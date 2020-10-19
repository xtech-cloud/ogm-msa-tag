package handler

import (
	"context"
	"omo-msa-tag/model"

	"github.com/micro/go-micro/v2/logger"
	proto "github.com/xtech-cloud/omo-msp-tag/proto/tag"
)

type Collection struct{}

func (this *Collection) AddTag(_ctx context.Context, _req *proto.CollectionAddTagRequest, _rsp *proto.BlankResponse) error {
	logger.Infof("Received Collection.AddTag, req is %v", _req)
	_rsp.Status = &proto.Status{}

	if "" == _req.Code {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "code is required"
		return nil
	}

	if "" == _req.Name {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "name is required"
		return nil
	}

	dao := model.NewCollectionDAO(nil)
	tag, err := dao.FindOne(_req.Code)
    if nil != err {
        return err
    }
	if nil != tag {
		_rsp.Status.Code = 2
		_rsp.Status.Message = "tag is exists"
		return nil
	}

	alias := make(map[string]string)
	if "" == _req.Alias {
		alias["en_US"] = _req.Name
	} else {
		alias["en_US"] = _req.Alias
	}

	tag = &model.Tag{
		Code:    _req.Code,
		Name:    _req.Name,
		Flag:    _req.Flag,
		Alias:   alias,
		Keyword: make([]string, 0),
        Dummy: make([]string, 0),
	}
	err = dao.InsertOne(tag)
	return err
}

func (this *Collection) RemoveTag(_ctx context.Context, _req *proto.CollectionRemoveTagRequest, _rsp *proto.BlankResponse) error {
	logger.Infof("Received Collection.RemoveTag, req is %v", _req)
	_rsp.Status = &proto.Status{}

	if "" == _req.Code {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "code is required"
		return nil
	}

	dao := model.NewCollectionDAO(nil)
	tag, err := dao.FindOne(_req.Code)
    if nil != err {
        return err
    }
    if nil == tag {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "tag not found"
		return nil
    }

    return dao.DeleteMany(_req.Code)
}

func (this *Collection) UpdateTag(_ctx context.Context, _req *proto.CollectionUpdateTagRequest, _rsp *proto.BlankResponse) error {
	logger.Infof("Received Collection.UpdateTag, req is %v", _req)
	_rsp.Status = &proto.Status{}

	if "" == _req.Code {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "code is required"
		return nil
	}

	if "" == _req.Name {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "name is required"
		return nil
	}

	dao := model.NewCollectionDAO(nil)
	tag, err := dao.FindOne(_req.Code)
    if nil != err {
        return err
    }
	if nil == tag {
		_rsp.Status.Code = 2
		_rsp.Status.Message = "tag not found"
		return nil
	}

	tag = &model.Tag{
		Code:    _req.Code,
		Name:    _req.Name,
		Flag:    _req.Flag,
		Alias:   _req.Alias,
		Keyword: _req.Keyword,
	}
	err = dao.UpdateOne(tag)
	return nil
}

func (this *Collection) ListTag(_ctx context.Context, _req *proto.CollectionListTagRequest, _rsp *proto.CollectionListTagResponse) error {
	logger.Infof("Received Collection.ListTag, req is %v", _req)
	_rsp.Status = &proto.Status{}

    offset := int64(0)
    count := int64(100)

    if _req.Offset > 0 {
        offset = _req.Offset
    }
    if _req.Count> 0 {
        count = _req.Count
    }

	dao := model.NewCollectionDAO(nil)
    total, err := dao.Count()
    if nil != err {
        return err
    }

    tags, err := dao.List(offset, count)
    if nil != err {
        return err
    }

    _rsp.Total = total
    _rsp.Tag = make([]*proto.TagEntity, len(tags))
    for i,tag := range tags {
        _rsp.Tag[i] = &proto.TagEntity {
            Code: tag.Code,
            Name: tag.Name,
            Flag: tag.Flag,
            Alias: tag.Alias,
            Keyword: tag.Keyword,
        }
    }

	return nil
}

func (this *Collection) SearchTag(_ctx context.Context, _req *proto.CollectionSearchTagRequest, _rsp *proto.CollectionSearchTagResponse) error {
	logger.Infof("Received Collection.SearchTag, req is %v", _req)
	_rsp.Status = &proto.Status{}

    if "" == _req.Filter{
        return nil
    }

    offset := int64(0)
    count := int64(100)

    if _req.Offset > 0 {
        offset = _req.Offset
    }
    if _req.Count> 0 {
        count = _req.Count
    }

	dao := model.NewCollectionDAO(nil)
    tags, err := dao.FindMany(_req.Filter, offset, count)
    if nil != err {
        return err
    }

    _rsp.Total = int64(len(tags))
    _rsp.Tag = make([]*proto.TagEntity, len(tags))
    for i,tag := range tags {
        _rsp.Tag[i] = &proto.TagEntity {
            Code: tag.Code,
            Name: tag.Name,
            Flag: tag.Flag,
            Alias: tag.Alias,
            Keyword: tag.Keyword,
        }
    }
    return nil
}

func (this *Collection) SuggestFilter(_ctx context.Context, _req *proto.CollectionSuggestFilterRequest, _rsp *proto.CollectionSuggestFilterResponse) error {
	logger.Infof("Received Collection.SuggestFilter, req is %v", _req)
	_rsp.Status = &proto.Status{}

    // TODO 收缩code,name and keyword
    result := model.SuggestFilter(_req.Input)
    _rsp.Filter= result
	return nil
}

func (this *Collection) ReplaceKeyword(_ctx context.Context, _req *proto.CollectionReplaceKeywordRequest, _rsp *proto.BlankResponse) error {
	logger.Infof("Received Collection.AddTag, req is %v", _req)
	_rsp.Status = &proto.Status{}

    if "" == _req.MatchedValue {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "matchedValue is required"
		return nil
    }

    if "" == _req.NewValue {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "newValue is required"
		return nil
    }

	return nil
}

func (this *Collection) ExtendKeyword(_ctx context.Context, _req *proto.CollectionExtendKeywordRequest, _rsp *proto.BlankResponse) error {
	logger.Infof("Received Collection.AddTag, req is %v", _req)
	_rsp.Status = &proto.Status{}

    if 0 == len(_req.MatchedValue) {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "matchedValue is required"
		return nil
    }

    if 0 == len(_req.NewValue) {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "newValue is required"
		return nil
    }
	return nil
}

func (this *Collection) MergeJson(_ctx context.Context, _req *proto.CollectionMergeJsonRequest, _rsp *proto.BlankResponse) error {
	logger.Infof("Received Collection.AddTag, req is %v", _req)
	_rsp.Status = &proto.Status{}
	return nil
}
