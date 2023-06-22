package repositories

type AlbumRepository struct {}

func NewAlbumRepository() *AlbumRepository {
	return &AlbumRepository{}
}

func (a *AlbumRepository) GetAlbums() string {
	return "Albums"
}

func (a *AlbumRepository) GetAlbum() string {
	return "Album by id"
}

func (a *AlbumRepository) CreateAlbum() string {
	return "Create Album"
}

func (a *AlbumRepository) DeleteAlbum() string {
	return "Delete Album"
}