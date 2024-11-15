package playlists

type PlaylistUsecase struct{}

func NewPlaylistUsecase() *PlaylistUsecase {
	return &PlaylistUsecase{}
}

func (p *PlaylistUsecase) GetPlaylists() string {
	return "Playlists"
}

func (p *PlaylistUsecase) GetPlaylist() string {
	return "Playlist by id"
}

func (p *PlaylistUsecase) CreatePlaylist() string {
	return "Create Playlist"
}

func (p *PlaylistUsecase) DeletePlaylist() string {
	return "Delete Playlist"
}

func (p *PlaylistUsecase) UpdatePlaylist() string {
	return "Update Playlist"
}

func (p *PlaylistUsecase) AddSongToPlaylist() string {
	return "Add Song to Playlist"
}

func (p *PlaylistUsecase) RemoveSongFromPlaylist() string {
	return "Remove Song from Playlist"
}

func (p *PlaylistUsecase) GetSongsFromPlaylist() string {
	return "Get Songs from Playlist"
}
