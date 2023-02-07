package business

import (
	"context"
	"retail-demo/retail-management/model/dto"
)

type ImportProductStorage interface {
	ImportProduct(ctx context.Context, dto dto.ImportProductDto) error
}

type importProductBiz struct {
	store ImportProductStorage
}

func ImportProductBiz(store ImportProductStorage) *importProductBiz {
	return &importProductBiz{store: store}
}

func (biz *importProductBiz) ImportProduct(ctx context.Context, dto dto.ImportProductDto) error {
	return biz.store.ImportProduct(ctx, dto)
}
