package handlers

import (
	"fmt"
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
	result, err := u.uc.GetArtists()
	if err != nil {
		u.helper.OutputError(constants.GET_FAILED, err.Error())
		return
	}

	if result == nil {
		u.helper.OutputError(constants.GET_FAILED, constants.NOT_FOUND_DATA)
		return
	}

	u.helper.OutputSuccess(constants.GET_SUCCESS)

	var output []string
	for _, v := range result {
		output = append(output, fmt.Sprintf("ID: %s, Name: %s", v.ID, v.Name))
	}

	helpers.TableOutput[string, string, interface{}]("Artists", output, nil)
}

func (u *ArtistHandler) GetArtist() {
	result, err := u.uc.GetArtist()
	if err != nil {
		u.helper.OutputError(constants.GET_FAILED, err.Error())
		return
	}

	if result == nil {
		u.helper.OutputError(constants.GET_FAILED, constants.NOT_FOUND_DATA)
		return
	}

	u.helper.OutputSuccess(constants.GET_SUCCESS)

	output := []string{fmt.Sprintf("ID: %s, Name: %s", result.ID, result.Name)}
	helpers.TableOutput[string, string, interface{}]("Artists", output, nil)
}

func (u *ArtistHandler) CreateArtist() {
	result, err := u.uc.CreateArtist()
	if err != nil {
		u.helper.OutputError(constants.CREATE_FAILED, err.Error())
		return
	}

	u.helper.OutputSuccess(constants.CREATE_SUCCESS)

	output := []string{fmt.Sprintf("ID: %s", result.ID)}
	helpers.TableOutput[string, string, interface{}]("Artists", output, nil)
}

func (u *ArtistHandler) DeleteArtist() {
	err := u.uc.DeleteArtist()
	if err != nil {
		u.helper.OutputError(constants.DELETE_FAILED, err.Error())
		return
	}

	u.helper.OutputSuccess(constants.DELETE_SUCCESS)
}

func (u *ArtistHandler) UpdateArtist() {
	result, err := u.uc.UpdateArtist()
	if err != nil {
		u.helper.OutputError(constants.UPDATE_FAILED, err.Error())
		return
	}

	u.helper.OutputSuccess(constants.UPDATE_SUCCESS)

	output := []string{fmt.Sprintf("ID: %s, Name: %s", result.ID, result.Name)}
	helpers.TableOutput[string, string, interface{}]("Artists", output, nil)
}

func (u *ArtistHandler) GetAlbumsOfArtist() {
	u.helper.OutputSuccess("GetAlbumsOfArtist")
}

func (u *ArtistHandler) GetTracksOfArtist() {
	u.helper.OutputSuccess("GetSongsOfArtist")
}
