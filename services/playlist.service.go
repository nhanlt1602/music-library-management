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
		playlistId,
		playlistRequest.Title,
		playlistRequest.Track,
		playlistRequest.Owner)

	err := mgm.Coll(playlist).Create(playlist)

	if err != nil {
		// return nil, errors.New("Failed to create playlist")
		return nil, err
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

func GetPlaylists(request *models.GetPlaylistRequest) ([]*db.Playlist, error) {
	playlists := []*db.Playlist{}
	limit := int64(request.Paging.Size)
	page := int64(request.Paging.Page)
	var err error

	if request.Paging.PagingIgnore {
		err = mgm.Coll(&db.Playlist{}).SimpleFind(&playlists, &options.FindOptions{})
	} else {
		findOptions := options.Find().
			SetSkip(int64(page * limit)).
			SetLimit(int64(limit))

		err = mgm.Coll(&db.Playlist{}).SimpleFind(&playlists, bson.M{}, findOptions)
	}

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

	playlistUpdate := bson.M{
		"$set": bson.M{
			"title": playlistRequest.Title,
			"track": playlistRequest.Track,
			"owner": playlistRequest.Owner,
		},
	}

	_, err = mgm.Coll(playlist).UpdateOne(mgm.Ctx(), bson.M{field.ID: playlistId}, playlistUpdate)

	if err != nil {
		return errors.New("Failed to update playlist")
	}

	return nil
}
