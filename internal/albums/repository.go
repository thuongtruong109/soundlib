package albums

import (
	"fmt"

	"github.com/thuongtruong109/gouse"
	"github.com/thuongtruong109/soundlib/pkg/constants"
)

type AlbumRepository struct{}

func NewAlbumRepository() *AlbumRepository {
	return &AlbumRepository{}
}

func (gr *AlbumRepository) GetAlbums() ([]*Album, error) {
	albums, err := gouse.ReadFileObj[*Album](constants.ALBUM_PATH)

	if err != nil {
		return nil, err
	}

	if albums == nil {
		return nil, nil
	}

	return albums, nil
}

func (ar *AlbumRepository) GetArtist(albumID string) (*Album, error) {
	albums, err := ar.GetAlbums()
	if err != nil {
		return nil, err
	}

	if albums == nil {
		return nil, nil
	}

	for _, v := range albums {
		if v.ID == albumID {
			return v, nil
		}
	}

	return nil, nil
}

func (gr *AlbumRepository) CreateAlbum(newGenre *Album) (*Album, error) {
	albums, _ := gr.GetAlbums()

	var albumInit []*Album

	if albums == nil {
		albumInit = make([]*Album, 0)
	} else {
		albumInit = make([]*Album, len(albums))
		copy(albumInit, albums)
	}

	albumInit = append(albumInit, newGenre)

	err2 := gouse.WriteFileObj(constants.ALBUM_PATH, albumInit)
	if err2 != nil {
		return nil, fmt.Errorf(gouse.DESC_CREATE_FAILED)
	}
	return newGenre, nil
}

func (gr *AlbumRepository) UpdateAlbum(updateGenre *Album) (*Album, error) {
	albums, err := gr.GetAlbums()
	if err != nil {
		return nil, err
	}

	if albums == nil {
		return nil, nil
	}

	var albumInit []*Album

	for i, genre := range albums {
		if genre.ID == updateGenre.ID {
			albumInit = append(albums[:i], updateGenre)
			break
		}
	}

	albumInit = append(albumInit, albums[len(albumInit):]...)

	err2 := gouse.WriteFileObj[[]*Album](constants.ALBUM_PATH, albumInit)
	if err2 != nil {
		return nil, fmt.Errorf(gouse.DESC_UPDATE_FAILED)
	}
	return updateGenre, nil
}

func (gr *AlbumRepository) DeleteAlbum(genreID string) error {
	albums, _ := gr.GetAlbums()

	if albums == nil {
		return nil
	}

	var albumInit []*Album

	for i, genre := range albums {
		if genre.ID == genreID {
			albumInit = append(albums[:i], albums[i+1:]...)
			break
		}
	}

	err2 := gouse.WriteFileObj[[]*Album](constants.ALBUM_PATH, albumInit)
	if err2 != nil {
		return fmt.Errorf(gouse.DESC_DELETE_FAILED)
	}

	return nil
}
