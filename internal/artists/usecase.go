package artists

import (
	"fmt"

	gu_helper "github.com/thuongtruong109/gouse/helper"
	"github.com/thuongtruong109/soundlib/pkg/constants"
	"github.com/thuongtruong109/soundlib/pkg/helpers"
)

type ArtistUsecase struct {
	repo   ArtistRepository
	helper helpers.Helper
}

func NewArtistUsecase(repo ArtistRepository, helper helpers.Helper) *ArtistUsecase {
	return &ArtistUsecase{
		repo:   repo,
		helper: helper,
	}
}

func (a *ArtistUsecase) GetArtists() ([]string, error) {
	result, err := helpers.QueryTimeTwoOutput[[]*Artist](a.repo.GetArtists)()
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, fmt.Errorf(constants.NOT_FOUND_DATA)
	}

	var output []string
	for _, v := range result {
		output = append(output, fmt.Sprintf("ID: %s, Name: %s", v.ID, v.Name))
	}

	return output, nil
}

func (a *ArtistUsecase) GetArtist() ([]string, error) {
	var id string
	fmt.Print("» Enter ID: ")
	fmt.Scanln(&id)

	result, err := helpers.QueryTimeTwoOutputWithParams[*Artist, string](a.repo.GetArtist)(id)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, fmt.Errorf(constants.NOT_FOUND_DATA)
	}

	output := []string{fmt.Sprintf("ID: %s, Name: %s", result.ID, result.Name)}

	return output, nil
}

func (a *ArtistUsecase) CreateArtist() ([]string, error) {
	var name string
	fmt.Print("» Enter name: ")
	fmt.Scanln(&name)

	artist := &Artist{
		ID:   gu_helper.RandomID(),
		Name: name,
	}

	result, err := helpers.QueryTimeTwoOutputWithParams[*Artist, *Artist](a.repo.CreateArtist)(artist)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, fmt.Errorf(constants.NOT_FOUND_DATA)
	}

	output := []string{fmt.Sprintf("ID: %s", result.ID)}

	return output, nil
}

func (a *ArtistUsecase) DeleteArtist() error {
	var id string
	fmt.Print("» Enter ID: ")
	fmt.Scanln(&id)

	err := helpers.QueryTimeOneOutputWithParams[error](a.repo.DeleteArtist)(id)
	if err != nil {
		return err
	}

	return nil
}

func (a *ArtistUsecase) UpdateArtist() ([]string, error) {
	var id string
	fmt.Print("» Enter ID: ")
	fmt.Scanln(&id)

	var name string
	fmt.Print("» Enter name: ")
	fmt.Scanln(&name)

	newArtist := &Artist{
		ID:   id,
		Name: name,
	}

	result, err := helpers.QueryTimeTwoOutputWithParams[*Artist, *Artist](a.repo.UpdateArtist)(newArtist)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, fmt.Errorf(constants.NOT_FOUND_DATA)
	}

	output := []string{fmt.Sprintf("ID: %s", result.ID)}

	return output, nil
}

func (a *ArtistUsecase) GetAlbumsOfArtist() string {
	return "Albums of Artist"
}

func (a *ArtistUsecase) GetTracksOfArtist() string {
	return "Tracks of Artist"
}
