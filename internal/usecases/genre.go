package usecases

type GenreUsecase struct {}

func NewGenreUsecase() *GenreUsecase {
	return &GenreUsecase{}
}

func (g *GenreUsecase) GetGenres() string {
	return "Genres"
}

func (g *GenreUsecase) GetGenre() string {
	return "Genre by id"
}

func (g *GenreUsecase) CreateGenre() string {
	return "Create Genre"
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