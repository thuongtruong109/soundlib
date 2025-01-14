package albums

import (
	"fmt"

	"github.com/thuongtruong109/soundlib/pkg/base"
)

type AlbumRepository struct {
	*base.Repository[Album]
}

func NewAlbumRepository(dbPath string) *AlbumRepository {
	return &AlbumRepository{
		Repository:  &base.Repository[Album]{DataPath: dbPath},
	}
}

func (ar *AlbumRepository) GetAlbums() ([]*Album, error) {
	albums, err := ar.Repository.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get all albums: %w", err)
	}

	var albumPointers []*Album
	for _, album := range albums {
		albumPointers = append(albumPointers, &album)
	}

	return albumPointers, nil
}

func (ar *AlbumRepository) GetAlbum(albumID string) (*Album, error) {
	return ar.Repository.GetByID(albumID)
}

func (ar *AlbumRepository) CreateAlbum(newAlbum *Album) (*Album, error) {
	if err := ar.Repository.Create(newAlbum); err != nil {
		return nil, fmt.Errorf("failed to create album: %w", err)
	}
	return newAlbum, nil
}

func (ar *AlbumRepository) UpdateAlbum(updateAlbum *Album) (*Album, error) {
	if err := ar.Repository.Update(updateAlbum); err != nil {
		return nil, fmt.Errorf("failed to update album: %w", err)
	}
	return updateAlbum, nil
}

func (ar *AlbumRepository) DeleteAlbum(albumID string) error {
	if err := ar.Repository.Delete(albumID); err != nil {
		return fmt.Errorf("failed to delete album: %w", err)
	}
	return nil
}