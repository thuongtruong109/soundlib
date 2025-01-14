package playlists

import (
	"github.com/thuongtruong109/gouse"
	"github.com/thuongtruong109/soundlib/pkg/base"
	"github.com/thuongtruong109/soundlib/pkg/helpers"
)

type PlaylistHandler struct {
	uc     PlaylistUsecase
	helper helpers.Helper
	ch     base.BaseHandler
}

func NewPlaylistHandler(uc PlaylistUsecase, helper helpers.Helper, ch base.BaseHandler) *PlaylistHandler {
	return &PlaylistHandler{
		uc:     uc,
		helper: helper,
		ch:     ch,
	}
}

func (u *PlaylistHandler) GetPlaylists() {
	result, time, err := u.uc.GetPlaylists()
	u.ch.ErrorWrapper(gouse.DESC_GET_FAILED, err)
	u.ch.SuccessDataWrapper(gouse.DESC_GET_SUCCESS, result, time)
}

func (u *PlaylistHandler) GetPlaylist() {
	result, time, err := u.uc.GetPlaylist()
	u.ch.ErrorWrapper(gouse.DESC_GET_FAILED, err)
	u.ch.SuccessDataWrapper(gouse.DESC_GET_SUCCESS, result, time)
}

func (u *PlaylistHandler) CreatePlaylist() {
	result, time, err := u.uc.CreatePlaylist()
	u.ch.ErrorWrapper(gouse.DESC_CREATE_FAILED, err)
	u.ch.SuccessDataWrapper(gouse.DESC_CREATE_SUCCESS, result, time)
}

func (u *PlaylistHandler) DeletePlaylist() {
	time, err := u.uc.DeletePlaylist()
	u.ch.ErrorWrapper(gouse.DESC_DELETE_FAILED, err)
	u.ch.SuccessNoDataWrapper(gouse.DESC_DELETE_SUCCESS, time)
}

func (u *PlaylistHandler) UpdatePlaylist() {
	result, time, err := u.uc.UpdatePlaylist()
	u.ch.ErrorWrapper(gouse.DESC_UPDATE_FAILED, err)
	u.ch.SuccessDataWrapper(gouse.DESC_UPDATE_SUCCESS, result, time)
}
