package models

import (
	"music-management/pkg/helpers"
)

type Track struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Year int `json:"year"`
	Duration int `json:"duration"`
	
	GenreID int `json:"genre_id"`
	AlbumID int `json:"album_id"`
	ArtistID int `json:"artist_id"`
}

func NewTrack(title string, year, duration, genreID, albumID, artistID int) *Track {
	return &Track{
		ID: helpers.GenerateID(),
		Title: title,
		Year: year,
		Duration: duration,
		ArtistID: artistID,
		AlbumID: albumID,
		GenreID: genreID,
	}
}