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

func CreatePlaylist(playlistId primitive.ObjectID, playlistRequest *models.PlaylistRequest) (*db.Playlist, error) {
	playlist := db.NewPlaylist(
		playlistRequest.Title,
		playlistRequest.Tracks,
		primitive.NewObjectID())
	// playlistRequest.Owner)

	err := mgm.Coll(playlist).Create(playlist)

	if err != nil {
		return nil, errors.New("Failed to create playlist")
	}

	return playlist, nil
}

func DeletePlaylist(playlistId primitive.ObjectID) error {
	deleteResult, err := mgm.Coll(&db.Playlist{}).DeleteOne(mgm.Ctx(), bson.M{field.ID: playlistId})

	if err != nil || deleteResult.DeletedCount <= 0 {
		return errors.New("Failed to delete playlist")
	}

	return nil
}

func GetPlaylistById(playlistId primitive.ObjectID) (*db.Playlist, error) {
	playlist := &db.Playlist{}
	err := mgm.Coll(playlist).FindByID(playlistId, playlist)

	if err != nil {
		return nil, errors.New("Failed to find playlist")
	}

	return playlist, nil
}

func GetPlaylists(page int, size int, sort string, filter map[string]interface{}) ([]*db.Playlist, error) {
	playlists := []*db.Playlist{}
	query := bson.M{}

	for key, value := range filter {
		query[key] = value
	}

	limit := int64(size)
	skip := int64((page - 1) * size)

	err := mgm.Coll(&db.Playlist{}).SimpleFind(&playlists, query, &options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
		Sort:  sort,
	})

	if err != nil {
		return nil, errors.New("Failed to get playlists")
	}

	return playlists, nil
}

func UpdatePlaylist(playlistId primitive.ObjectID, playlistRequest *models.PlaylistRequest) error {
	playlist := &db.Playlist{}
	err := mgm.Coll(playlist).FindByID(playlistId, playlist)

	if err != nil {
		return errors.New("Failed to find playlist")
	}

	playlist.Title = playlistRequest.Title
	playlist.Track = playlistRequest.Tracks
	playlist.Owner = playlistRequest.Owner

	err = mgm.Coll(playlist).Update(playlist)

	if err != nil {
		return errors.New("Failed to update playlist")
	}
	return nil
}
