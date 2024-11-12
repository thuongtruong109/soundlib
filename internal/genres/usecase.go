package genres

import (
	"fmt"

	gu_helper "github.com/thuongtruong109/gouse/helper"
	"github.com/thuongtruong109/soundlib/pkg/constants"
	"github.com/thuongtruong109/soundlib/pkg/helpers"
)

type GenreUsecase struct {
	repo   GenreRepository
	helper helpers.Helper
}

func NewGenreUsecase(repo GenreRepository, helper helpers.Helper) *GenreUsecase {
	return &GenreUsecase{
		repo:   repo,
		helper: helper,
	}
}

func (g *GenreUsecase) GetGenres() ([]string, error) {
	result, err := helpers.QueryTimeTwoOutput[[]*Genre](g.repo.GetGenres)()

	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, fmt.Errorf(constants.NOT_FOUND_DATA)
	}

	var items []string
	for _, item := range result {
		items = append(items, fmt.Sprintf("Id: %s, Name: %s, Description: %s", item.ID, item.Name, item.Description))
	}

	return items, nil
}

func (g *GenreUsecase) GetGenre() ([]string, error) {
	var id string
	fmt.Print("» Enter id: ")
	fmt.Scanln(&id)

	result, err := helpers.QueryTimeTwoOutputWithParams[*Genre, string](g.repo.GetGenre)(id)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, fmt.Errorf(constants.NOT_FOUND_DATA)
	}

	newItem := []string{fmt.Sprintf("Id: %s, Name: %s, Description: %s", result.ID, result.Name, result.Description)}

	return newItem, nil
}

func (g *GenreUsecase) CreateGenre() ([]string, error) {
	var name string
	fmt.Print("» Enter name: ")
	fmt.Scanln(&name)

	var description string
	fmt.Print("» Enter description: ")
	fmt.Scanln(&description)

	newGenre := &Genre{
		ID:          gu_helper.RandomID(),
		Name:        name,
		Description: description,
	}

	result, err := helpers.QueryTimeTwoOutputWithParams[*Genre, *Genre](g.repo.CreateGenre)(newGenre)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, fmt.Errorf(constants.NOT_FOUND_DATA)
	}

	newItem := []string{fmt.Sprintf("Id: %s", result.ID)}

	return newItem, nil
}

func (g *GenreUsecase) DeleteGenre() error {
	var id string
	fmt.Print("» Enter id: ")
	fmt.Scanln(&id)

	err := helpers.QueryTimeOneOutputWithParams[error](g.repo.DeleteGenre)(id)
	if err != nil {
		return err
	}

	return nil
}

func (g *GenreUsecase) UpdateGenre() ([]string, error) {
	var id string
	fmt.Print("» Enter id: ")
	fmt.Scanln(&id)

	var name string
	fmt.Print("» Enter name: ")
	fmt.Scanln(&name)

	var description string
	fmt.Print("» Enter description: ")
	fmt.Scanln(&description)

	newGenre := &Genre{
		ID:          id,
		Name:        name,
		Description: description,
	}

	result, err := helpers.QueryTimeTwoOutputWithParams[*Genre, *Genre](g.repo.UpdateGenre)(newGenre)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, fmt.Errorf(constants.NOT_FOUND_DATA)
	}

	newItem := []string{fmt.Sprintf("Id: %s, Name: %s, Description: %s", result.ID, result.Name, result.Description)}

	return newItem, nil

}

func (g *GenreUsecase) GetTracksOfGenre() string {
	return "Tracks of Genre"
}
