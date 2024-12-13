package cmd

import (
	"github.com/thuongtruong109/soundlib/internal/albums"
	"github.com/thuongtruong109/soundlib/internal/artists"
	"github.com/thuongtruong109/soundlib/internal/genres"
	"github.com/thuongtruong109/soundlib/internal/playlists"
	"github.com/thuongtruong109/soundlib/internal/tracks"

	"github.com/thuongtruong109/soundlib/pkg/common"
	"github.com/thuongtruong109/soundlib/pkg/helpers"
)

func App() {
	helper := helpers.NewHelper()

	albumUC := albums.NewAlbumUsecase()
	albumHandler := albums.NewAlbumHandler(*albumUC, *helper)

	artistRepo := artists.NewArtistRepository()
	artistUC := artists.NewArtistUsecase(*artistRepo, *helper)
	artistHandler := artists.NewArtistHandler(*artistUC, *helper, *common.NewCommonHandler(*helper, "Artists"))

	genreRepo := genres.NewGenreRepository()
	genreUC := genres.NewGenreUsecase(*genreRepo, *helper)
	genreHandler := genres.NewGenreHandler(*genreUC, *helper, *common.NewCommonHandler(*helper, "Genres"))

	playlistUC := playlists.NewPlaylistUsecase()
	playlistHandler := playlists.NewPlaylistHandler(*playlistUC, *helper)

	trackRepo := tracks.NewTrackRepository()
	trackUC := tracks.NewTrackUsecase(*trackRepo, *helper)
	trackHandler := tracks.NewTrackHandler(*trackUC, *helper, *common.NewCommonHandler(*helper, "Tracks"))

	exe := NewDelivery(*helper, *albumHandler, *artistHandler, *genreHandler, *playlistHandler, *trackHandler)
	/**trackHandler)*/

	exe.Execution()
}
