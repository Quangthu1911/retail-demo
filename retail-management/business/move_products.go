package business

import (
	"context"
	"retail-demo/retail-management/model/dto"
)

type MoveProductStorage interface {
	MoveProduct(ctx context.Context, dto dto.MoveProductDto) error
}

type moveProductBiz struct {
	store MoveProductStorage
}

func MoveProductBiz(store MoveProductStorage) *moveProductBiz {
	return &moveProductBiz{store: store}
}

func (biz *moveProductBiz) MoveProduct(ctx context.Context, dto dto.MoveProductDto) error {
	return biz.store.MoveProduct(ctx, dto)
}
