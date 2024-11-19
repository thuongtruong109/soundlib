package artists

// import (
// 	"github.com/thuongtruong109/soundlib/internal/common"
// 	"github.com/thuongtruong109/soundlib/pkg/constants"
// 	"github.com/thuongtruong109/soundlib/pkg/helpers"
// )

// type ArtistHandler struct {
// 	uc     ArtistUsecase
// 	helper helpers.Helper
// 	ch     common.CommonHandler
// }

// func NewArtistHandler(uc ArtistUsecase, helper helpers.Helper, ch common.CommonHandler) *ArtistHandler {
// 	return &ArtistHandler{
// 		uc:     uc,
// 		helper: helper,
// 		ch:     ch,
// 	}
// }

// func (u *ArtistHandler) GetArtists() {
// 	result, err := u.uc.GetArtists()
// 	u.ch.ErrorWrapper(constants.GET_FAILED, err)
// 	u.ch.SuccessWrapper(constants.GET_SUCCESS, result)
// }

// func (u *ArtistHandler) GetArtist() {
// 	result, err := u.uc.GetArtist()
// 	u.ch.ErrorWrapper(constants.GET_FAILED, err)
// 	u.ch.SuccessWrapper(constants.GET_SUCCESS, result)
// }

// func (u *ArtistHandler) CreateArtist() {
// 	result, err := u.uc.CreateArtist()
// 	u.ch.ErrorWrapper(constants.CREATE_FAILED, err)
// 	u.ch.SuccessWrapper(constants.CREATE_SUCCESS, result)
// }

// func (u *ArtistHandler) DeleteArtist() {
// 	err := u.uc.DeleteArtist()
// 	u.ch.ErrorWrapper(constants.DELETE_FAILED, err)
// 	u.helper.OutputSuccess(constants.DELETE_SUCCESS)
// }

// func (u *ArtistHandler) UpdateArtist() {
// 	result, err := u.uc.UpdateArtist()
// 	u.ch.ErrorWrapper(constants.UPDATE_FAILED, err)
// 	u.ch.SuccessWrapper(constants.UPDATE_SUCCESS, result)
// }
