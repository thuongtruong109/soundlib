package track_in_album

import (
	"github.com/thuongtruong109/gouse"
	"github.com/thuongtruong109/soundlib/pkg/common"
	"github.com/thuongtruong109/soundlib/pkg/helpers"
)

type TrackInAlbumHandler struct {
	uc     TrackInAlbumUsecase
	helper helpers.Helper
	common common.CommonHandler
}

func NewTrackInAlbumHandler(uc TrackInAlbumUsecase, helper helpers.Helper, common common.CommonHandler) *TrackInAlbumHandler {
	return &TrackInAlbumHandler{
		uc:     uc,
		helper: helper,
		common: common,
	}
}

func (h *TrackInAlbumHandler) AddTrackToAlbum() {
	result, time, err := h.uc.AddTrackToAlbum()
	h.common.ErrorWrapper(gouse.DESC_GET_FAILED, err)
	h.common.SuccessDataWrapper(gouse.DESC_GET_SUCCESS, result, time)
}

func (h *TrackInAlbumHandler) GetTracksOfAlbum() {
	result, time, err := h.uc.GetTracksOfAlbum()
	h.common.ErrorWrapper(gouse.DESC_GET_FAILED, err)
	h.common.SuccessDataWrapper(gouse.DESC_GET_SUCCESS, result, time)
}

func (h *TrackInAlbumHandler) DeleteTrackFromAlbum() {
	time, err := h.uc.DeleteTrackFromAlbum()
	h.common.ErrorWrapper(gouse.DESC_GET_FAILED, err)
	h.common.SuccessNoDataWrapper(gouse.DESC_GET_SUCCESS, time)
}
