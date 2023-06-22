package handlers

import (
	"music-management/pkg/helpers"
	"music-management/pkg/constants"

	"music-management/internal/usecases"
)

type PlaylistHandler struct {
	uc usecases.PlaylistUsecase
	helper helpers.Helper
}

func NewPlaylistHandler(uc usecases.PlaylistUsecase, helper helpers.Helper) *PlaylistHandler {
	return &PlaylistHandler{
		uc: uc,
		helper: helper,
	}
}

func (u *PlaylistHandler) GetPlaylists() {
	u.helper.Output(constants.INFO, "GetPlaylists")
}

func (u *PlaylistHandler) GetPlaylist() {
	u.helper.Output(constants.INFO, "GetPlaylist")
}

func (u *PlaylistHandler) CreatePlaylist() {
	u.helper.Output(constants.INFO, "CreatePlaylist")
}

func (u *PlaylistHandler) DeletePlaylist() {
	u.helper.Output(constants.INFO, "DeletePlaylist")
}

func (u *PlaylistHandler) UpdatePlaylist() {
	u.helper.Output(constants.INFO, "UpdatePlaylist")
}

func (u *PlaylistHandler) GetPlaylistByAlbum() {
	u.helper.Output(constants.INFO, "GetPlaylistByAlbum")
}

func (u *PlaylistHandler) GetPlaylistByArtist() {
	u.helper.Output(constants.INFO, "GetPlaylistByArtist")
}
