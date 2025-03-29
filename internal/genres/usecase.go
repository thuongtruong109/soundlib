package genres

import (
	"fmt"

	"github.com/thuongtruong109/gouse"
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

func (g *GenreUsecase) GetGenres() ([]string, string, error) {
	result, time, err := helpers.QueryTimeTwoOutput(g.repo.GetGenres)()

	if err != nil {
		return nil, "", err
	}

	if result == nil {
		return nil, "", fmt.Errorf(gouse.DESC_NOT_FOUND_DATA)
	}

	var items []string
	for _, item := range result {
		items = append(items, fmt.Sprintf("Id: %s, Name: %s, Description: %s", item.ID, item.Name, item.Description))
	}

	return items, time, nil
}

func (g *GenreUsecase) GetGenre() ([]string, string, error) {
	var id string
	fmt.Print("» Enter id: ")
	fmt.Scanln(&id)

	result, time, err := helpers.QueryTimeTwoOutputWithParams(g.repo.GetGenre)(id)
	if err != nil {
		return nil, "", err
	}

	if result == nil {
		return nil, "", fmt.Errorf(gouse.DESC_NOT_FOUND_DATA)
	}

	newItem := []string{fmt.Sprintf("Id: %s, Name: %s, Description: %s", result.ID, result.Name, result.Description)}

	return newItem, time, nil
}

func (g *GenreUsecase) CreateGenre() ([]string, string, error) {
	var name string
	fmt.Print("» Enter name: ")
	fmt.Scanln(&name)

	var description string
	fmt.Print("» Enter description: ")
	fmt.Scanln(&description)

	newGenre := &Genre{
		ID:          gouse.RandID(),
		Name:        name,
		Description: description,
	}

	result, time, err := helpers.QueryTimeTwoOutputWithParams(g.repo.CreateGenre)(newGenre)
	if err != nil {
		return nil, time, err
	}

	if result == nil {
		return nil, time, fmt.Errorf(gouse.DESC_NOT_FOUND_DATA)
	}

	newItem := []string{fmt.Sprintf("Id: %s", result.ID)}

	return newItem, time, err
}

func (g *GenreUsecase) DeleteGenre() (string, error) {
	var id string
	fmt.Print("» Enter id: ")
	fmt.Scanln(&id)

	time, err := helpers.QueryTimeErrorWithOneParam[error](g.repo.DeleteGenre)(id)
	if err != nil {
		return "", fmt.Errorf(gouse.DESC_NOT_FOUND_DATA)
	}

	return time, err
}

func (g *GenreUsecase) UpdateGenre() ([]string, string, error) {
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

	result, time, err := helpers.QueryTimeTwoOutputWithParams(g.repo.UpdateGenre)(newGenre)
	if err != nil {
		return nil, time, err
	}

	if result == nil {
		return nil, time, fmt.Errorf(gouse.DESC_NOT_FOUND_DATA)
	}

	newItem := []string{fmt.Sprintf("Id: %s, Name: %s, Description: %s", result.ID, result.Name, result.Description)}

	return newItem, time, nil

}
