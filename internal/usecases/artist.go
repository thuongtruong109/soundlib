package usecases

type ArtistUsecase struct {}

func NewArtistUsecase() *ArtistUsecase {
	return &ArtistUsecase{}
}

func (a *ArtistUsecase) GetArtists() string {
	return "Artists"
}

func (a *ArtistUsecase) GetArtist() string {
	return "Artist by id"
}

func (a *ArtistUsecase) CreateArtist() string {
	return "Create Artist"
}

func (a *ArtistUsecase) DeleteArtist() string {
	return "Delete Artist"
}

func (a *ArtistUsecase) UpdateArtist() string {
	return "Update Artist"
}

func (a *ArtistUsecase) GetAlbumsOfArtist() string {
	return "Albums of Artist"
}

func (a *ArtistUsecase) GetTracksOfArtist() string {
	return "Tracks of Artist"
}