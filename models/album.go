package models

import (
	"music-management/pkg/helpers"
)

type Album struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Duration float32 `json:"duration"`
	Year int `json:"year"`

	GenreID int `json:"genre_id"`
	ArtistID int `json:"artist_id"`
}

func (a *Album) NewAlbum(title string, duration float32, year int, genreID int, artistID int) *Album {
	return &Album{
		ID: helpers.GenerateID(),
		Title: title,
		Duration: duration,
		Year: year,
		GenreID: genreID,
		ArtistID: artistID,
	}
}

