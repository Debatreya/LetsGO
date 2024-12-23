package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct { // this struct will be used to define the schema of the collection and it is exported
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string             `json:"movie,omitempty" bson:"movie,omitempty"`
	Watched bool               `json:"watched,omitempty"`
}

