package track_in_playlist

import (
	"github.com/thuongtruong109/gouse"
	"github.com/thuongtruong109/soundlib/pkg/common"
	"github.com/thuongtruong109/soundlib/pkg/helpers"
)

type TrackInPlaylistHandler struct {
	uc     TrackInPlaylistUsecase
	helper helpers.Helper
	common common.CommonHandler
}

func NewTrackInPlaylistHandler(uc TrackInPlaylistUsecase, helper helpers.Helper, common common.CommonHandler) *TrackInPlaylistHandler {
	return &TrackInPlaylistHandler{
		uc:     uc,
		helper: helper,
		common: common,
	}
}

func (h *TrackInPlaylistHandler) AddTrackToPlaylist() {
	result, time, err := h.uc.AddTrackToPlaylist()
	h.common.ErrorWrapper(gouse.DESC_GET_FAILED, err)
	h.common.SuccessDataWrapper(gouse.DESC_GET_SUCCESS, result, time)
}

func (h *TrackInPlaylistHandler) GetTracksOfPlaylist() {
	result, time, err := h.uc.GetTracksOfPlaylist()
	h.common.ErrorWrapper(gouse.DESC_GET_FAILED, err)
	h.common.SuccessDataWrapper(gouse.DESC_GET_SUCCESS, result, time)
}

func (h *TrackInPlaylistHandler) DeleteTrackFromPlaylist() {
	time, err := h.uc.DeleteTrackFromPlaylist()
	h.common.ErrorWrapper(gouse.DESC_GET_FAILED, err)
	h.common.SuccessNoDataWrapper(gouse.DESC_GET_SUCCESS, time)
}
