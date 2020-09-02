package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Tag struct {
	Code    string            `bson:"code"`
	Name    string            `bson:"name"`
	Flag    int64             `bson:"flag"`
	Alias   map[string]string `bson:"alias"`
	Keyword []string          `bson:"keyword"`
}

const (
	CollectionName = "msa_tag_collection"
)

type CollectionDAO struct {
	conn *Conn
}

func NewCollectionDAO(_conn *Conn) *CollectionDAO {
	if nil == _conn {
		return &CollectionDAO{
			conn: defaultConn,
		}
	} else {
		return &CollectionDAO{
			conn: _conn,
		}
	}
}

func (this *CollectionDAO) InsertOne(_tag *Tag) error {
	ctx, cancel := NewContext()
	defer cancel()

	document, err := bson.Marshal(_tag)
	if nil != err {
		return err
	}

	_, err = this.conn.DB.Collection(CollectionName).InsertOne(ctx, document)
	if nil != err {
		return err
	}
	return nil
}

func (this *CollectionDAO) FindOne(_code string) (*Tag, error) {
	ctx, cancel := NewContext()
	defer cancel()

	filter := bson.D{{"code", _code}}
	res := this.conn.DB.Collection(CollectionName).FindOne(ctx, filter)
	if res.Err() == mongo.ErrNoDocuments {
		return nil, nil
	}
	var tag Tag
	err := res.Decode(&tag)
	return &tag, err
}

func (this *CollectionDAO) Count() (int64, error) {
	ctx, cancel := NewContext()
	defer cancel()
	count, err := this.conn.DB.Collection(CollectionName).EstimatedDocumentCount(ctx)
	return count, err
}

func (this *CollectionDAO) List(_offset int64, _count int64) ([]*Tag, error) {
	ctx, cancel := NewContext()
	defer cancel()

	filter := bson.D{}
	// 1: 倒叙  -1：正序
	sort := bson.D{{"code", -1}}

	findOptions := options.Find()
	findOptions.SetSort(sort)
	findOptions.SetSkip(_offset)
	findOptions.SetLimit(_count)

	cur, err := this.conn.DB.Collection(CollectionName).Find(ctx, filter, findOptions)
	if nil != err {
		return make([]*Tag, 0), err
	}
	defer cur.Close(ctx)

	var tags []*Tag
	for cur.Next(ctx) {
		var tag Tag
		err = cur.Decode(&tag)
		if nil != err {
			return make([]*Tag, 0), err
		}
		tags = append(tags, &tag)
	}
	return tags, nil
}

func (this *CollectionDAO) FindMany(_filter string, _offset int64, _count int64) ([]*Tag, error) {
	ctx, cancel := NewContext()
	defer cancel()

	// 模糊查询并忽略大小写
    /*
	regex := bson.M{
		"$regex":   _filter,
		"$options": "$i",
	}

	filter := bson.D{
		{"$or", bson.A{
			bson.M{"code": regex},
			bson.M{"name": regex},
			//bson.M{"alias": regex},
			bson.M{"keyword": regex},
		}},
	}
    */

    //TODO 从别名中搜索 
	filter := bson.D{
		{"$or", bson.A{
			bson.M{"code": _filter},
			bson.M{"name": _filter},
			//bson.M{"alias.en_US": _filter},
			bson.M{"keyword": _filter},
		}},
	}

	// 1: 倒叙  -1：正序
	sort := bson.D{{"code", -1}}

	findOptions := options.Find()
	findOptions.SetSort(sort)
	findOptions.SetSkip(_offset)
	findOptions.SetLimit(_count)

	cur, err := this.conn.DB.Collection(CollectionName).Find(ctx, filter, findOptions)
	if nil != err {
		return make([]*Tag, 0), err
	}
	defer cur.Close(ctx)

	var tags []*Tag
	for cur.Next(ctx) {
		var tag Tag
		err = cur.Decode(&tag)
		if nil != err {
			return make([]*Tag, 0), err
		}
		tags = append(tags, &tag)
	}
	return tags, nil
}

func (this *CollectionDAO) UpdateOne(_tag *Tag) error {
	ctx, cancel := NewContext()
	defer cancel()

	filter := bson.D{{"code", _tag.Code}}
	update := bson.D{
		{"$set", bson.D{
			{"name", _tag.Name},
			{"flag", _tag.Flag},
			{"alias", _tag.Alias},
			{"keyword", _tag.Keyword},
		}},
	}
	_, err := this.conn.DB.Collection(CollectionName).UpdateOne(ctx, filter, update)
	if nil != err {
		return err
	}

	increaseFilter([]*Tag{_tag})
	return nil
}

func (this *CollectionDAO) DeleteMany(_code string) error {
	ctx, cancel := NewContext()
	defer cancel()

	filter := bson.D{{"code", _code}}
	_, err := this.conn.DB.Collection(CollectionName).DeleteMany(ctx, filter)

	//TODO
	//reduceKeyword()
	return err
}
