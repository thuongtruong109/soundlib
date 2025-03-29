package track_in_playlist

import (
	"fmt"

	"github.com/thuongtruong109/gouse"
	"github.com/thuongtruong109/soundlib/pkg/helpers"
)

type TrackInPlaylistUsecase struct {
	repo   TrackInPlaylistRepository
	helper helpers.Helper
}

func NewTrackInPlaylistUsecase(repo TrackInPlaylistRepository, helper helpers.Helper) *TrackInPlaylistUsecase {
	return &TrackInPlaylistUsecase{
		repo:   repo,
		helper: helper,
	}
}

func (g *TrackInPlaylistUsecase) GetTracksOfPlaylist() ([]string, string, error) {
	result, time, err := helpers.QueryTimeTwoOutput(g.repo.GetTracksOfPlaylist)()

	if err != nil {
		return nil, "", err
	}

	if result == nil {
		return nil, "", fmt.Errorf(gouse.DESC_NOT_FOUND_DATA)
	}

	var items []string
	for _, item := range result {
		items = append(items, fmt.Sprintf("TrackID: %s, PlaylistID: %s, AddedAt: %s", item.TrackID, item.PlaylistID, item.AddedAt))
	}

	return items, time, nil
}

func (g *TrackInPlaylistUsecase) AddTrackToPlaylist() ([]string, string, error) {
	var trackID string
	fmt.Print("» Enter track ID: ")
	fmt.Scanln(&trackID)

	var albumID string
	fmt.Print("» Enter playlist ID: ")
	fmt.Scanln(&trackID)

	var releasedAt string
	fmt.Print("» Enter added time: ")
	fmt.Scanln(&releasedAt)

	newTrack := &TrackInPlaylist{
		TrackID:    trackID,
		PlaylistID: albumID,
		AddedAt:    gouse.ISODate(),
	}

	result, time, err := helpers.QueryTimeTwoOutputWithParams(g.repo.AddTrackToAlbum)(newTrack)
	if err != nil {
		return nil, time, err
	}

	if result == nil {
		return nil, time, fmt.Errorf(gouse.DESC_NOT_FOUND_DATA)
	}

	newItem := []string{fmt.Sprintf("TrackID: %s, PlaylistID: %s, AddedAt: %s", result.TrackID, result.PlaylistID, result.AddedAt)}

	return newItem, time, err
}

func (g *TrackInPlaylistUsecase) DeleteTrackFromPlaylist() (string, error) {
	var trackId string
	fmt.Print("» Enter track id: ")
	fmt.Scanln(&trackId)

	var playlistId string
	fmt.Print("» Enter playlist id: ")
	fmt.Scanln(&playlistId)

	time, err := helpers.QueryTimeErrorWithTwoParams[error](g.repo.DeleteTrackFromAlbum)(trackId, playlistId)
	if err != nil {
		return "", fmt.Errorf(gouse.DESC_NOT_FOUND_DATA)
	}

	return time, err
}
