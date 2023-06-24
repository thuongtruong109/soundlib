package usecases

import (
	"fmt"
	"music-management/database"
	"music-management/pkg/constants"
	"music-management/pkg/helpers"
	"music-management/internal/models"
)

type GenreUsecase struct {
	helper helpers.Helper
}

func NewGenreUsecase(helper helpers.Helper) *GenreUsecase {
	return &GenreUsecase{
		helper: helper,
	}
}

func (g *GenreUsecase) GetGenres() ([]*models.Genre, error) {
	genresInterface, err := database.ReadDB[*models.Genre](constants.GENRE_PATH)

	if err != nil {
		return nil, err
	}

	if genresInterface == nil {
		return nil, nil
	}

	genres := make([]*models.Genre, len(genresInterface))
	copy(genres, genresInterface)

	return genres, nil
}

func (g *GenreUsecase) GetGenre() (*models.Genre, error) {
	allGenres, err := g.GetGenres()
	if err != nil {
		return nil, err
	}

	if allGenres == nil {
		return nil, nil
	}

	var id string
	fmt.Print("» Enter id: ")
	fmt.Scanln(&id)

	for _, genre := range allGenres {
		if genre.ID == id {
			return genre, nil
		}
	}

	return nil, nil
}

func (g *GenreUsecase) CreateGenre() (*models.Genre, error) {
	allGenres, _ := g.GetGenres()

	var genresInit []*models.Genre

	if allGenres == nil {
		genresInit = make([]*models.Genre, 0)
	} else {
		genresInit = make([]*models.Genre, len(allGenres))
		copy(genresInit, allGenres)
	}

	var name string
	fmt.Print("» Enter name: ")
	fmt.Scanln(&name)

	var description string
	fmt.Print("» Enter description: ")
	fmt.Scanln(&description)

	newGenre := &models.Genre{
		ID:          g.helper.GenerateID(),
		Name:        name,
		Description: description,
	}

	genresInit = append(genresInit, newGenre)

	err2 := database.SaveDB[[]*models.Genre](constants.GENRE_PATH, genresInit)
	if err2 != nil {
		return nil, fmt.Errorf(constants.CREATE_FAILED)
	}
	return newGenre, nil
}

func (g *GenreUsecase) DeleteGenre() error {
	allGenres, err := g.GetGenres()
	if err != nil {
		return err
	}

	if allGenres == nil {
		return nil
	}

	var id string
	fmt.Print("» Enter id: ")
	fmt.Scanln(&id)

	var genresInit []*models.Genre

	for i, genre := range allGenres {
		if genre.ID == id {
			genresInit = append(allGenres[:i], allGenres[i+1:]...)
			break
		}
	}

	err2 := database.SaveDB[[]*models.Genre](constants.GENRE_PATH, genresInit)
	if err2 != nil {
		return fmt.Errorf(constants.DELETE_FAILED)
	}

	return nil
}

func (g *GenreUsecase) UpdateGenre() (*models.Genre, error) {
	allGenres, err := g.GetGenres()
	if err != nil {
		return nil, err
	}

	if allGenres == nil {
		return nil, nil
	}

	var id string
	fmt.Print("» Enter id: ")
	fmt.Scanln(&id)

	var name string
	fmt.Print("» Enter name: ")
	fmt.Scanln(&name)

	var description string
	fmt.Print("» Enter description: ")
	fmt.Scanln(&description)

	newGenre := &models.Genre{
		ID:          id,
		Name:        name,
		Description: description,
	}

	var genresInit []*models.Genre

	for i, genre := range allGenres {
		if genre.ID == id {
			genresInit = append(allGenres[:i], newGenre)
			break
		}
	}

	genresInit = append(genresInit, allGenres[len(genresInit):]...)

	err2 := database.SaveDB[[]*models.Genre](constants.GENRE_PATH, genresInit)
	if err2 != nil {
		return nil, fmt.Errorf(constants.UPDATE_FAILED)
	}
	return newGenre, nil
}

func (g *GenreUsecase) GetTracksOfGenre() string {
	return "Tracks of Genre"
}