package storage

import (
	"context"
	"errors"
	"fmt"
	"retail-demo/retail-management/model"
	"retail-demo/retail-management/model/dto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *mongodbStorage) MoveProduct(ctx context.Context, dto dto.MoveProductDto) error {
	var inventory model.Inventory
	coll := s.db.Collection("inventory")

	errInventory := coll.FindOne(ctx, bson.M{"requestId": dto.InputInventory.RequestId}).Decode(&inventory)
	if errInventory != nil {
		if errors.Is(errInventory, mongo.ErrNoDocuments) {
			return errors.New("Inventory not found")
		}
		return errInventory
	}
	for _, reqPd := range dto.InputInventory.Products {
		for _, product := range inventory.Products {
			if reqPd.RequestId == product.RequestId {
				if reqPd.Quantity > product.Quantity {
					continue
				}
				var outputInventory model.Inventory
				coll.FindOne(ctx, bson.M{"requestId": dto.OutputInventory}).Decode(&outputInventory)
				flag := false
				for _, prd := range outputInventory.Products {
					if reqPd.RequestId == prd.RequestId {
						fmt.Println("qqqqq", reqPd.RequestId, reqPd.Quantity, outputInventory.RequestId)
						query := bson.M{
							"requestId":          outputInventory.RequestId,
							"products.requestId": reqPd.RequestId,
						}
						update := bson.D{
							{"$inc", bson.D{{"products.$.quantity", reqPd.Quantity}}},
						}
						coll.UpdateOne(ctx, query, update)
						updateInventoryOld(reqPd, dto.InputInventory.RequestId, ctx, s)
						flag = true
						break
					}
				}
				if !flag {
					var sp model.Product
					if errProduct := s.db.Collection("product").FindOne(ctx, bson.M{"requestId": reqPd.RequestId}).Decode(&sp); errProduct != nil {
						continue
					}
					sp.Quantity = reqPd.Quantity
					change := bson.M{
						"$push": bson.M{
							"products": sp,
						},
					}
					query := bson.M{
						"requestId": outputInventory.RequestId,
					}
					coll.UpdateOne(ctx, query, change)
					updateInventoryOld(reqPd, dto.InputInventory.RequestId, ctx, s)
				}
			}
		}
	}
	return nil
}

func updateInventoryOld(reqPd model.Product, requestId string, ctx context.Context, s *mongodbStorage) {
	query := bson.M{
		"requestId":          requestId,
		"products.requestId": reqPd.RequestId,
	}
	update := bson.D{
		{"$inc", bson.D{{"products.$.quantity", reqPd.Quantity * -1}}},
	}
	s.db.Collection("inventory").UpdateOne(ctx, query, update)
}
