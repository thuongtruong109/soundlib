package track_in_playlist

import (
	"github.com/thuongtruong109/gouse"
	"github.com/thuongtruong109/soundlib/pkg/base"
	"github.com/thuongtruong109/soundlib/pkg/helpers"
)

type TrackInPlaylistHandler struct {
	uc     TrackInPlaylistUsecase
	helper helpers.Helper
	base base.BaseHandler
}

func NewTrackInPlaylistHandler(uc TrackInPlaylistUsecase, helper helpers.Helper, base base.BaseHandler) *TrackInPlaylistHandler {
	return &TrackInPlaylistHandler{
		uc:     uc,
		helper: helper,
		base: base,
	}
}

func (h *TrackInPlaylistHandler) AddTrackToPlaylist() {
	result, time, err := h.uc.AddTrackToPlaylist()
	h.base.ErrorWrapper(gouse.DESC_GET_FAILED, err)
	h.base.SuccessDataWrapper(gouse.DESC_GET_SUCCESS, result, time)
}

func (h *TrackInPlaylistHandler) GetTracksOfPlaylist() {
	result, time, err := h.uc.GetTracksOfPlaylist()
	h.base.ErrorWrapper(gouse.DESC_GET_FAILED, err)
	h.base.SuccessDataWrapper(gouse.DESC_GET_SUCCESS, result, time)
}

func (h *TrackInPlaylistHandler) DeleteTrackFromPlaylist() {
	time, err := h.uc.DeleteTrackFromPlaylist()
	h.base.ErrorWrapper(gouse.DESC_GET_FAILED, err)
	h.base.SuccessNoDataWrapper(gouse.DESC_GET_SUCCESS, time)
}
