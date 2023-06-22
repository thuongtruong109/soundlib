package usecases

type SongUsecase struct {}

func NewSongUsecase() *SongUsecase {
	return &SongUsecase{}
}

func (s *SongUsecase) GetSongs() string {
	return "Songs"
}

func (s *SongUsecase) GetSong() string {
	return "Song by id"
}

func (s *SongUsecase) CreateSong() string {
	return "Create Song"
}

func (s *SongUsecase) DeleteSong() string {
	return "Delete Song"
}

func (s *SongUsecase) UpdateSong() string {
	return "Update Song"
}