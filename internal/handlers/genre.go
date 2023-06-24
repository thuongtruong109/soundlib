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
	result, err := u.uc.GetGenres()
	if err != nil {
		u.helper.OutputError(constants.GET_FAILED, err.Error())
		return
	}

	if result == nil {
		u.helper.OutputError(constants.GET_SUCCESS, constants.NOT_FOUND_DATA)
		return
	}

	u.helper.OutputSuccess(constants.GET_SUCCESS)

	var items []string
	for _, item := range result {
		items = append(items, fmt.Sprintf("Id: %s, Name: %s, Description: %s", item.ID, item.Name, item.Description))
	}

	helpers.TableOutput[string, string, interface{}]("Genre", items, nil)
}

func (u *GenreHandler) GetGenre() {
	result, err := u.uc.GetGenre()
	if err != nil {
		u.helper.OutputError(constants.GET_FAILED, err.Error())
		return
	}

	if result == nil {
		u.helper.OutputError(constants.GET_FAILED, constants.NOT_FOUND_DATA)
		return
	}

	u.helper.OutputSuccess(constants.GET_SUCCESS)

	newItem := []string{fmt.Sprintf("Id: %s, Name: %s, Description: %s", result.ID, result.Name, result.Description)}

	helpers.TableOutput[string, string, interface{}]("Genre", newItem, nil)
}

func (u *GenreHandler) CreateGenre() {
	result, err := u.uc.CreateGenre()
	if err != nil {
		u.helper.OutputError(constants.CREATE_FAILED, err.Error())
		return
	}

	if result == nil {
		u.helper.OutputError(constants.CREATE_FAILED, constants.NOT_FOUND_DATA)
		return
	}

	u.helper.OutputSuccess(constants.CREATE_SUCCESS)

	newItem := []string{fmt.Sprintf("Id: %s", result.ID)}

	helpers.TableOutput[string, string, interface{}]("Genre", newItem, nil)
}

func (u *GenreHandler) DeleteGenre() {
	err := u.uc.DeleteGenre()
	if err != nil {
		u.helper.OutputError(constants.DELETE_FAILED, err.Error())
		return
	}

	u.helper.OutputSuccess(constants.DELETE_SUCCESS)
}

func (u *GenreHandler) UpdateGenre() {
	result, err := u.uc.UpdateGenre()
	if err != nil {
		u.helper.OutputError(constants.UPDATE_FAILED, err.Error())
		return
	}

	if result == nil {
		u.helper.OutputError(constants.UPDATE_FAILED, constants.NOT_FOUND_DATA)
		return
	}

	u.helper.OutputSuccess(constants.UPDATE_SUCCESS)

	newItem := []string{fmt.Sprintf("Id: %s, Name: %s, Description: %s", result.ID, result.Name, result.Description)}

	helpers.TableOutput[string, string, interface{}]("Genre", newItem, nil)
}

func (u *GenreHandler) GetTracksOfGenre() {
	u.helper.OutputSuccess(constants.GET_SUCCESS)
}