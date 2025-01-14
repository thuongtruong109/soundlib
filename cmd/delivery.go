package cmd

import (
	"github.com/thuongtruong109/soundlib/internal/playlists"
	"github.com/thuongtruong109/soundlib/internal/tracks"
	"github.com/thuongtruong109/soundlib/pkg/constants"
	"github.com/thuongtruong109/soundlib/pkg/helpers"

	"github.com/thuongtruong109/soundlib/internal/albums"
	"github.com/thuongtruong109/soundlib/internal/artists"
	"github.com/thuongtruong109/soundlib/internal/genres"

	"github.com/thuongtruong109/soundlib/internal/track_in_album"
	"github.com/thuongtruong109/soundlib/internal/track_in_playlist"
)

type Delivery struct {
	helper          helpers.Helper
	albumHandler    albums.AlbumHandler
	artistHandler   artists.ArtistHandler
	genreHandler    genres.GenreHandler
	playlistHandler playlists.PlaylistHandler
	trackHandler    tracks.TrackHandler

	trackInAlbumHandler    track_in_album.TrackInAlbumHandler
	trackInPlaylistHandler track_in_playlist.TrackInPlaylistHandler
}

func NewDelivery(helper helpers.Helper, albumHandler albums.AlbumHandler, artistHandler artists.ArtistHandler, genreHandler genres.GenreHandler, playlistHandler playlists.PlaylistHandler, trackHandler tracks.TrackHandler, trackInAlbumHandler track_in_album.TrackInAlbumHandler, trackInPlaylistHandler track_in_playlist.TrackInPlaylistHandler) *Delivery {
	return &Delivery{
		helper:          helper,
		albumHandler:    albumHandler,
		artistHandler:   artistHandler,
		genreHandler:    genreHandler,
		playlistHandler: playlistHandler,
		trackHandler:    trackHandler,

		trackInAlbumHandler:    trackInAlbumHandler,
		trackInPlaylistHandler: trackInPlaylistHandler,
	}
}

func (h *Delivery) HandleOption(option int8) {
	optionHandlers := map[int8]func(){
		1: h.albumHandler.CreateAlbum,
		2: h.albumHandler.GetAlbums,
		3: h.albumHandler.GetAlbum,
		4: h.albumHandler.DeleteAlbum,
		5: h.albumHandler.UpdateAlbum,

		6: h.trackInAlbumHandler.AddTrackToAlbum,
		7: h.trackInAlbumHandler.GetTracksOfAlbum,
		8: h.trackInAlbumHandler.DeleteTrackFromAlbum,

		9:  h.artistHandler.CreateArtist,
		10: h.artistHandler.GetArtists,
		11: h.artistHandler.GetArtist,
		12: h.artistHandler.DeleteArtist,
		13: h.artistHandler.UpdateArtist,

		14: h.genreHandler.CreateGenre,
		15: h.genreHandler.GetGenres,
		16: h.genreHandler.GetGenre,
		17: h.genreHandler.DeleteGenre,
		18: h.genreHandler.UpdateGenre,

		19: h.trackHandler.CreateTrack,
		20: h.trackHandler.GetTracks,
		21: h.trackHandler.GetTrack,
		22: h.trackHandler.DeleteTrack,
		23: h.trackHandler.UpdateTrack,

		24: h.playlistHandler.CreatePlaylist,
		25: h.playlistHandler.GetPlaylists,
		26: h.playlistHandler.GetPlaylist,
		27: h.playlistHandler.DeletePlaylist,
		28: h.playlistHandler.UpdatePlaylist,

		29: h.trackInPlaylistHandler.AddTrackToPlaylist,
		30: h.trackInPlaylistHandler.GetTracksOfPlaylist,
		31: h.trackInPlaylistHandler.DeleteTrackFromPlaylist,
	}

	handler, exists := optionHandlers[option]
	if exists {
		handler()
	} else {
		h.helper.OutputDefault(constants.ERROR, "Invalid option")
	}
}
