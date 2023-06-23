package models

import (
	"music-management/pkg/helpers"
)

type Album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Duration float32 `json:"duration"`
	Year int `json:"year"`

	GenreID int `json:"genre_id"`
	ArtistID int `json:"artist_id"`
}

func NewAlbum(helper helpers.Helper) *Album {
	return &Album{
		ID: helper.GenerateID(),
	}
}

