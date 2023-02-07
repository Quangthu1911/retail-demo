package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"retail-demo/retail-management/model"
	"retail-demo/retail-management/model/dto"
)

func (s *mongodbStorage) GetInfo(ctx context.Context) []dto.GetInfoDto {
	var branchs []model.Branch
	coll := s.db.Collection("branch")
	cur, _ := coll.Find(ctx, bson.D{})
	_ = cur.All(context.TODO(), &branchs)
	var ouput []dto.GetInfoDto
	for _, branch := range branchs {
		cur, _ = s.db.Collection("inventory").Find(ctx, bson.M{"requestId": bson.M{"$in": branch.Inventory}})
		var inventory []model.Inventory
		_ = cur.All(context.TODO(), &inventory)
		ouput = append(ouput, dto.GetInfoDto{branch, inventory})
	}
	return ouput
}
