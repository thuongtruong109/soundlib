package models

type Album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Duration float32 `json:"duration"`
	Year int `json:"year"`

	GenreID int `json:"genre_id"`
	ArtistID int `json:"artist_id"`
}

