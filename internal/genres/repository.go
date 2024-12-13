package genres

import (
	"fmt"

	"github.com/thuongtruong109/gouse"
	"github.com/thuongtruong109/soundlib/pkg/constants"
)

type GenreRepository struct{}

func NewGenreRepository() *GenreRepository {
	return &GenreRepository{}
}

func (gr *GenreRepository) GetGenres() ([]*Genre, error) {
	genres, err := gouse.ReadFileObj[*Genre](constants.GENRE_PATH)

	if err != nil {
		return nil, err
	}

	if genres == nil {
		return nil, nil
	}

	return genres, nil
}

func (gr *GenreRepository) GetGenre(genreID string) (*Genre, error) {
	allGenres, err := gr.GetGenres()
	if err != nil {
		return nil, err
	}

	if allGenres == nil {
		return nil, nil
	}

	for _, genre := range allGenres {
		if genre.ID == genreID {
			return genre, nil
		}
	}

	return nil, nil
}

func (gr *GenreRepository) CreateGenre(newGenre *Genre) (*Genre, error) {
	allGenres, _ := gr.GetGenres()

	var genresInit []*Genre

	if allGenres == nil {
		genresInit = make([]*Genre, 0)
	} else {
		genresInit = make([]*Genre, len(allGenres))
		copy(genresInit, allGenres)
	}

	genresInit = append(genresInit, newGenre)

	err2 := gouse.WriteFileObj(constants.GENRE_PATH, genresInit)
	if err2 != nil {
		return nil, fmt.Errorf(constants.CREATE_FAILED)
	}
	return newGenre, nil
}

func (gr *GenreRepository) UpdateGenre(updateGenre *Genre) (*Genre, error) {
	allGenres, err := gr.GetGenres()
	if err != nil {
		return nil, err
	}

	if allGenres == nil {
		return nil, nil
	}

	var genresInit []*Genre

	for i, genre := range allGenres {
		if genre.ID == updateGenre.ID {
			genresInit = append(allGenres[:i], updateGenre)
			break
		}
	}

	genresInit = append(genresInit, allGenres[len(genresInit):]...)

	err2 := gouse.WriteFileObj[[]*Genre](constants.GENRE_PATH, genresInit)
	if err2 != nil {
		return nil, fmt.Errorf(constants.UPDATE_FAILED)
	}
	return updateGenre, nil
}

func (gr *GenreRepository) DeleteGenre(genreID string) error {
	allGenres, _ := gr.GetGenres()

	if allGenres == nil {
		return nil
	}

	var genresInit []*Genre

	for i, genre := range allGenres {
		if genre.ID == genreID {
			genresInit = append(allGenres[:i], allGenres[i+1:]...)
			break
		}
	}

	err2 := gouse.WriteFileObj[[]*Genre](constants.GENRE_PATH, genresInit)
	if err2 != nil {
		return fmt.Errorf(constants.DELETE_FAILED)
	}

	return nil
}
