package artists

import (
	"fmt"

	"github.com/thuongtruong109/gouse"
	"github.com/thuongtruong109/soundlib/pkg/constants"
)

type ArtistRepository struct{}

func NewArtistRepository() *ArtistRepository {
	return &ArtistRepository{}
}

func (ar *ArtistRepository) GetArtists() ([]*Artist, error) {
	allArtist, err := gouse.ReadFileObj[*Artist](constants.ARTIST_PATH)
	if err != nil {
		return nil, err
	}

	if allArtist == nil {
		return nil, nil
	}

	return allArtist, nil
}

func (ar *ArtistRepository) GetArtist(artistID string) (*Artist, error) {
	allArtist, err := ar.GetArtists()
	if err != nil {
		return nil, err
	}

	if allArtist == nil {
		return nil, nil
	}

	for _, v := range allArtist {
		if v.ID == artistID {
			return v, nil
		}
	}

	return nil, nil
}

func (ar *ArtistRepository) CreateArtist(newArtist *Artist) (*Artist, error) {
	allGenres, _ := ar.GetArtists()

	var artistInit []*Artist

	if allGenres == nil {
		artistInit = make([]*Artist, 0)
	} else {
		artistInit = make([]*Artist, len(allGenres))
		copy(artistInit, allGenres)
	}

	artistInit = append(artistInit, newArtist)

	err2 := gouse.WriteFileObj[[]*Artist](constants.ARTIST_PATH, artistInit)
	if err2 != nil {
		return nil, fmt.Errorf(gouse.DESC_CREATE_FAILED)
	}
	return newArtist, nil
}

func (ar *ArtistRepository) UpdateArtist(artistUpdate *Artist) (*Artist, error) {
	allArtist, _ := ar.GetArtists()

	if allArtist == nil {
		return nil, fmt.Errorf(gouse.DESC_UPDATE_FAILED)
	}

	var artistInit []*Artist

	for i, v := range allArtist {
		if v.ID == artistUpdate.ID {
			artistInit = append(allArtist[:i], artistUpdate)
			break
		}
	}

	artistInit = append(artistInit, allArtist[len(artistInit):]...)

	err2 := gouse.WriteFileObj[[]*Artist](constants.ARTIST_PATH, artistInit)
	if err2 != nil {
		return nil, fmt.Errorf(gouse.DESC_UPDATE_FAILED)
	}
	return artistUpdate, nil
}

func (ar *ArtistRepository) DeleteArtist(artlistID string) error {
	allArtist, _ := ar.GetArtists()

	if allArtist == nil {
		return fmt.Errorf(gouse.DESC_DELETE_FAILED)
	}

	for i, v := range allArtist {
		if v.ID == artlistID {
			allArtist = append(allArtist[:i], allArtist[i+1:]...)
			break
		}
	}

	err2 := gouse.WriteFileObj[[]*Artist](constants.ARTIST_PATH, allArtist)
	if err2 != nil {
		return fmt.Errorf(gouse.DESC_DELETE_FAILED)
	}
	return nil
}
