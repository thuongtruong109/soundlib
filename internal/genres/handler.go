package genres

import (
	"github.com/thuongtruong109/gouse"
	"github.com/thuongtruong109/soundlib/pkg/common"
	"github.com/thuongtruong109/soundlib/pkg/helpers"
)

type GenreHandler struct {
	uc     GenreUsecase
	helper helpers.Helper
	ch     common.CommonHandler
}

func NewGenreHandler(uc GenreUsecase, helper helpers.Helper, ch common.CommonHandler) *GenreHandler {
	return &GenreHandler{
		uc:     uc,
		helper: helper,
		ch:     ch,
	}
}

func (u *GenreHandler) GetGenres() {
	result, time, err := u.uc.GetGenres()
	u.ch.ErrorWrapper(gouse.DESC_GET_FAILED, err)
	u.ch.SuccessDataWrapper(gouse.DESC_GET_SUCCESS, result, time)
}

func (u *GenreHandler) GetGenre() {
	result, time, err := u.uc.GetGenre()
	u.ch.ErrorWrapper(gouse.DESC_GET_FAILED, err)
	u.ch.SuccessDataWrapper(gouse.DESC_GET_SUCCESS, result, time)
}

func (u *GenreHandler) CreateGenre() {
	result, time, err := u.uc.CreateGenre()
	u.ch.ErrorWrapper(gouse.DESC_CREATE_FAILED, err)
	u.ch.SuccessDataWrapper(gouse.DESC_CREATE_SUCCESS, result, time)
}

func (u *GenreHandler) DeleteGenre() {
	time, err := u.uc.DeleteGenre()
	u.ch.ErrorWrapper(gouse.DESC_DELETE_FAILED, err)
	u.ch.SuccessNoDataWrapper(gouse.DESC_DELETE_SUCCESS, time)
}

func (u *GenreHandler) UpdateGenre() {
	result, time, err := u.uc.UpdateGenre()
	u.ch.ErrorWrapper(gouse.DESC_UPDATE_FAILED, err)
	u.ch.SuccessDataWrapper(gouse.DESC_UPDATE_SUCCESS, result, time)
}
