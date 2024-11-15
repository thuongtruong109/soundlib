package track_in_album

import (
	"time"

	"github.com/thuongtruong109/soundlib/internal/albums"
	"github.com/thuongtruong109/soundlib/internal/tracks"
)

type TrackInAlbum struct {
	TrackID    tracks.Track `json:"track_id"`
	AlbumID    albums.Album `json:"album_id"`
	ReleasedAt time.Time    `json:"released_at"`
}
