package handlers

import (
	"fmt"
	"music-management/pkg/helpers"
	"music-management/pkg/constants"

	"music-management/internal/usecases"
)

type GenreHandler struct {
	uc usecases.GenreUsecase
	helper helpers.Helper
}

func NewGenreHandler(uc usecases.GenreUsecase, helper helpers.Helper) *GenreHandler {
	return &GenreHandler{
		uc: uc,
		helper: helper,
	}
}

func (u *GenreHandler) GetGenres() {
	u.helper.Output(constants.INFO, "GetGenres")
}

func (u *GenreHandler) GetGenre() {
	u.helper.Output(constants.INFO, "GetGenre")
}

func (u *GenreHandler) CreateGenre() {
	result := u.uc.CreateGenre()

	display := fmt.Sprintf("ID: %s\n", result.ID)

	u.helper.Output(constants.INFO, "Create successfully")

	u.helper.TableOutput("Genre", []interface{}{display}, nil)

}

func (u *GenreHandler) DeleteGenre() {
	u.helper.Output(constants.INFO, "DeleteGenre")
}

func (u *GenreHandler) UpdateGenre() {
	u.helper.Output(constants.INFO, "UpdateGenre")
}

func (u *GenreHandler) GetAlbumsOfGenre() {
	u.helper.Output(constants.INFO, "GetAlbumsOfGenre")
}

func (u *GenreHandler) GetSongsOfGenre() {
	u.helper.Output(constants.INFO, "GetSongsOfGenre")
}