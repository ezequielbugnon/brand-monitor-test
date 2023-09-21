package repository

import (
	"context"

	"github.com/ezequiel-bugnon/brandmonitor/db"
	"github.com/ezequiel-bugnon/brandmonitor/entity"
	"go.mongodb.org/mongo-driver/bson"
)

type repository struct {
	conection db.MongoDB
}

func NewRepository(conection db.MongoDB) *repository {
	return &repository{
		conection,
	}
}

func (s *repository) Post(file entity.FileMongoEntity) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, err := s.conection.Query().InsertOne(ctx, file)
	if err != nil {
		return err
	}

	return nil
}

func (s *repository) Get() ([]entity.FileMongoEntity, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cursor, err := s.conection.Query().Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []entity.FileMongoEntity

	for cursor.Next(ctx) {
		var file entity.FileMongoEntity

		if err := cursor.Decode(&file); err != nil {
			return nil, err
		}

		results = append(results, file)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return results, nil
}
