package track_in_album

type TrackInAlbum struct {
	TrackID    string `json:"track_id"`
	AlbumID    string `json:"album_id"`
	ReleasedAt string `json:"released_at"`
}
