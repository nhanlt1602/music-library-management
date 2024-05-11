package models

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MusicTrack struct {
	mgm.DefaultModel `bson:",inline"`
	Id               primitive.ObjectID `json:"id" bson:"_id"`
	Title            string             `json:"title" bson:"title"`
	Artist           string             `json:"artist" bson:"artist"`
	Album            string             `json:"album" bson:"album"`
	Genre            string             `json:"genre" bson:"genre"`
	ReleaseYear      int                `json:"release_year" bson:"release_year"`
	Duration         int                `json:"duration" bson:"duration"`
	Mp3File          string             `json:"mp3_file" bson:"mp3_file"`
}

func NewMusicTrack(title string, artist string, album string, genre string, releaseYear int, duration int, mp3File string) *MusicTrack {
	return &MusicTrack{
		Title:       title,
		Artist:      artist,
		Album:       album,
		Genre:       genre,
		ReleaseYear: releaseYear,
		Duration:    duration,
		Mp3File:     mp3File,
	}
}

func (model *MusicTrack) CollectionName() string {
	return "musictracks"
}
