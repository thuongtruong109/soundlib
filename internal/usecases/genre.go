package usecases

import (
	"music-management/database"
	"music-management/internal/models"
	"music-management/pkg/helpers"
)

type GenreUsecase struct {
	db database.Database
	helper helpers.Helper
}

func NewGenreUsecase(db database.Database, helper helpers.Helper) *GenreUsecase {
	return &GenreUsecase{
		db: db,
		helper: helper,
	}
}

func (g *GenreUsecase) GetGenres() string {
	return "Genres"
}

func (g *GenreUsecase) GetGenre() string {
	return "Genre by id"
}

func (g *GenreUsecase) CreateGenre() *models.Genre {
	newGenre := &models.Genre{
		ID: g.helper.GenerateID(),
		Name: "Genre",
		Description: "Description",
	}

	g.db.SaveToDB("genres.json", *newGenre)
	return newGenre
}

func (g *GenreUsecase) DeleteGenre() string {
	return "Delete Genre"
}

func (g *GenreUsecase) UpdateGenre() string {
	return "Update Genre"
}

func (g *GenreUsecase) GetAlbumsOfGenre() string {
	return "Albums of genre"
}

func (g *GenreUsecase) GetSongsOfGenre() string {
	return "Songs of genre"
}