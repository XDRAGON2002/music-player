package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Song struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	SongID   string             `json:"songId,omitempty"`
	Image    string             `json:"image,omitempty"`
	SongName string             `json:"songName,omitempty"`
	Likes    int                `json:"likes,omitempty"`
}
