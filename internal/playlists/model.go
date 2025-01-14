package playlists

type Playlist struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Duration    string   `json:"duration"`
	CreatedAt   string `json:"created_at"`

	CreatorID string `json:"creator_id"`
}

func (m Playlist) GetID() string {
	return m.ID
}