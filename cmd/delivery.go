package cmd

import (
	"fmt"
	"music-management/pkg/helpers"
	"music-management/pkg/constants"
	"music-management/internal/handlers"
)

type Delivery struct {
	albumHandler handlers.AlbumHandler
	artistHandler handlers.ArtistHandler
	genreHandler handlers.GenreHandler
	playlistHandler handlers.PlaylistHandler
	songHandler handlers.SongHandler
	helper helpers.Helper
}

func NewDelivery(albumHandler handlers.AlbumHandler, artistHandler handlers.ArtistHandler, genreHandler handlers.GenreHandler, playlistHandler handlers.PlaylistHandler, songHandler handlers.SongHandler, helper helpers.Helper) *Delivery {
	return &Delivery{
		albumHandler: albumHandler,
		artistHandler: artistHandler,
		genreHandler: genreHandler,
		playlistHandler: playlistHandler,
		songHandler: songHandler,
		helper: helper,
	}
}

func (d *Delivery) DisplayOptions() {
	options := []map[int]string{
		{1: "Create new album"},
		{2: "Get all albums"},
		{3: "Get album by id"},
		{4: "Delete album by id"},
		{5: "Create new artist"},
		{6: "Get all artists"},
		{7: "Get artist by id"},
		{8: "Get all albums of artist"},
		{9: "Get all songs of artist"},
		{10: "Delete artist by id"},
		{11: "Create new genre"},
		{12: "Get all genres"},
		{13: "Get genre by id"},
		{14: "Get all albums of genre"},
		{15: "Get all songs of genre"},
		{16: "Delete genre by id"},
		{17: "Create new song"},
		{18: "Get all songs"},
		{19: "Get song by id"},
		{20: "Delete song by id"},
	}
	
	for _, option := range options {
		for key, value := range option {
			d.helper.Output(constants.DESC, fmt.Sprintf("%v", key) + ". " + value)
		}
	}
}

func (h *Delivery) HandleOption(option int) {
	switch option {
	case 1:
		h.albumHandler.CreateAlbum()
	case 2:
		h.albumHandler.GetAlbums()
	case 3:
		h.albumHandler.GetAlbum()
	case 4:
		h.albumHandler.DeleteAlbum()
	case 5:
		h.artistHandler.CreateArtist()
	case 6:
		h.artistHandler.GetArtists()
	case 7:
		h.artistHandler.GetArtist()
	case 8:
		h.artistHandler.GetAlbumsOfArtist()
	case 9:
		h.artistHandler.GetSongsOfArtist()
	case 10:
		h.artistHandler.DeleteArtist()
	case 11:
		h.genreHandler.CreateGenre()
	case 12:
		h.genreHandler.GetGenres()
	case 13:
		h.genreHandler.GetGenre()
	case 14:
		h.genreHandler.GetAlbumsOfGenre()
	case 15:
		h.genreHandler.GetSongsOfGenre()
	case 16:
		h.genreHandler.DeleteGenre()
	case 17:
		h.songHandler.CreateSong()
	case 18:
		h.songHandler.GetSongs()
	case 19:
		h.songHandler.GetSong()
	case 20:
		h.songHandler.DeleteSong()
	default:
		h.helper.Output(constants.ERROR, "Invalid option")
	}
}