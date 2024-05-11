package services

import (
	"context"
	"errors"
	"music-library-management/models"
	db "music-library-management/models/db"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteMusicTrack(ctx context.Context, musicTrackId primitive.ObjectID) error {
	panic("unimplemented")
}

func CreateMusicTrack(musicTrackId primitive.ObjectID, musicTrackRequest *models.MusicTrackRequest) (*db.MusicTrack, error) {
	musicTrack := db.NewMusicTrack(
		musicTrackId,
		musicTrackRequest.Title,
		musicTrackRequest.Artist,
		musicTrackRequest.Album,
		musicTrackRequest.Genre,
		musicTrackRequest.ReleaseYear,
		musicTrackRequest.Duration,
		musicTrackRequest.Mp3File,
	)

	err := mgm.Coll(musicTrack).Create(musicTrack)
	if err != nil {
		return nil, errors.New("Failed to create music track")
	}

	return musicTrack, nil
}
