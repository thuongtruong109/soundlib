package tracks

import (
	"github.com/thuongtruong109/gouse"
	"github.com/thuongtruong109/soundlib/pkg/base"
	"github.com/thuongtruong109/soundlib/pkg/helpers"
)

type TrackHandler struct {
	uc     TrackUsecase
	helper helpers.Helper
	ch     base.BaseHandler
}

func NewTrackHandler(uc TrackUsecase, helper helpers.Helper, ch base.BaseHandler) *TrackHandler {
	return &TrackHandler{
		uc:     uc,
		helper: helper,
		ch:     ch,
	}
}

func (u *TrackHandler) GetTracks() {
	result, time, err := u.uc.GetTracks()
	u.ch.ErrorWrapper(gouse.DESC_GET_FAILED, err)
	u.ch.SuccessDataWrapper(gouse.DESC_GET_SUCCESS, result, time)
}

func (u *TrackHandler) GetTrack() {
	result, time, err := u.uc.GetTrack()
	u.ch.ErrorWrapper(gouse.DESC_GET_FAILED, err)
	u.ch.SuccessDataWrapper(gouse.DESC_GET_SUCCESS, result, time)
}

func (u *TrackHandler) CreateTrack() {
	result, time, err := u.uc.CreateTrack()
	u.ch.ErrorWrapper(gouse.DESC_CREATE_FAILED, err)
	u.ch.SuccessDataWrapper(gouse.DESC_CREATE_SUCCESS, result, time)
}

func (u *TrackHandler) DeleteTrack() {
	time, err := u.uc.DeleteTrack()
	u.ch.ErrorWrapper(gouse.DESC_DELETE_FAILED, err)
	u.ch.SuccessNoDataWrapper(gouse.DESC_DELETE_SUCCESS, time)
}

func (u *TrackHandler) UpdateTrack() {
	result, time, err := u.uc.UpdateTrack()
	u.ch.ErrorWrapper(gouse.DESC_UPDATE_FAILED, err)
	u.ch.SuccessDataWrapper(gouse.DESC_UPDATE_SUCCESS, result, time)
}
