package handlers

import (
	"github.com/thuongtruong109/soundlib/pkg/helpers"

	"github.com/thuongtruong109/soundlib/internal/usecases"
)

type PlaylistHandler struct {
	uc usecases.PlaylistUsecase
	helper helpers.Helper
}

func NewPlaylistHandler(uc usecases.PlaylistUsecase, helper helpers.Helper) *PlaylistHandler {
	return &PlaylistHandler{
		uc: uc,
		helper: helper,
	}
}

func (u *PlaylistHandler) GetPlaylists() {
	u.helper.OutputSuccess("GetPlaylists")
}

func (u *PlaylistHandler) GetPlaylist() {
	u.helper.OutputSuccess("GetPlaylist")
}

func (u *PlaylistHandler) CreatePlaylist() {
	u.helper.OutputSuccess("CreatePlaylist")
}

func (u *PlaylistHandler) DeletePlaylist() {
	u.helper.OutputSuccess("DeletePlaylist")
}

func (u *PlaylistHandler) UpdatePlaylist() {
	u.helper.OutputSuccess("UpdatePlaylist")
}

func (u *PlaylistHandler) GetTracksOfPlaylist() {
	u.helper.OutputSuccess("GetTracksOfPlaylist")
}

func (u *PlaylistHandler) AddTrackToPlaylist() {
	u.helper.OutputSuccess("AddTrackToPlaylist")
}

func (u *PlaylistHandler) DeleteTrackFromPlaylist() {
	u.helper.OutputSuccess("DeleteTrackFromPlaylist")
}

func (u *PlaylistHandler) GetPlaylistsHaveTrack() {
	u.helper.OutputSuccess("GetPlaylistsHaveTrack")
}