package artists

import (
	"fmt"

	gu_date "github.com/thuongtruong109/gouse/date"
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

func (a *ArtistUsecase) GetArtists() ([]string, string, error) {
	result, time, err := helpers.QueryTimeTwoOutput[[]*Artist](a.repo.GetArtists)()
	if err != nil {
		return nil, time, err
	}

	if result == nil {
		return nil, time, fmt.Errorf(constants.NOT_FOUND_DATA)
	}

	var output []string
	for _, v := range result {
		output = append(output, fmt.Sprintf("ID: %s, Username: %s, Full Name: %s, Bio: %s, Avatar URL: %s, Debut At: %s", v.ID, v.Username, v.FullName, v.Bio, v.AvatarUrl, v.DebutAt))
	}

	return output, time, nil
}

func (a *ArtistUsecase) GetArtist() ([]string, string, error) {
	var id string
	fmt.Print("» Enter ID: ")
	fmt.Scanln(&id)

	result, time, err := helpers.QueryTimeTwoOutputWithParams[*Artist, string](a.repo.GetArtist)(id)
	if err != nil {
		return nil, time, err
	}

	if result == nil {
		return nil, time, fmt.Errorf(constants.NOT_FOUND_DATA)
	}

	output := []string{fmt.Sprintf("ID: %s, Username: %s, Full Name: %s, Bio: %s, Avatar URL: %s, Debut At: %s", result.ID, result.Username, result.FullName, result.Bio, result.AvatarUrl, result.DebutAt)}

	return output, time, nil
}

func (a *ArtistUsecase) CreateArtist() ([]string, string, error) {
	var username, fullName, bio, avatarUrl string
	fmt.Print("» Enter username: ")
	fmt.Scanln(&username)

	fmt.Print("» Enter full name: ")
	fmt.Scanln(&fullName)

	fmt.Print("» Enter bio: ")
	fmt.Scanln(&bio)

	fmt.Print("» Enter avatar URL: ")
	fmt.Scanln(&avatarUrl)

	artist := &Artist{
		ID:        gu_helper.RandomID(),
		Username:  username,
		FullName:  fullName,
		Bio:       bio,
		AvatarUrl: avatarUrl,
		DebutAt:   gu_date.ISO(),
	}

	result, time, err := helpers.QueryTimeTwoOutputWithParams[*Artist, *Artist](a.repo.CreateArtist)(artist)
	if err != nil {
		return nil, time, err
	}

	if result == nil {
		return nil, time, fmt.Errorf(constants.NOT_FOUND_DATA)
	}

	output := []string{fmt.Sprintf("ID: %s", result.ID)}

	return output, time, nil
}

func (a *ArtistUsecase) DeleteArtist() (string, error) {
	var id string
	fmt.Print("» Enter ID: ")
	fmt.Scanln(&id)

	time, err := helpers.QueryTimeErrorWithParams[error](a.repo.DeleteArtist)(id)
	if err != nil {
		return time, err
	}

	return time, nil
}

func (a *ArtistUsecase) UpdateArtist() ([]string, string, error) {
	var id, username, fullName, bio, avatarUrl string
	fmt.Print("» Enter ID: ")
	fmt.Scanln(&id)

	fmt.Print("» Enter username: ")
	fmt.Scanln(&username)

	fmt.Print("» Enter full name: ")
	fmt.Scanln(&fullName)

	fmt.Print("» Enter bio: ")
	fmt.Scanln(&bio)

	fmt.Print("» Enter avatar URL: ")
	fmt.Scanln(&avatarUrl)

	newArtist := &Artist{
		ID:        id,
		Username:  username,
		FullName:  fullName,
		Bio:       bio,
		AvatarUrl: avatarUrl,
		DebutAt:   gu_date.ISO(),
	}

	result, time, err := helpers.QueryTimeTwoOutputWithParams[*Artist, *Artist](a.repo.UpdateArtist)(newArtist)
	if err != nil {
		return nil, time, err
	}

	if result == nil {
		return nil, time, fmt.Errorf(constants.NOT_FOUND_DATA)
	}

	output := []string{fmt.Sprintf("ID: %s", result.ID)}

	return output, time, nil
}

func (a *ArtistUsecase) GetTrackOfArtist() string {
	return "Tracks of Artist"
}

func (a *ArtistUsecase) GetTracksOfGenre() string {
	return "Tracks of Genre"
}
