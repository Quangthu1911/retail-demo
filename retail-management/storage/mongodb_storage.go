package storage

import "go.mongodb.org/mongo-driver/mongo"

type mongodbStorage struct {
	db *mongo.Database
}

func NewMongoDbStorage(db *mongo.Database) *mongodbStorage {
	return &mongodbStorage{db: db}
}
