package cmd

import (
	"github.com/thuongtruong109/soundlib/internal/albums"
	"github.com/thuongtruong109/soundlib/internal/artists"
	"github.com/thuongtruong109/soundlib/internal/genres"
	"github.com/thuongtruong109/soundlib/internal/playlists"
	"github.com/thuongtruong109/soundlib/internal/tracks"

	"github.com/thuongtruong109/soundlib/internal/track_in_album"
	"github.com/thuongtruong109/soundlib/internal/track_in_playlist"

	"github.com/thuongtruong109/soundlib/pkg/base"
	"github.com/thuongtruong109/soundlib/pkg/constants"
	"github.com/thuongtruong109/soundlib/pkg/helpers"
)

func App() {
	helper := helpers.NewHelper()

	albumRepo := albums.NewAlbumRepository(constants.ALBUM_PATH)
	albumUC := albums.NewAlbumUsecase(*albumRepo, *helper)
	albumHandler := albums.NewAlbumHandler(*albumUC, *helper, *base.NewBaseHandler(*helper, "Albums"))

	artistRepo := artists.NewArtistRepository(constants.ARTIST_PATH)
	artistUC := artists.NewArtistUsecase(*artistRepo, *helper)
	artistHandler := artists.NewArtistHandler(*artistUC, *helper, *base.NewBaseHandler(*helper, "Artists"))

	genreRepo := genres.NewGenreRepository(constants.GENRE_PATH)
	genreUC := genres.NewGenreUsecase(*genreRepo, *helper)
	genreHandler := genres.NewGenreHandler(*genreUC, *helper, *base.NewBaseHandler(*helper, "Genres"))

	playlistRepo := playlists.NewPlaylistRepository(constants.PLAYLIST_PATH)
	playlistUC := playlists.NewPlaylistUsecase(*playlistRepo, *helper)
	playlistHandler := playlists.NewPlaylistHandler(*playlistUC, *helper, *base.NewBaseHandler(*helper, "Playlists"))

	trackRepo := tracks.NewTrackRepository(constants.TRACK_PATH)
	trackUC := tracks.NewTrackUsecase(*trackRepo, *helper)
	trackHandler := tracks.NewTrackHandler(*trackUC, *helper, *base.NewBaseHandler(*helper, "Tracks"))

	trackInAlbumRepo := track_in_album.NewTrackInAlbumRepository(constants.TRACK_IN_ALBUM_PATH)
	trackInAlbumUC := track_in_album.NewTrackInAlbumUsecase(*trackInAlbumRepo, *helper)
	trackInAlbumHandler := track_in_album.NewTrackInAlbumHandler(*trackInAlbumUC, *helper, *base.NewBaseHandler(*helper, "Tracks in Album"))

	trackInPlaylistRepo := track_in_playlist.NewTrackInPlaylistRepository(constants.TRACK_IN_PLAYLIST_PATH)
	trackInPlaylistUC := track_in_playlist.NewTrackInPlaylistUsecase(*trackInPlaylistRepo, *helper)
	trackInPlaylistHandler := track_in_playlist.NewTrackInPlaylistHandler(*trackInPlaylistUC, *helper, *base.NewBaseHandler(*helper, "Tracks in Playlist"))

	exe := NewDelivery(*helper, *albumHandler, *artistHandler, *genreHandler, *playlistHandler, *trackHandler, *trackInAlbumHandler, *trackInPlaylistHandler)

	exe.Execution()
}
