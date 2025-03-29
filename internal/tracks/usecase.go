package tracks

import (
	"fmt"

	"github.com/thuongtruong109/gouse"
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

func (a *TrackUsecase) GetTracks() ([]string, string, error) {
	result, time, err := helpers.QueryTimeTwoOutput(a.repo.GetTracks)()
	if err != nil {
		return nil, time, err
	}

	if result == nil {
		return nil, "", fmt.Errorf(gouse.DESC_NOT_FOUND_DATA)
	}

	var output []string
	for _, v := range result {
		output = append(output, fmt.Sprintf("ID: %s, Name: %s, Play Count: %d, Duration: %f, File URL: %s, Created At: %s", v.ID, v.Name, v.PlayCount, v.Duration, v.FileUrl, v.CreatedAt))
	}

	return output, time, nil
}

func (a *TrackUsecase) GetTrack() ([]string, string, error) {
	var id string
	fmt.Print("» Enter ID: ")
	fmt.Scanln(&id)

	result, time, err := helpers.QueryTimeTwoOutputWithParams(a.repo.GetTrackById)(id)
	if err != nil {
		return nil, time, err
	}

	if result == nil {
		return nil, time, fmt.Errorf(gouse.DESC_NOT_FOUND_DATA)
	}

	output := []string{fmt.Sprintf("ID: %s, Name: %s, Play Count: %d, Duration: %f, File URL: %s, Created At: %s", result.ID, result.Name, result.PlayCount, result.Duration, result.FileUrl, result.CreatedAt)}

	return output, time, nil
}

func (a *TrackUsecase) CreateTrack() ([]string, string, error) {
	var name, genreID, artistID, fileURL string
	fmt.Print("» Enter track title: ")
	fmt.Scanln(&name)

	fmt.Print("» Enter genre ID: ")
	fmt.Scanln(&genreID)

	fmt.Print("» Enter artist ID: ")
	fmt.Scanln(&artistID)

	fmt.Print("» Enter file URL: ")
	fmt.Scanln(&fileURL)

	var duration float32
	fmt.Print("» Enter duration: ")
	fmt.Scanln(&duration)

	var playCount int
	fmt.Print("» Enter play count: ")
	fmt.Scanln(&playCount)

	newTrack := &Track{
		ID:        gouse.RandID(),
		Name:      name,
		GenreID:   genreID,
		ArtistID:  artistID,
		FileUrl:   fileURL,
		Duration:  duration,
		PlayCount: playCount,
		CreatedAt: gouse.ISODate(),
	}

	result, time, err := helpers.QueryTimeTwoOutputWithParams(a.repo.CreateTrack)(newTrack)
	if err != nil {
		return nil, time, err
	}

	if result == nil {
		return nil, time, fmt.Errorf(gouse.DESC_NOT_FOUND_DATA)
	}

	output := []string{fmt.Sprintf("ID: %s", result.ID)}

	return output, time, nil
}

func (a *TrackUsecase) DeleteTrack() (string, error) {
	var id string
	fmt.Print("» Enter ID: ")
	fmt.Scanln(&id)

	time, err := helpers.QueryTimeErrorWithOneParam[error](a.repo.DeleteTrack)(id)
	if err != nil {
		return "", err
	}

	return time, nil
}

func (a *TrackUsecase) UpdateTrack() ([]string, string, error) {
	var id string
	fmt.Print("» Enter ID: ")
	fmt.Scanln(&id)

	var name, genreID, artistID, fileURL string
	fmt.Print("» Enter track name: ")
	fmt.Scanln(&name)

	fmt.Print("» Enter genre ID: ")
	fmt.Scanln(&genreID)

	fmt.Print("» Enter artist ID: ")
	fmt.Scanln(&artistID)

	fmt.Print("» Enter file URL: ")
	fmt.Scanln(&fileURL)

	var duration float32
	fmt.Print("» Enter duration: ")
	fmt.Scanln(&duration)

	var playCount int
	fmt.Print("» Enter play count: ")
	fmt.Scanln(&playCount)

	newTrack := &Track{
		ID:        id,
		Name:      name,
		GenreID:   genreID,
		ArtistID:  artistID,
		FileUrl:   fileURL,
		Duration:  duration,
		PlayCount: playCount,
		CreatedAt: gouse.ISODate(),
	}

	result, time, err := helpers.QueryTimeTwoOutputWithParams(a.repo.UpdateTrack)(newTrack)
	if err != nil {
		return nil, time, err
	}

	if result == nil {
		return nil, time, fmt.Errorf(gouse.DESC_NOT_FOUND_DATA)
	}

	output := []string{fmt.Sprintf("ID: %s", result.ID)}

	return output, time, nil
}

// func (a *TrackUsecase) GetTracksOfGenre() string {
// 	return "Tracks of Genre"
// }

// func (a *TrackUsecase) GetTracksOfArtist() string {
// 	return "Tracks of Artist"
// }
