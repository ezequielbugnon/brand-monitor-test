package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client *mongo.Client
}

func NewMongoDB() (*MongoDB, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("¡Conexión exitosa a MongoDB!")

	return &MongoDB{client}, nil
}

func (m *MongoDB) Query() *mongo.Collection {
	database := m.client.Database("miBaseDeDatos")
	collection := database.Collection("miColeccion")
	return collection
}

func (m *MongoDB) Close() {
	if m.client != nil {
		m.client.Disconnect(context.Background())
	}
}
