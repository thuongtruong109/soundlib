package artists

import (
	"github.com/thuongtruong109/soundlib/pkg/common"
	"github.com/thuongtruong109/soundlib/pkg/constants"
	"github.com/thuongtruong109/soundlib/pkg/helpers"
)

type ArtistHandler struct {
	uc     ArtistUsecase
	helper helpers.Helper
	ch     common.CommonHandler
}

func NewArtistHandler(uc ArtistUsecase, helper helpers.Helper, ch common.CommonHandler) *ArtistHandler {
	return &ArtistHandler{
		uc:     uc,
		helper: helper,
		ch:     ch,
	}
}

func (u *ArtistHandler) GetArtists() {
	result, time, err := u.uc.GetArtists()
	u.ch.ErrorWrapper(constants.GET_FAILED, err)
	u.ch.SuccessDataWrapper(constants.GET_SUCCESS, result, time)
}

func (u *ArtistHandler) GetArtist() {
	result, time, err := u.uc.GetArtist()
	u.ch.ErrorWrapper(constants.GET_FAILED, err)
	u.ch.SuccessDataWrapper(constants.GET_SUCCESS, result, time)
}

func (u *ArtistHandler) CreateArtist() {
	result, time, err := u.uc.CreateArtist()
	u.ch.ErrorWrapper(constants.CREATE_FAILED, err)
	u.ch.SuccessDataWrapper(constants.CREATE_SUCCESS, result, time)
}

func (u *ArtistHandler) DeleteArtist() {
	time, err := u.uc.DeleteArtist()
	u.ch.ErrorWrapper(constants.DELETE_FAILED, err)
	u.ch.SuccessNoDataWrapper(constants.DELETE_SUCCESS, time)
}

func (u *ArtistHandler) UpdateArtist() {
	result, time, err := u.uc.UpdateArtist()
	u.ch.ErrorWrapper(constants.UPDATE_FAILED, err)
	u.ch.SuccessDataWrapper(constants.UPDATE_SUCCESS, result, time)
}
