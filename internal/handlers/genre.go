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
	result := u.uc.GetGenres()
	if result == nil {
		u.helper.Output(constants.ERROR, constants.GET_FAILED)
		return
	}

	u.helper.Output(constants.INFO, constants.GET_SUCCESS)

	var items []string
	for _, item := range result {
		items = append(items, fmt.Sprintf("Id: %s, Name: %s, Description: %s", item.ID, item.Name, item.Description))
	}

	helpers.TableOutput[string, string, interface{}]("Genre", items, nil)
}

func (u *GenreHandler) GetGenre() {
	result := u.uc.GetGenre()
	if result == nil {
		u.helper.Output(constants.ERROR, constants.GET_FAILED)
		return
	}

	u.helper.Output(constants.INFO, constants.GET_SUCCESS)

	newItem := []string{fmt.Sprintf("Id: %s, Name: %s, Description: %s", result.ID, result.Name, result.Description)}

	helpers.TableOutput[string, string, interface{}]("Genre", newItem, nil)
}

func (u *GenreHandler) CreateGenre() {
	result := u.uc.CreateGenre()
	if result == nil {
		u.helper.Output(constants.ERROR, constants.CREATE_FAILED)
		return
	}

	u.helper.Output(constants.INFO, constants.CREATE_SUCCESS)

	newItem := []string{fmt.Sprintf("Id: %s", result.ID)}

	helpers.TableOutput[string, string, interface{}]("Genre", newItem, nil)
}

func (u *GenreHandler) DeleteGenre() {
	result := u.uc.DeleteGenre()
	if result != nil {
		u.helper.Output(constants.ERROR, constants.DELETE_FAILED)
		return
	}

	u.helper.Output(constants.INFO, constants.DELETE_SUCCESS)
}

func (u *GenreHandler) UpdateGenre() {
	result := u.uc.UpdateGenre()
	if result == nil {
		u.helper.Output(constants.ERROR, constants.UPDATE_FAILED)
		return
	}

	u.helper.Output(constants.INFO, constants.UPDATE_SUCCESS)

	newItem := []string{fmt.Sprintf("Id: %s, Name: %s, Description: %s", result.ID, result.Name, result.Description)}

	helpers.TableOutput[string, string, interface{}]("Genre", newItem, nil)
}

func (u *GenreHandler) GetAlbumsOfGenre() {
	u.helper.Output(constants.INFO, "GetAlbumsOfGenre")
}

func (u *GenreHandler) GetSongsOfGenre() {
	u.helper.Output(constants.INFO, "GetSongsOfGenre")
}