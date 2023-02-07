package business

import (
	"context"
	"retail-demo/retail-management/model/dto"
)

type GetInfoStorage interface {
	GetInfo(ctx context.Context) []dto.GetInfoDto
}

type getInfoBiz struct {
	store GetInfoStorage
}

func GetInfoBiz(store GetInfoStorage) *getInfoBiz {
	return &getInfoBiz{store: store}
}

func (biz *getInfoBiz) GetInfo(ctx context.Context) []dto.GetInfoDto {
	return biz.store.GetInfo(ctx)
}
