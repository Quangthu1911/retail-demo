package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Inventory struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name,omitempty"`
	Products  []Product          `bson:"products,omitempty"`
	RequestId string             `bson:"requestId,omitempty"`
}
