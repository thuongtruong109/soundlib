package genres

import (
	"github.com/thuongtruong109/soundlib/pkg/helpers"
	"github.com/thuongtruong109/soundlib/pkg/constants"
	"github.com/thuongtruong109/soundlib/internal/common"
)

type GenreHandler struct {
	uc GenreUsecase
	helper helpers.Helper
	ch common.CommonHandler
}

func NewGenreHandler(uc GenreUsecase, helper helpers.Helper, ch common.CommonHandler) *GenreHandler {
	return &GenreHandler{
		uc: uc,
		helper: helper,
		ch: ch,
	}
}

func (u *GenreHandler) GetGenres() {
	result, err := u.uc.GetGenres()
	u.ch.ErrorWrapper(constants.GET_FAILED, err)
	u.ch.SuccessWrapper(constants.GET_SUCCESS, result)
}

func (u *GenreHandler) GetGenre() {
	result, err := u.uc.GetGenre()
	u.ch.ErrorWrapper(constants.GET_FAILED, err)
	u.ch.SuccessWrapper(constants.GET_SUCCESS, result)
}

func (u *GenreHandler) CreateGenre() {
	result, err := u.uc.CreateGenre()
	u.ch.ErrorWrapper(constants.CREATE_FAILED, err)
	u.ch.SuccessWrapper(constants.CREATE_SUCCESS, result)
}

func (u *GenreHandler) DeleteGenre() {
	err := u.uc.DeleteGenre()
	u.ch.ErrorWrapper(constants.DELETE_FAILED, err)
	u.helper.OutputSuccess(constants.DELETE_SUCCESS)
}

func (u *GenreHandler) UpdateGenre() {
	result, err := u.uc.UpdateGenre()
	u.ch.ErrorWrapper(constants.UPDATE_FAILED, err)
	u.ch.SuccessWrapper(constants.UPDATE_SUCCESS, result)
}

func (u *GenreHandler) GetTracksOfGenre() {
	u.helper.OutputSuccess(constants.GET_SUCCESS)
}