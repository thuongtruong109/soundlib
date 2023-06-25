package usecases

type TrackUsecase struct {}

func NewTrackUsecase() *TrackUsecase {
	return &TrackUsecase{}
}

func (s *TrackUsecase) GetSongs() string {
	return "Songs"
}

func (s *TrackUsecase) GetSong() string {
	return "Song by id"
}

func (s *TrackUsecase) CreateSong() string {
	return "Create Song"
}

func (s *TrackUsecase) DeleteSong() string {
	return "Delete Song"
}

func (s *TrackUsecase) UpdateSong() string {
	return "Update Song"
}