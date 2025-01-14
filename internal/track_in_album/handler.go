package track_in_album

import (
	"github.com/thuongtruong109/gouse"
	"github.com/thuongtruong109/soundlib/pkg/base"
	"github.com/thuongtruong109/soundlib/pkg/helpers"
)

type TrackInAlbumHandler struct {
	uc     TrackInAlbumUsecase
	helper helpers.Helper
	base base.BaseHandler
}

func NewTrackInAlbumHandler(uc TrackInAlbumUsecase, helper helpers.Helper, base base.BaseHandler) *TrackInAlbumHandler {
	return &TrackInAlbumHandler{
		uc:     uc,
		helper: helper,
		base: base,
	}
}

func (h *TrackInAlbumHandler) AddTrackToAlbum() {
	result, time, err := h.uc.AddTrackToAlbum()
	h.base.ErrorWrapper(gouse.DESC_GET_FAILED, err)
	h.base.SuccessDataWrapper(gouse.DESC_GET_SUCCESS, result, time)
}

func (h *TrackInAlbumHandler) GetTracksOfAlbum() {
	result, time, err := h.uc.GetTracksOfAlbum()
	h.base.ErrorWrapper(gouse.DESC_GET_FAILED, err)
	h.base.SuccessDataWrapper(gouse.DESC_GET_SUCCESS, result, time)
}

func (h *TrackInAlbumHandler) DeleteTrackFromAlbum() {
	time, err := h.uc.DeleteTrackFromAlbum()
	h.base.ErrorWrapper(gouse.DESC_GET_FAILED, err)
	h.base.SuccessNoDataWrapper(gouse.DESC_GET_SUCCESS, time)
}
