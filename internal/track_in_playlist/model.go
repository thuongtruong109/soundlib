package track_in_playlist

import (
	"time"

	"github.com/thuongtruong109/soundlib/internal/playlists"
	"github.com/thuongtruong109/soundlib/internal/tracks"
)

type TrackInPlaylist struct {
	TrackID    tracks.Track       `json:"track_id"`
	PlaylistID playlists.Playlist `json:"playlist_id"`
	AddedAt    time.Time          `json:"added_at"`
}
