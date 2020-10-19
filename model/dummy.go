package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DummyDAO struct {
	conn *Conn
}

func NewDummyDAO(_conn *Conn) *DummyDAO {
	if nil == _conn {
		return &DummyDAO{
			conn: defaultConn,
		}
	} else {
		return &DummyDAO{
			conn: _conn,
		}
	}
}

func (this *DummyDAO) Filter(_offset int64, _count int64, _code []string) (int64, []string, error) {
	ctx, cancel := NewContext()
	defer cancel()

	filter := bson.M{
		"code": bson.M{
			"$in": _code,
		},
	}

	// 1: 倒叙  -1：正序
	sort := bson.D{{"code", -1}}

	findOptions := options.Find()
	findOptions.SetSort(sort)

	cur, err := this.conn.DB.Collection(CollectionName).Find(ctx, filter, findOptions)
	if nil != err {
		return 0, make([]string, 0), err
	}
	defer cur.Close(ctx)

	total := int64(0)
	dummyCache := make(map[int][]string)
	row := 0
	for cur.Next(ctx) {
		var tag Tag
		err = cur.Decode(&tag)
		if nil != err {
			return 0, make([]string, 0), err
		}
		total = total + int64(len(tag.Dummy))
		dummyCache[row] = tag.Dummy
		row = row + 1
	}

	offset := _offset
	if _offset >= total {
		offset = total
	}
	count := _count
	if count+offset >= total {
		count = total - offset
	}

	dummyAry := make([]string, count)
	pos := int64(0)
	index := int64(0)
	for i := 0; i < row; i++ {
		if pos+int64(len(dummyCache[i])) < offset {
			pos = pos + int64(len(dummyCache[i]))
			continue
		}
		left := offset - pos
		right := int64(len(dummyCache[i]))
		if right > left+count {
			right = left + count
		}
		size := right - left
		copy(dummyAry[index:index+size], dummyCache[i][left:right])
		index = index + size
		if index >= count {
			break
		}
	}
	return total, dummyAry, nil
}
