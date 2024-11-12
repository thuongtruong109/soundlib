package genres

import (
	"fmt"

	"github.com/thuongtruong109/gouse/io"
	"github.com/thuongtruong109/soundlib/pkg/constants"
	"github.com/thuongtruong109/soundlib/pkg/helpers"
)

type GenreRepository struct {
	helper helpers.Helper
}

func NewGenreRepository(helper helpers.Helper) *GenreRepository {
	return &GenreRepository{
		helper: helper,
	}
}

func (g *GenreRepository) GetGenres() ([]*Genre, error) {
	genres, err := io.ReadFileObj[*Genre](constants.GENRE_PATH)

	if err != nil {
		return nil, err
	}

	if genres == nil {
		return nil, nil
	}

	return genres, nil
}

func (g *GenreRepository) GetGenre(genreID string) (*Genre, error) {
	allGenres, err := g.GetGenres()
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

func (g *GenreRepository) CreateGenre(newGenre *Genre) (*Genre, error) {
	allGenres, _ := g.GetGenres()

	var genresInit []*Genre

	if allGenres == nil {
		genresInit = make([]*Genre, 0)
	} else {
		genresInit = make([]*Genre, len(allGenres))
		copy(genresInit, allGenres)
	}

	genresInit = append(genresInit, newGenre)

	err2 := io.WriteFileObj[[]*Genre](constants.GENRE_PATH, genresInit)
	if err2 != nil {
		return nil, fmt.Errorf(constants.CREATE_FAILED)
	}
	return newGenre, nil
}

func (g *GenreRepository) UpdateGenre(updateGenre *Genre) (*Genre, error) {
	allGenres, err := g.GetGenres()
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

	err2 := io.WriteFileObj[[]*Genre](constants.GENRE_PATH, genresInit)
	if err2 != nil {
		return nil, fmt.Errorf(constants.UPDATE_FAILED)
	}
	return updateGenre, nil
}

func (g *GenreRepository) DeleteGenre(genreID string) error {
	allGenres, err := g.GetGenres()
	if err != nil {
		return err
	}

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

	err2 := io.WriteFileObj[[]*Genre](constants.GENRE_PATH, genresInit)
	if err2 != nil {
		return fmt.Errorf(constants.DELETE_FAILED)
	}

	return nil
}
