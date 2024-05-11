package models

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Playlist struct {
	mgm.DefaultModel `bson:",inline"`
	Id               primitive.ObjectID   `json:"id" bson:"_id"`
	Title            string               `json:"title" bson:"title"`
	Track            []primitive.ObjectID `json:"track" bson:"track"`
	Owner            primitive.ObjectID   `json:"owner" bson:"owner"`
}

func NewPlaylist(id primitive.ObjectID, title string, track []primitive.ObjectID, owner primitive.ObjectID) *Playlist {
	return &Playlist{
		Id:    id,
		Title: title,
		Track: track,
		Owner: owner,
	}
}

func (model *Playlist) CollectionName() string {
	return "playlists"
}
