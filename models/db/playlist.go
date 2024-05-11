package models

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Playlist struct {
	mgm.DefaultModel `bson:",inline"`
	Id               string               `json:"id" bson:"_id"`
	Title            string               `json:"title" bson:"title"`
	Track            []primitive.ObjectID `json:"track" bson:"track"`
}

func NewPlaylist(title string, track []primitive.ObjectID) *Playlist {
	return &Playlist{
		Title: title,
		Track: track,
	}
}

func (model *Playlist) CollectionName() string {
	return "playlists"
}
