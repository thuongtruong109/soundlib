package tracks

// import (
// 	"github.com/thuongtruong109/soundlib/pkg/helpers"
// 	"github.com/thuongtruong109/soundlib/pkg/constants"
// 	"github.com/thuongtruong109/soundlib/internal/common"
// )

// type TrackHandler struct {
// 	uc TrackUsecase
// 	helper helpers.Helper
// 	ch common.CommonHandler
// }

// func NewTrackHandler(uc TrackUsecase, helper helpers.Helper, ch common.CommonHandler) *TrackHandler {
// 	return &TrackHandler{
// 		uc: uc,
// 		helper: helper,
// 		ch: ch,
// 	}
// }

// func (u *TrackHandler) GetTracks() {
// 	result, err := u.uc.GetTracks()
// 	u.ch.ErrorWrapper(constants.GET_FAILED, err)
// 	u.ch.SuccessWrapper(constants.GET_SUCCESS, result)
// }

// func (u *TrackHandler) GetTrack() {
// 	result, err := u.uc.GetTrack()
// 	u.ch.ErrorWrapper(constants.GET_FAILED, err)
// 	u.ch.SuccessWrapper(constants.GET_SUCCESS, result)
// }

// func (u *TrackHandler) CreateTrack() {
// 	result, err := u.uc.CreateTrack()
// 	u.ch.ErrorWrapper(constants.CREATE_FAILED, err)
// 	u.ch.SuccessWrapper(constants.CREATE_SUCCESS, result)
// }

// func (u *TrackHandler) DeleteTrack() {
// 	err := u.uc.DeleteTrack()
// 	u.ch.ErrorWrapper(constants.DELETE_FAILED, err)
// 	u.helper.OutputSuccess(constants.DELETE_SUCCESS)
// }

// func (u *TrackHandler) UpdateTrack() {
// 	result, err := u.uc.UpdateTrack()
// 	u.ch.ErrorWrapper(constants.UPDATE_FAILED, err)
// 	u.ch.SuccessWrapper(constants.UPDATE_SUCCESS, result)
// }

// func (u *TrackHandler) GetAlbumsOfTrack() {
// 	u.helper.OutputSuccess("GetAlbumsOfTrack")
// }

// func (u *TrackHandler) GetTracksOfTrack() {
// 	u.helper.OutputSuccess("GetSongsOfTrack")
// }
