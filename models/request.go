package models

import validation "github.com/go-ozzo/ozzo-validation"

type MusicTrackRequest struct {
	Title       string `json:"title"`
	Artist      string `json:"artist"`
	Album       string `json:"album"`
	Genre       string `json:"genre"`
	ReleaseYear int    `json:"release_year"`
	Duration    int    `json:"duration"`
	Mp3File     string `json:"mp3_file"`
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
