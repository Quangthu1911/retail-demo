package dto

import "retail-demo/retail-management/model"

type SellProductDto struct {
	BranchName string          `bson:"branchName,omitempty"`
	Products   []model.Product `bson:"products,omitempty"`
}
