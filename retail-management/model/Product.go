package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Name           string             `bson:"name,omitempty"`
	Amount         int                `bson:"amount,omitempty"`
	OriginalAmount int                `bson:"originalAmount,omitempty"`
	Quantity       int                `bson:"quantity,omitempty"`
	RequestId      string             `bson:"requestId,omitempty"`
}
