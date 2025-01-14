package playlists

import (
	"fmt"

	"github.com/thuongtruong109/soundlib/pkg/base"
)

type PlaylistRepository struct {
	*base.Repository[Playlist]
}

func NewPlaylistRepository(dbPath string) *PlaylistRepository {
	return &PlaylistRepository{
		Repository:  &base.Repository[Playlist]{DataPath: dbPath},
	}
}

func (ar *PlaylistRepository) GetPlaylists() ([]*Playlist, error) {
	playlists, err := ar.Repository.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get all playlists: %w", err)
	}

	var playlistPointers []*Playlist
	for _, playlist := range playlists {
		playlistPointers = append(playlistPointers, &playlist)
	}

	return playlistPointers, nil
}

func (ar *PlaylistRepository) GetPlaylist(playlistID string) (*Playlist, error) {
	return ar.Repository.GetByID(playlistID)
}

func (ar *PlaylistRepository) CreatePlaylist(newPlaylist *Playlist) (*Playlist, error) {
	if err := ar.Repository.Create(newPlaylist); err != nil {
		return nil, fmt.Errorf("failed to create playlist: %w", err)
	}
	return newPlaylist, nil
}

func (ar *PlaylistRepository) UpdatePlaylist(updatePlaylist *Playlist) (*Playlist, error) {
	if err := ar.Repository.Update(updatePlaylist); err != nil {
		return nil, fmt.Errorf("failed to update playlist: %w", err)
	}
	return updatePlaylist, nil
}

func (ar *PlaylistRepository) DeletePlaylist(playlistID string) error {
	if err := ar.Repository.Delete(playlistID); err != nil {
		return fmt.Errorf("failed to delete playlist: %w", err)
	}
	return nil
}