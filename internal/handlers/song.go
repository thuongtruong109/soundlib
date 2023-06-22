package handlers

import (
	"music-management/pkg/helpers"
	"music-management/pkg/constants"
	
	"music-management/internal/usecases"
)

type SongHandler struct {
	uc usecases.SongUsecase
	helper helpers.Helper
}

func NewSongHandler(uc usecases.SongUsecase, helper helpers.Helper) *SongHandler {
	return &SongHandler{
		uc: uc,
		helper: helper,
	}
}

func (u *SongHandler) GetSong() {
	u.helper.Output(constants.INFO, "GetSong")
}

func (u *SongHandler) GetSongs() {
	u.helper.Output(constants.INFO, "GetSongs")
}

func (u *SongHandler) CreateSong() {
	u.helper.Output(constants.INFO, "CreateSong")
}

func (u *SongHandler) DeleteSong() {
	u.helper.Output(constants.INFO, "DeleteSong")
}

func (u *SongHandler) UpdateSong() {
	u.helper.Output(constants.INFO, "UpdateSong")
}

func (u *SongHandler) GetSongByAlbum() {
	u.helper.Output(constants.INFO, "GetSongByAlbum")
}

func (u *SongHandler) GetSongByArtist() {
	u.helper.Output(constants.INFO, "GetSongByArtist")
}