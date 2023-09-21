package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type FileMongoEntity struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Indicador1 int                `bson:"Indicador1" json:"Indicador1" validate:"required"`
	Indicador2 int                `bson:"Indicador2" json:"Indicador2" validate:"required"`
	Indicador3 int                `bson:"Indicador3" json:"Indicador3" validate:"required"`
}
