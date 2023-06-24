package usecases

import (
	"fmt"
	"music-management/database"
	"music-management/pkg/constants"
	"music-management/internal/models"
	"music-management/pkg/helpers"
)

type ArtistUsecase struct {
	helper helpers.Helper
}

func NewArtistUsecase(helper helpers.Helper) *ArtistUsecase {
	return &ArtistUsecase{
		helper: helper,
	}
}

func (a *ArtistUsecase) GetArtists() ([]*models.Artist, error) {
	allArtist, err := database.ReadDB[*models.Artist](constants.ARTIST_PATH)
	if err != nil {
		return nil, err
	}

	if allArtist == nil {
		return nil, nil
	}

	artists := make([]*models.Artist, len(allArtist))
	copy(artists, allArtist)

	return artists, nil
}

func (a *ArtistUsecase) GetArtist() (*models.Artist, error) {
	allArtist, err := a.GetArtists()
	if err != nil {
		return nil, err
	}

	if allArtist == nil {
		return nil, nil
	}

	var id string
	fmt.Print("» Enter ID: ")
	fmt.Scanln(&id)

	for _, v := range allArtist {
		if v.ID == id {
			return v, nil
		}
	}

	return nil, nil
}

func (a *ArtistUsecase) CreateArtist() (*models.Artist, error) {
	allGenres, _ := a.GetArtists()

	var artistInit []*models.Artist

	if allGenres == nil {
		artistInit = make([]*models.Artist, 0)
	} else {
		artistInit = make([]*models.Artist, len(allGenres))
		copy(artistInit, allGenres)
	}

	var name string
	fmt.Print("» Enter name: ")
	fmt.Scanln(&name)

	artist := &models.Artist{
		ID: a.helper.GenerateID(),
		Name: name,
	}

	artistInit = append(artistInit, artist)

	err2 := database.SaveDB[[]*models.Artist](constants.ARTIST_PATH, artistInit)
	if err2 != nil {
		return nil, fmt.Errorf(constants.CREATE_FAILED)
	}
	return artist, nil
}

func (a *ArtistUsecase) DeleteArtist() error {
	allArtist, _ := a.GetArtists()

	if allArtist == nil {
		return fmt.Errorf(constants.DELETE_FAILED)
	}

	var id string
	fmt.Print("» Enter ID: ")
	fmt.Scanln(&id)

	for i, v := range allArtist {
		if v.ID == id {
			allArtist = append(allArtist[:i], allArtist[i+1:]...)
			break
		}
	}

	allArtist = append(allArtist, allArtist[len(allArtist)-1])

	err2 := database.SaveDB[[]*models.Artist](constants.ARTIST_PATH, allArtist)
	if err2 != nil {
		return fmt.Errorf(constants.DELETE_FAILED)
	}
	return nil
}

func (a *ArtistUsecase) UpdateArtist() (*models.Artist, error) {
	allArtist, _ := a.GetArtists()

	if allArtist == nil {
		return nil, fmt.Errorf(constants.UPDATE_FAILED)
	}

	var id string
	fmt.Print("» Enter ID: ")
	fmt.Scanln(&id)
	
	var name string
	fmt.Print("» Enter name: ")
	fmt.Scanln(&name)

	newArtist := &models.Artist{
		ID: id,
		Name: name,
	}

	var artistInit []*models.Artist

	for i, v := range allArtist {
		if v.ID == id {
			artistInit = append(allArtist[:i], newArtist)
			break
		}
	}

	artistInit = append(artistInit, allArtist[len(artistInit):]...)

	err2 := database.SaveDB[[]*models.Artist](constants.ARTIST_PATH, allArtist)
	if err2 != nil {
		return nil, fmt.Errorf(constants.UPDATE_FAILED)
	}
	return newArtist, nil
}

func (a *ArtistUsecase) GetAlbumsOfArtist() string {
	return "Albums of Artist"
}

func (a *ArtistUsecase) GetTracksOfArtist() string {
	return "Tracks of Artist"
}