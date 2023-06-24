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
	trackHandler handlers.TrackHandler
	helper helpers.Helper
}

func NewDelivery(albumHandler handlers.AlbumHandler, artistHandler handlers.ArtistHandler, genreHandler handlers.GenreHandler, playlistHandler handlers.PlaylistHandler, trackHandler handlers.TrackHandler, helper helpers.Helper) *Delivery {
	return &Delivery{
		albumHandler: albumHandler,
		artistHandler: artistHandler,
		genreHandler: genreHandler,
		playlistHandler: playlistHandler,
		trackHandler: trackHandler,
		helper: helper,
	}
}

func (d *Delivery) DisplayOptions() {
	options := []map[int]string{
		{1: "Create new album"},
		{2: "Get all albums"},
		{3: "Get album info by id"},
		{4: "Delete album by id"},
		{5: "Update album by id"},
		{6: "Get all tracks of album"},

		{7: "Create new artist"},
		{8: "Get all artists"},
		{9: "Get artist by id"},
		{10: "Get all albums of artist"},
		{11: "Get all songs of artist"},
		{12: "Delete artist by id"},

		{13: "Create new genre"},
		{14: "Get all genres"},
		{15: "Get genre by id"},
		{16: "Delete genre by id"},
		{17: "Update genre by id"},
		{18: "Get all tracks of genre"},

		{19: "Create new track"},
		{20: "Get all tracks"},
		{21: "Get track by id"},
		{22: "Delete track by id"},
		{23: "Update track by id"},

		{24: "Create new playlist"},
		{25: "Get all playlists"},
		{26: "Get playlist by id"},
		{27: "Delete playlist by id"},
		{28: "Update playlist by id"},

		{29: "Add track to playlist"},
		{30: "Remove track from playlist"},
		{31: "Get all tracks of playlist"},
		{32: "Get all playlists have track"},
		{33: "Get all playlists have track"},
	}
	
	for _, option := range options {
		for key, value := range option {
			d.helper.OutputNomal(constants.DESC, fmt.Sprintf("⦿ %v", key) + ". " + value)
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
		h.albumHandler.UpdateAlbum()
	case 6:
		h.albumHandler.GetTracksOfAlbum()

	case 7:
		h.artistHandler.CreateArtist()
	case 8:
		h.artistHandler.GetArtists()
	case 9:
		h.artistHandler.GetArtist()
	case 10:
		h.artistHandler.GetAlbumsOfArtist()
	case 11:
		h.artistHandler.GetTracksOfArtist()
	case 12:
		h.artistHandler.DeleteArtist()

	case 13:
		h.genreHandler.CreateGenre()
	case 14:
		h.genreHandler.GetGenres()
	case 15:
		h.genreHandler.GetGenre()
	case 16:
		h.genreHandler.DeleteGenre()
	case 17:
		h.genreHandler.UpdateGenre()
	case 18:
		h.genreHandler.GetTracksOfGenre()

	case 19:
		h.trackHandler.CreateTrack()
	case 20:
		h.trackHandler.GetTracks()
	case 21:
		h.trackHandler.GetTrack()
	case 22:
		h.trackHandler.DeleteTrack()
	case 23:
		h.trackHandler.UpdateTrack()

	case 24:
		h.playlistHandler.CreatePlaylist()
	case 25:
		h.playlistHandler.GetPlaylists()
	case 26:
		h.playlistHandler.GetPlaylist()
	case 27:
		h.playlistHandler.DeletePlaylist()
	case 28:
		h.playlistHandler.UpdatePlaylist()

	case 29:
		h.playlistHandler.AddTrackToPlaylist()
	case 30:
		h.playlistHandler.DeleteTrackFromPlaylist()
	case 31:
		h.playlistHandler.GetTracksOfPlaylist()
	case 32:
		h.playlistHandler.GetPlaylistsHaveTrack()
	default:
		h.helper.OutputNomal(constants.ERROR, "Invalid option")
	}
}