package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Invoice struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	CreatedDate time.Time          `bson:"createdDate,omitempty"`
	Products    []Product          `bson:"products,omitempty"`
	TotalAmount int                `bson:"totalAmount,omitempty"`
}
