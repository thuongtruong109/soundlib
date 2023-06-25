package usecases

type AlbumUsecase struct {}

func NewAlbumUsecase() *AlbumUsecase {
	return &AlbumUsecase{}
}

func (a *AlbumUsecase) GetAlbums() string {
	return "Albums"
}

func (a *AlbumUsecase) GetAlbum() string {
	return "Album by id"
}

func (a *AlbumUsecase) CreateAlbum() string {
	return "Create Album"
}

func (a *AlbumUsecase) DeleteAlbum() string {
	return "Delete Album"
}

func (a *AlbumUsecase) UpdateAlbum() string {
	return "Update Album"
}

func (a *AlbumUsecase) GetTracksOfAlbum() string {
	return "Tracks of Album"
}