package storage

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"retail-demo/retail-management/model"
	"retail-demo/retail-management/model/dto"
)

func (s *mongodbStorage) ImportProduct(ctx context.Context, dto dto.ImportProductDto) error {
	var inventory model.Inventory
	coll := s.db.Collection("inventory")

	errInventory := coll.FindOne(ctx, bson.M{"requestId": dto.InventoryId}).Decode(&inventory)
	if errInventory != nil {
		if errors.Is(errInventory, mongo.ErrNoDocuments) {
			return errors.New("Inventory not found")
		}
		return errInventory
	}
	var err error
	for _, reqPd := range dto.Products {
		flag := false
		fmt.Println(inventory.Name)
		for _, product := range inventory.Products {
			fmt.Println(reqPd.RequestId, product.RequestId)
			if reqPd.RequestId == product.RequestId {
				flag = true
				break
			}
		}
		if flag {
			query := bson.M{
				"requestId":          dto.InventoryId,
				"products.requestId": reqPd.RequestId,
			}
			update := bson.D{
				{"$inc", bson.D{{"products.$.quantity", reqPd.Quantity}}},
			}
			_, err = coll.UpdateOne(ctx, query, update)
		} else {
			var product model.Product
			if errProduct := s.db.Collection("product").FindOne(ctx, bson.M{"requestId": reqPd.RequestId}).Decode(&product); errProduct != nil {
				continue
			}
			product.Quantity = reqPd.Quantity
			change := bson.M{
				"$push": bson.M{
					"products": product,
				},
			}
			query := bson.M{
				"requestId": dto.InventoryId,
			}
			_, err = coll.UpdateOne(ctx, query, change)
		}
	}
	if err != nil {
		return err
	}
	return nil
}
