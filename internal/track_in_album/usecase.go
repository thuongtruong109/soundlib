package track_in_album

import (
	"fmt"

	"github.com/thuongtruong109/gouse"
	"github.com/thuongtruong109/soundlib/pkg/helpers"
)

type TrackInAlbumUsecase struct {
	repo   TrackInAlbumRepository
	helper helpers.Helper
}

func NewTrackInAlbumUsecase(repo TrackInAlbumRepository, helper helpers.Helper) *TrackInAlbumUsecase {
	return &TrackInAlbumUsecase{
		repo:   repo,
		helper: helper,
	}
}

func (g *TrackInAlbumUsecase) GetTracksOfAlbum() ([]string, string, error) {
	result, time, err := helpers.QueryTimeTwoOutput[[]*TrackInAlbum](g.repo.GetTracksOfAlbum)()

	if err != nil {
		return nil, "", err
	}

	if result == nil {
		return nil, "", fmt.Errorf(gouse.DESC_NOT_FOUND_DATA)
	}

	var items []string
	for _, item := range result {
		items = append(items, fmt.Sprintf("TrackID: %s, AlbumID: %s, ReleasedAt: %s", item.TrackID, item.AlbumID, item.ReleasedAt))
	}

	return items, time, nil
}

func (g *TrackInAlbumUsecase) AddTrackToAlbum() ([]string, string, error) {
	var trackID string
	fmt.Print("» Enter track ID: ")
	fmt.Scanln(&trackID)

	var albumID string
	fmt.Print("» Enter album ID: ")
	fmt.Scanln(&trackID)

	var releasedAt string
	fmt.Print("» Enter released time: ")
	fmt.Scanln(&releasedAt)

	newTrack := &TrackInAlbum{
		TrackID:    trackID,
		AlbumID:    albumID,
		ReleasedAt: gouse.ISODate(),
	}

	result, time, err := helpers.QueryTimeTwoOutputWithParams[*TrackInAlbum, *TrackInAlbum](g.repo.AddTrackToAlbum)(newTrack)
	if err != nil {
		return nil, time, err
	}

	if result == nil {
		return nil, time, fmt.Errorf(gouse.DESC_NOT_FOUND_DATA)
	}

	newItem := []string{fmt.Sprintf("TrackID: %s, AlbumID: %s, ReleasedAt: %s", result.TrackID, result.AlbumID, result.ReleasedAt)}

	return newItem, time, err
}

func (g *TrackInAlbumUsecase) DeleteTrackFromAlbum() (string, error) {
	var trackId string
	fmt.Print("» Enter track id: ")
	fmt.Scanln(&trackId)

	var albumId string
	fmt.Print("» Enter album id: ")
	fmt.Scanln(&albumId)

	time, err := helpers.QueryTimeErrorWithTwoParams[error](g.repo.DeleteTrackFromAlbum)(trackId, albumId)
	if err != nil {
		return "", fmt.Errorf(gouse.DESC_NOT_FOUND_DATA)
	}

	return time, err
}
