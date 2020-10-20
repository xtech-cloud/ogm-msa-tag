package handler

import (
	"context"
	"omo-msa-tag/model"

	"github.com/micro/go-micro/v2/logger"
	proto "github.com/xtech-cloud/omo-msp-tag/proto/tag"
)

type Dummy struct{}

func (this *Dummy) AddTag(_ctx context.Context, _req *proto.DummyAddTagRequest, _rsp *proto.BlankResponse) error {
	logger.Infof("Received Dummy.AddTag, req is %v", _req)
	_rsp.Status = &proto.Status{}

	if "" == _req.Code {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "code is required"
		return nil
	}

	if "" == _req.Owner {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "owner is required"
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

    return dao.AppendDummy(tag, _req.Owner)
}

func (this *Dummy) RemoveTag(_ctx context.Context, _req *proto.DummyRemoveTagRequest, _rsp *proto.BlankResponse) error {
	logger.Infof("Received Dummy.RemoveTag, req is %v", _req)
	_rsp.Status = &proto.Status{}

	if "" == _req.Code {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "code is required"
		return nil
	}

	if "" == _req.Owner {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "owner is required"
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

	return dao.RemoveDummy(tag, _req.Owner)
}

func (this *Dummy) FilterTag(_ctx context.Context, _req *proto.DummyFilterTagRequest, _rsp *proto.DummyFilterTagResponse) error {
	logger.Infof("Received Dummy.ListTag, req is %v", _req)
	_rsp.Status = &proto.Status{}

	offset := int64(0)
	count := int64(100)

	if _req.Offset > 0 {
		offset = _req.Offset
	}
	if _req.Count > 0 {
		count = _req.Count
	}

	dao := model.NewDummyDAO(nil)
    total, dummy, err := dao.Filter(offset, count, _req.Code)
	if nil != err {
		return err
	}
    _rsp.Total = total
    _rsp.Owner = dummy
	return nil
}

func (this *Dummy) ListTag(_ctx context.Context, _req *proto.DummyListTagRequest, _rsp *proto.DummyListTagResponse) error {
	logger.Infof("Received Dummy.ListTag, req is %v", _req)
	_rsp.Status = &proto.Status{}

	offset := int64(0)
	count := int64(100)

	if "" == _req.Owner{
		_rsp.Status.Code = 1
		_rsp.Status.Message = "owner is required"
		return nil
	}

	if _req.Offset > 0 {
		offset = _req.Offset
	}
	if _req.Count > 0 {
		count = _req.Count
	}
	dao := model.NewDummyDAO(nil)
    total, dummy, err := dao.List(offset, count, _req.Owner)
	if nil != err {
		return err
	}
    _rsp.Total = total
    _rsp.Owner = dummy
	return nil
}
