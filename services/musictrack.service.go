package services

import (
	"errors"
	"log"
	"music-library-management/models"
	db "music-library-management/models/db"
	"strings"

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
func GetMusicTracks(request *models.GetMusicTrackRequest) ([]*db.MusicTrack, error) {
	musicTracks := []*db.MusicTrack{}
	limit := int64(request.Paging.Size)
	page := int64(request.Paging.Page)
	var err error

	query := bson.M{}

	if request.Title != "" {
		query["title"] = bson.M{"$regex": request.Title, "$options": "i"}
	}

	if request.Artist != "" {
		query["artist"] = bson.M{"$regex": request.Artist, "$options": "i"}
	}

	if request.Album != "" {
		query["album"] = bson.M{"$regex": request.Album, "$options": "i"}
	}

	if request.Genre != "" {
		query["genre"] = bson.M{"$regex": request.Genre, "$options": "i"}
	}

	findOptions := options.Find()

	if request.Paging.Sort != "" {
		sort := strings.Split(request.Paging.Sort, ",")
		sortField := sort[0]
		sortOrder := 1 // ascending order by default

		if sort[1] == "desc" {
			sortOrder = -1 // descending order
		}

		findOptions.SetSort(bson.M{sortField: sortOrder})
	}

	if request.Paging.PagingIgnore {
		err = mgm.Coll(&db.MusicTrack{}).SimpleFind(&musicTracks, query, findOptions)
	} else {
		findOptions.SetSkip(int64(page * limit))
		findOptions.SetLimit(int64(limit))
		err = mgm.Coll(&db.MusicTrack{}).SimpleFind(&musicTracks, query, findOptions)
	}

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
		log.Println("Failed to create music track with error: ", err)
		return nil, err
	}

	return musicTrack, nil
}

func UpdateMusicTrack(musicTrackId primitive.ObjectID, musicTrackRequest *models.MusicTrackRequest) error {
	musicTrack := &db.MusicTrack{}
	err := mgm.Coll(musicTrack).FindByID(musicTrackId, musicTrack)

	if err != nil {
		return errors.New("Failed to find music track")
	}

	musicTrackUpdate := bson.M{
		"$set": bson.M{
			"title":        musicTrackRequest.Title,
			"artist":       musicTrackRequest.Artist,
			"album":        musicTrackRequest.Album,
			"genre":        musicTrackRequest.Genre,
			"release_year": musicTrackRequest.ReleaseYear,
			"duration":     musicTrackRequest.Duration,
			"mp3_file":     musicTrackRequest.Mp3File,
		},
	}

	_, err = mgm.Coll(musicTrack).UpdateOne(mgm.Ctx(), bson.M{field.ID: musicTrackId}, musicTrackUpdate)

	if err != nil {
		return errors.New("Failed to update music track")
	}

	return nil
}
