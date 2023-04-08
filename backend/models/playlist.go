package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Playlist struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name,omitempty"`
	Songs []string           `json:"songs"`
}
