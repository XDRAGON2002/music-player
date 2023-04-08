package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty"`
	Email     string             `json:"email,omitempty"`
	Password  string             `json:"password,omitempty"`
	Playlists []string           `json:"playlists"`
}
