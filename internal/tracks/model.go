package tracks

type Track struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	PlayCount int     `json:"play_count"`
	Duration  float32 `json:"duration"`
	FileUrl   string  `json:"file_url"`
	CreatedAt string  `json:"created_at"`

	GenreID  string `json:"genre_id"`
	ArtistID string `json:"artist_id"`
}
