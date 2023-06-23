package usecases

import (
	"fmt"
	"music-management/database"
	"music-management/pkg/constants"
	"music-management/pkg/helpers"
	"music-management/internal/models"
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

func (g *GenreUsecase) GetGenres() []interface{} {
	genresInterface, err := g.db.ReadFromDB(constants.GENRE_PATH)
	if err != nil {
		return nil
	}

	genres := make([]models.Genre, 0)
	for _, genreInterface := range genresInterface {
		if genreMap, ok := genreInterface.(map[string]interface{}); ok {
			genre := models.Genre{
				ID:          genreMap["id"].(string),
				Name:        genreMap["name"].(string),
				Description: genreMap["description"].(string),
			}
			genres = append(genres, genre)
		}
	}

	interfaceSlice := make([]interface{}, len(genres))
	for i, v := range genres {
		interfaceSlice[i] = fmt.Sprintf("ID: %s\nName: %s\nDescription: %s\n", v.ID, v.Name, v.Description)
	}

	return interfaceSlice
}

func (g *GenreUsecase) GetGenre() string {
	return "Genre by id"
}

func (g *GenreUsecase) CreateGenre() *models.Genre {
	genresInterface, err := g.db.ReadFromDB(constants.GENRE_PATH)
	if err != nil {
		return nil
	}

	var allGenres []interface{}

	if genresInterface == nil {
		allGenres = make([]interface{}, 0)
		
	} else {
		allGenres = make([]interface{}, len(genresInterface))
		copy(allGenres, genresInterface)
	}

	newGenre := &models.Genre{
		ID:          g.helper.GenerateID(),
		Name:        "Genre",
		Description: "Description",
	}

	allGenres = append(allGenres, newGenre)

	err2 := g.db.SaveToDB(constants.GENRE_PATH, allGenres)
	if err2 != nil {
		return nil
	}
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