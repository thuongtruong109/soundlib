package track_in_playlist

type TrackInPlaylist struct {
	TrackID    string `json:"track_id"`
	PlaylistID string `json:"playlist_id"`
	AddedAt    string `json:"added_at"`
}
