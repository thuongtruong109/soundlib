package tracks

import (
	"fmt"

	gu_helper "github.com/thuongtruong109/gouse/helper"
	"github.com/thuongtruong109/soundlib/pkg/constants"
	"github.com/thuongtruong109/soundlib/pkg/helpers"
)

type TrackUsecase struct {
	repo   TrackRepository
	helper helpers.Helper
}

func NewTrackUsecase(repo TrackRepository, helper helpers.Helper) *TrackUsecase {
	return &TrackUsecase{
		repo:   repo,
		helper: helper,
	}
}

func (a *TrackUsecase) GetTracks() ([]string, error) {
	result, err := helpers.QueryTimeTwoOutput[[]*Track](a.repo.GetTracks)()
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, fmt.Errorf(constants.NOT_FOUND_DATA)
	}

	var output []string
	for _, v := range result {
		output = append(output, fmt.Sprintf("ID: %s, Name: %s", v.ID, v.Title))
	}

	return output, nil
}

func (a *TrackUsecase) GetTrack() ([]string, error) {
	var id string
	fmt.Print("» Enter ID: ")
	fmt.Scanln(&id)

	result, err := helpers.QueryTimeTwoOutputWithParams[*Track, string](a.repo.GetTrackById)(id)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, fmt.Errorf(constants.NOT_FOUND_DATA)
	}

	output := []string{fmt.Sprintf("ID: %s, Name: %s", result.ID, result.Title)}

	return output, nil
}

func (a *TrackUsecase) CreateTrack() ([]string, error) {
	var name string
	fmt.Print("» Enter track title: ")
	fmt.Scanln(&name)

	newTrack := &Track{
		ID:    gu_helper.RandomID(),
		Title: name,
	}

	result, err := helpers.QueryTimeTwoOutputWithParams[*Track, *Track](a.repo.CreateTrack)(newTrack)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, fmt.Errorf(constants.NOT_FOUND_DATA)
	}

	output := []string{fmt.Sprintf("ID: %s", result.ID)}

	return output, nil
}

func (a *TrackUsecase) DeleteTrack() error {
	var id string
	fmt.Print("» Enter ID: ")
	fmt.Scanln(&id)

	err := helpers.QueryTimeOneOutputWithParams[error](a.repo.DeleteTrack)(id)
	if err != nil {
		return err
	}

	return nil
}

func (a *TrackUsecase) UpdateTrack() ([]string, error) {
	var id string
	fmt.Print("» Enter ID: ")
	fmt.Scanln(&id)

	var name string
	fmt.Print("» Enter name: ")
	fmt.Scanln(&name)

	newTrack := &Track{
		ID:    id,
		Title: name,
	}

	result, err := helpers.QueryTimeTwoOutputWithParams[*Track, *Track](a.repo.UpdateTrack)(newTrack)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, fmt.Errorf(constants.NOT_FOUND_DATA)
	}

	output := []string{fmt.Sprintf("ID: %s", result.ID)}

	return output, nil
}

func (a *TrackUsecase) GetAlbumsOfTrack() string {
	return "Albums of Track"
}

func (a *TrackUsecase) GetTracksOfTrack() string {
	return "Tracks of Track"
}
