package artists

import (
	"github.com/thuongtruong109/gouse"
	"github.com/thuongtruong109/soundlib/pkg/common"
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
	u.ch.ErrorWrapper(gouse.DESC_GET_FAILED, err)
	u.ch.SuccessDataWrapper(gouse.DESC_GET_SUCCESS, result, time)
}

func (u *ArtistHandler) GetArtist() {
	result, time, err := u.uc.GetArtist()
	u.ch.ErrorWrapper(gouse.DESC_GET_FAILED, err)
	u.ch.SuccessDataWrapper(gouse.DESC_GET_SUCCESS, result, time)
}

func (u *ArtistHandler) CreateArtist() {
	result, time, err := u.uc.CreateArtist()
	u.ch.ErrorWrapper(gouse.DESC_CREATE_FAILED, err)
	u.ch.SuccessDataWrapper(gouse.DESC_CREATE_SUCCESS, result, time)
}

func (u *ArtistHandler) DeleteArtist() {
	time, err := u.uc.DeleteArtist()
	u.ch.ErrorWrapper(gouse.DESC_DELETE_FAILED, err)
	u.ch.SuccessNoDataWrapper(gouse.DESC_DELETE_SUCCESS, time)
}

func (u *ArtistHandler) UpdateArtist() {
	result, time, err := u.uc.UpdateArtist()
	u.ch.ErrorWrapper(gouse.DESC_UPDATE_FAILED, err)
	u.ch.SuccessDataWrapper(gouse.DESC_UPDATE_SUCCESS, result, time)
}
