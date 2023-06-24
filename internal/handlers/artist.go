package handlers

import (
	"music-management/pkg/helpers"

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
	u.helper.OutputSuccess("GetArtists")
}

func (u *ArtistHandler) GetArtist() {
	u.helper.OutputSuccess("GetArtist")
}

func (u *ArtistHandler) CreateArtist() {
	u.helper.OutputSuccess("CreateArtist")
}

func (u *ArtistHandler) DeleteArtist() {
	u.helper.OutputSuccess("DeleteArtist")
}

func (u *ArtistHandler) UpdateArtist() {
	u.helper.OutputSuccess("UpdateArtist")
}

func (u *ArtistHandler) GetAlbumsOfArtist() {
	u.helper.OutputSuccess("GetAlbumsOfArtist")
}

func (u *ArtistHandler) GetTracksOfArtist() {
	u.helper.OutputSuccess("GetSongsOfArtist")
}
