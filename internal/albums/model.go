package albums

type Album struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Cover     string    `json:"cover"`
	CreatedAt string `json:"created_at"`
	OwnerID string `json:"owner_id"`
}

func (a Album) GetID() string {
	return a.ID
}