package albums

import (
	"time"

	"github.com/thuongtruong109/soundlib/internal/artists"
)

type Album struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Cover     string    `json:"cover"`
	CreatedAt time.Time `json:"created_at"`

	OwnerID artists.Artist `json:"owner_id"`
}
