package usecases

import (
	"music-management/repositories"
	"music-management/pkg/helpers"
	"music-management/pkg/constants"
)

type AlbumUsecases struct {
	repo repositories.AlbumRepository
}

func NewAlbumUsecases(repo repositories.AlbumRepository) *AlbumUsecases {
	return &AlbumUsecases{
		repo: repo,
	}
}

func (u *AlbumUsecases) GetAlbums() {
	helpers.Output(constants.INFO, "GetAlbums")
}

func (u *AlbumUsecases) GetAlbum() {
	helpers.Output(constants.INFO, "GetAlbum")
}

func (u *AlbumUsecases) CreateAlbum() {
	helpers.Output(constants.INFO, "CreateAlbum")
}

func (u *AlbumUsecases) DeleteAlbum() {
	helpers.Output(constants.INFO, "DeleteAlbum")
}