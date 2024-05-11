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
	Owner            primitive.ObjectID   `json:"owner" bson:"owner"`
}

func NewPlaylist(title string, track []primitive.ObjectID, owner primitive.ObjectID) *Playlist {
	return &Playlist{
		Title: title,
		Track: track,
		Owner: owner,
	}
}

func (model *Playlist) CollectionName() string {
	return "playlists"
}
