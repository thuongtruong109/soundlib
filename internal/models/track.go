package models

import (
	"music-management/pkg/helpers"
)

type Track struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Year int `json:"year"`
	Duration int `json:"duration"`
	
	GenreID int `json:"genre_id"`
	AlbumID int `json:"album_id"`
	ArtistID int `json:"artist_id"`
}

func NewTrack(helper helpers.Helper) *Track {
	return &Track{
		ID: helper.GenerateID(),
	}
}