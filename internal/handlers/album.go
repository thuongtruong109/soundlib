package handlers

import (
	"music-management/pkg/helpers"
	"music-management/pkg/constants"
	
	"music-management/internal/usecases"
)

type AlbumHandler struct {
	uc usecases.AlbumUsecase
	helper helpers.Helper
}

func NewAlbumHandler(uc usecases.AlbumUsecase, helper helpers.Helper) *AlbumHandler {
	return &AlbumHandler{
		uc: uc,
		helper: helper,
	}
}

func (u *AlbumHandler) GetAlbums() {
	u.helper.Output(constants.INFO, "GetAlbums")
}

func (u *AlbumHandler) GetAlbum() {
	u.helper.Output(constants.INFO, "GetAlbum")
}

func (u *AlbumHandler) CreateAlbum() {
	u.helper.Output(constants.INFO, "CreateAlbum")
}

func (u *AlbumHandler) DeleteAlbum() {
	u.helper.Output(constants.INFO, "DeleteAlbum")
}

func (u *AlbumHandler) GetAlbumsOfArtist() {
	u.helper.Output(constants.INFO, "GetAlbumsOfArtist")
}

func (u *AlbumHandler) GetAlbumsOfGenre() {
	u.helper.Output(constants.INFO, "GetAlbumsOfGenre")
}
