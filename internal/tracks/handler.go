package tracks

import (
	"github.com/thuongtruong109/soundlib/pkg/common"
	"github.com/thuongtruong109/soundlib/pkg/constants"
	"github.com/thuongtruong109/soundlib/pkg/helpers"
)

type TrackHandler struct {
	uc TrackUsecase
	helper helpers.Helper
	ch common.CommonHandler
}

func NewTrackHandler(uc TrackUsecase, helper helpers.Helper, ch common.CommonHandler) *TrackHandler {
	return &TrackHandler{
		uc: uc,
		helper: helper,
		ch: ch,
	}
}

func (u *TrackHandler) GetTracks() {
	result, time, err := u.uc.GetTracks()
	u.ch.ErrorWrapper(constants.GET_FAILED, err)
	u.ch.SuccessDataWrapper(constants.GET_SUCCESS, result, time)
}

func (u *TrackHandler) GetTrack() {
	result, time, err := u.uc.GetTrack()
	u.ch.ErrorWrapper(constants.GET_FAILED, err)
	u.ch.SuccessDataWrapper(constants.GET_SUCCESS, result, time)
}

func (u *TrackHandler) CreateTrack() {
	result, time, err := u.uc.CreateTrack()
	u.ch.ErrorWrapper(constants.CREATE_FAILED, err)
	u.ch.SuccessDataWrapper(constants.CREATE_SUCCESS, result, time)
}

func (u *TrackHandler) DeleteTrack() {
	time, err := u.uc.DeleteTrack()
	u.ch.ErrorWrapper(constants.DELETE_FAILED, err)
	u.ch.SuccessNoDataWrapper(constants.DELETE_SUCCESS, time)
}

func (u *TrackHandler) UpdateTrack() {
	result, time, err := u.uc.UpdateTrack()
	u.ch.ErrorWrapper(constants.UPDATE_FAILED, err)
	u.ch.SuccessDataWrapper(constants.UPDATE_SUCCESS, result, time)
}

// func (u *TrackHandler) GetAlbumsOfTrack() {
// 	u.helper.OutputSuccess("GetAlbumsOfTrack")
// }

// func (u *TrackHandler) GetTracksOfTrack() {
// 	u.helper.OutputSuccess("GetSongsOfTrack")
// }
