package business

import (
	"context"
	"retail-demo/retail-management/model"
	"retail-demo/retail-management/model/dto"
)

type SellProductStorage interface {
	SellProduct(ctx context.Context, dto dto.SellProductDto) (model.Invoice, error)
}

type sellProductBiz struct {
	store SellProductStorage
}

func SellProductBiz(store SellProductStorage) *sellProductBiz {
	return &sellProductBiz{store: store}
}

func (biz *sellProductBiz) SellProduct(ctx context.Context, dto dto.SellProductDto) (model.Invoice, error) {
	return biz.store.SellProduct(ctx, dto)
}
