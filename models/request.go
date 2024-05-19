package models

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var passwordRule = []validation.Rule{
	validation.Required,
	validation.Length(8, 32),
	validation.Match(regexp.MustCompile("^\\S+$")).Error("cannot contain whitespaces"),
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a RegisterRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Name, validation.Required, validation.Length(3, 64)),
		validation.Field(&a.Email, validation.Required, is.Email),
		validation.Field(&a.Password, passwordRule...),
	)
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a LoginRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Email, validation.Required, is.Email),
		validation.Field(&a.Password, passwordRule...),
	)
}

type RefreshRequest struct {
	Token string `json:"token"`
}

func (a RefreshRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Token, validation.Required, validation.Match(regexp.MustCompile("^\\S+$")).Error("cannot contain whitespaces")),
	)
}

type MusicTrackRequest struct {
	Title       string `json:"title"`
	Artist      string `json:"artist"`
	Album       string `json:"album"`
	Genre       string `json:"genre"`
	ReleaseYear int    `json:"release_year"`
	Duration    int    `json:"duration"`
	Mp3File     string `json:"mp3_file"`
}

type GetPlaylistRequest struct {
	Title  string               `json:"title"`
	Owner  primitive.ObjectID   `json:"owner"`
	Track  []primitive.ObjectID `json:"track"`
	Paging PagingRequest        `json:"paging"`
}

type GetMusicTrackRequest struct {
	Title  string        `json:"title"`
	Artist string        `json:"artist"`
	Album  string        `json:"album"`
	Genre  string        `json:"genre"`
	Paging PagingRequest `json:"paging"`
}

type PlaylistRequest struct {
	Title string               `json:"title"`
	Track []primitive.ObjectID `json:"track"`
	Owner primitive.ObjectID   `json:"owner"`
}

func (a PlaylistRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Title, validation.Required),
		validation.Field(&a.Track, validation.Required),
		validation.Field(&a.Owner, validation.Required),
	)
}

func (a MusicTrackRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Title, validation.Required),
		validation.Field(&a.Artist, validation.Required),
		validation.Field(&a.Album, validation.Required),
		validation.Field(&a.Genre, validation.Required),
		validation.Field(&a.ReleaseYear, validation.Required),
		validation.Field(&a.Duration, validation.Required),
		validation.Field(&a.Mp3File, validation.Required),
	)
}
