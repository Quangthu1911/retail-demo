package storage

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"retail-demo/retail-management/model"
	"retail-demo/retail-management/model/dto"
	"time"
)

func (s *mongodbStorage) SellProduct(ctx context.Context, dto dto.SellProductDto) (model.Invoice, error) {
	var branch model.Branch
	coll := s.db.Collection("branch")
	err := coll.FindOne(ctx, bson.M{"name": dto.BranchName}).Decode(&branch)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.Invoice{}, errors.New("Branch not found")
		}
		return model.Invoice{}, err
	}
	coll = s.db.Collection("inventory")
	cur, err := coll.Find(ctx, bson.M{"requestId": bson.M{"$in": branch.Inventory}})
	var inventory []model.Inventory
	if err = cur.All(context.TODO(), &inventory); err != nil {
		return model.Invoice{}, errors.New("Has some error in process")
	}
	var arrBuy []model.Product
	totalAmount := 0
	for _, reqProduct := range dto.Products {
		flag := false
		// Check inventory of branch
		for _, inventory := range inventory {
			// Check product of inventory
			for _, product := range inventory.Products {
				if reqProduct.RequestId == product.RequestId {
					if product.Quantity < reqProduct.Quantity {
						continue
					}
					query := bson.M{
						"requestId":          inventory.RequestId,
						"products.requestId": reqProduct.RequestId,
					}
					update := bson.D{
						{"$inc", bson.D{{"products.$.quantity", reqProduct.Quantity * -1}}},
					}
					_, err = coll.UpdateOne(ctx, query, update)
					var data model.Product
					err = s.db.Collection("product").FindOne(ctx, bson.M{"requestId": reqProduct.RequestId}).Decode(&data)
					totalAmount = totalAmount + (data.Amount * reqProduct.Quantity)
					data.Quantity = reqProduct.Quantity
					arrBuy = append(arrBuy, data)
					flag = true
					break
				}
			}
			if flag {
				break
			}
		}
	}
	invoice := model.Invoice{primitive.ObjectID{}, time.Now(), arrBuy, totalAmount}
	_, err = s.db.Collection("invoice").InsertOne(ctx, invoice)
	return invoice, nil
}
