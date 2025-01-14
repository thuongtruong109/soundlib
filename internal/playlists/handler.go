package playlists

import (
	"github.com/thuongtruong109/soundlib/pkg/helpers"
)

type PlaylistHandler struct {
	uc     PlaylistUsecase
	helper helpers.Helper
}

func NewPlaylistHandler(uc PlaylistUsecase, helper helpers.Helper) *PlaylistHandler {
	return &PlaylistHandler{
		uc:     uc,
		helper: helper,
	}
}

func (h *PlaylistHandler) GetPlaylists() {
	// result, err := h.uc.GetPlaylists()
	// h.helper.ErrorWrapper(constants.gouse.DESC_GET_FAILED, err)
	// h.helper.SuccessDataWrapper(constants.gouse.DESC_GET_SUCCESS, result)
	h.helper.OutputSuccess("GetPlaylist")
}

func (h *PlaylistHandler) GetPlaylist() {
	h.helper.OutputSuccess("GetPlaylist")
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
