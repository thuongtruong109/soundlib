package genres

import (
	"github.com/thuongtruong109/soundlib/pkg/common"
	"github.com/thuongtruong109/soundlib/pkg/constants"
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
	u.ch.ErrorWrapper(constants.GET_FAILED, err)
	u.ch.SuccessDataWrapper(constants.GET_SUCCESS, result, time)
}

func (u *GenreHandler) GetGenre() {
	result, time, err := u.uc.GetGenre()
	u.ch.ErrorWrapper(constants.GET_FAILED, err)
	u.ch.SuccessDataWrapper(constants.GET_SUCCESS, result, time)
}

func (u *GenreHandler) CreateGenre() {
	result, time, err := u.uc.CreateGenre()
	u.ch.ErrorWrapper(constants.CREATE_FAILED, err)
	u.ch.SuccessDataWrapper(constants.CREATE_SUCCESS, result, time)
}

func (u *GenreHandler) DeleteGenre() {
	time, err := u.uc.DeleteGenre()
	u.ch.ErrorWrapper(constants.DELETE_FAILED, err)
	u.ch.SuccessNoDataWrapper(constants.DELETE_SUCCESS, time)
}

func (u *GenreHandler) UpdateGenre() {
	result, time, err := u.uc.UpdateGenre()
	u.ch.ErrorWrapper(constants.UPDATE_FAILED, err)
	u.ch.SuccessDataWrapper(constants.UPDATE_SUCCESS, result, time)
}
