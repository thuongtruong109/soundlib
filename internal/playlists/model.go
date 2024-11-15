package playlists

import (
	"time"

	"github.com/thuongtruong109/soundlib/internal/artists"
)

type Playlist struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Duration    float32   `json:"duration"`
	CreatedAt   time.Time `json:"created_at"`

	CreatorID artists.Artist `json:"creator_id"`
}
