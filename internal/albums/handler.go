package albums

import (
	"github.com/thuongtruong109/gouse"
	"github.com/thuongtruong109/soundlib/pkg/base"
	"github.com/thuongtruong109/soundlib/pkg/helpers"
)

type AlbumHandler struct {
	uc     AlbumUsecase
	helper helpers.Helper
	ch     base.BaseHandler
}

func NewAlbumHandler(uc AlbumUsecase, helper helpers.Helper, ch base.BaseHandler) *AlbumHandler {
	return &AlbumHandler{
		uc:     uc,
		helper: helper,
		ch:     ch,
	}
}

func (u *AlbumHandler) GetAlbums() {
	result, time, err := u.uc.GetAlbums()
	u.ch.ErrorWrapper(gouse.DESC_GET_FAILED, err)
	u.ch.SuccessDataWrapper(gouse.DESC_GET_SUCCESS, result, time)
}

func (u *AlbumHandler) GetAlbum() {
	result, time, err := u.uc.GetAlbum()
	u.ch.ErrorWrapper(gouse.DESC_GET_FAILED, err)
	u.ch.SuccessDataWrapper(gouse.DESC_GET_SUCCESS, result, time)
}

func (u *AlbumHandler) CreateAlbum() {
	result, time, err := u.uc.CreateAlbum()
	u.ch.ErrorWrapper(gouse.DESC_CREATE_FAILED, err)
	u.ch.SuccessDataWrapper(gouse.DESC_CREATE_SUCCESS, result, time)
}

func (u *AlbumHandler) DeleteAlbum() {
	time, err := u.uc.DeleteAlbum()
	u.ch.ErrorWrapper(gouse.DESC_DELETE_FAILED, err)
	u.ch.SuccessNoDataWrapper(gouse.DESC_DELETE_SUCCESS, time)
}

func (u *AlbumHandler) UpdateAlbum() {
	result, time, err := u.uc.UpdateAlbum()
	u.ch.ErrorWrapper(gouse.DESC_UPDATE_FAILED, err)
	u.ch.SuccessDataWrapper(gouse.DESC_UPDATE_SUCCESS, result, time)
}
