package dto

import "retail-demo/retail-management/model"

type ImportProductDto struct {
	InventoryId string          `bson:"inventoryName,omitempty"`
	Products    []model.Product `bson:"products,omitempty"`
}
