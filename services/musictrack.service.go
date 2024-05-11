package services

import (
	"errors"
	"music-library-management/models"
	db "music-library-management/models/db"

	"github.com/kamva/mgm/v3"

	"github.com/kamva/mgm/v3/field"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Delete Music Track delete a music track by id
func DeleteMusicTrack(musicTrackId primitive.ObjectID) error {
	deleteResult, err := mgm.Coll(&db.MusicTrack{}).DeleteOne(mgm.Ctx(), bson.M{field.ID: musicTrackId})

	if err != nil || deleteResult.DeletedCount <= 0 {
		return errors.New("Failed to delete music track")
	}

	return nil
}

func GetMusicTrackById(musicTrackId primitive.ObjectID) (*db.MusicTrack, error) {
	musicTrack := &db.MusicTrack{}
	err := mgm.Coll(musicTrack).FindByID(musicTrackId, musicTrack)

	if err != nil {
		return nil, errors.New("Failed to find music track")
	}

	return musicTrack, nil
}

// Get Music Tracks get all music tracks, have paging, sorting, and filtering
func GetMusicTracks(page int, size int, sort string, filter map[string]interface{}) ([]*db.MusicTrack, error) {
	musicTracks := []*db.MusicTrack{}
	query := bson.M{}

	for key, value := range filter {
		query[key] = value
	}

	limit := int64(size)
	skip := int64((page - 1) * size)

	err := mgm.Coll(&db.MusicTrack{}).SimpleFind(&musicTracks, query, &options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
		Sort:  sort,
	})

	if err != nil {
		return nil, errors.New("Failed to get music tracks")
	}

	return musicTracks, nil
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

func UpdateMusicTrack(musicTrackId primitive.ObjectID, musicTrackRequest *models.MusicTrackRequest) error {
	musicTrack := &db.MusicTrack{}
	err := mgm.Coll(musicTrack).FindByID(musicTrackId, musicTrack)

	if err != nil {
		return errors.New("Failed to find music track")
	}

	musicTrack.Title = musicTrackRequest.Title
	musicTrack.Artist = musicTrackRequest.Artist
	musicTrack.Album = musicTrackRequest.Album
	musicTrack.Genre = musicTrackRequest.Genre
	musicTrack.ReleaseYear = musicTrackRequest.ReleaseYear
	musicTrack.Duration = musicTrackRequest.Duration
	musicTrack.Mp3File = musicTrackRequest.Mp3File
	err = mgm.Coll(musicTrack).Update(musicTrack)

	if err != nil {
		return errors.New("Failed to update music track")
	}

	return nil
}
