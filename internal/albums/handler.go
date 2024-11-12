package albums

import (
	"github.com/thuongtruong109/soundlib/pkg/helpers"
)

type AlbumHandler struct {
	uc     AlbumUsecase
	helper helpers.Helper
}

func NewAlbumHandler(uc AlbumUsecase, helper helpers.Helper) *AlbumHandler {
	return &AlbumHandler{
		uc:     uc,
		helper: helper,
	}
}

func (u *AlbumHandler) GetAlbums() {
	u.helper.OutputSuccess("GetAlbums")
}

func (u *AlbumHandler) GetAlbum() {
	u.helper.OutputSuccess("GetAlbum")
}

func (u *AlbumHandler) CreateAlbum() {
	u.helper.OutputSuccess("CreateAlbum")
}

func (u *AlbumHandler) DeleteAlbum() {
	u.helper.OutputSuccess("DeleteAlbum")
}

func (u *AlbumHandler) UpdateAlbum() {
	u.helper.OutputSuccess("UpdateAlbum")
}

func (u *AlbumHandler) GetTracksOfAlbum() {
	u.helper.OutputSuccess("GetTracksOfAlbum")
}
