package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Branch struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name,omitempty"`
	Address   string             `bson:"address,omitempty"`
	Inventory []string           `bson:"inventory,omitempty"`
}
