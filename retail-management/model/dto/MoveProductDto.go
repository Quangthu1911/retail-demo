package dto

import "retail-demo/retail-management/model"

type MoveProductDto struct {
	InputInventory  model.Inventory `bson:"inputInventory,omitempty"`
	OutputInventory string          `bson:"outputInventory,omitempty"`
}
