package dto

import "retail-demo/retail-management/model"

type GetInfoDto struct {
	Branch          model.Branch      `bson:"branch,omitempty"`
	InventoryDetail []model.Inventory `bson:"inventoryDetail,omitempty"`
}
