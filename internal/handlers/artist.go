package handlers

import (
	"music-management/pkg/helpers"
	"music-management/pkg/constants"

	"music-management/internal/usecases"
)

type ArtistHandler struct {
	uc usecases.ArtistUsecase
	helper helpers.Helper
}

func NewArtistHandler(uc usecases.ArtistUsecase, helper helpers.Helper) *ArtistHandler {
	return &ArtistHandler{
		uc: uc,
		helper: helper,
	}
}

func (u *ArtistHandler) GetArtists() {
	u.helper.Output(constants.INFO, "GetArtists")
}

func (u *ArtistHandler) GetArtist() {
	u.helper.Output(constants.INFO, "GetArtist")
}

func (u *ArtistHandler) CreateArtist() {
	u.helper.Output(constants.INFO, "CreateArtist")
}

func (u *ArtistHandler) DeleteArtist() {
	u.helper.Output(constants.INFO, "DeleteArtist")
}

func (u *ArtistHandler) UpdateArtist() {
	u.helper.Output(constants.INFO, "UpdateArtist")
}

func (u *ArtistHandler) GetAlbumsOfArtist() {
	u.helper.Output(constants.INFO, "GetAlbumsOfArtist")
}

func (u *ArtistHandler) GetSongsOfArtist() {
	u.helper.Output(constants.INFO, "GetSongsOfArtist")
}
