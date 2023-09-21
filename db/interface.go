package db

import "go.mongodb.org/mongo-driver/mongo"

type MongoDBInterface interface {
	Query() *mongo.Collection
	Close()
}
