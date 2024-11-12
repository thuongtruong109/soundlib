package cmd

import (
	"github.com/thuongtruong109/soundlib/internal/handlers"
	"github.com/thuongtruong109/soundlib/internal/usecases"

	"github.com/thuongtruong109/soundlib/internal/albums"
	"github.com/thuongtruong109/soundlib/internal/artists"
	"github.com/thuongtruong109/soundlib/internal/common"
	"github.com/thuongtruong109/soundlib/internal/genres"

	"github.com/thuongtruong109/soundlib/pkg/helpers"
)

func App() {
	helper := helpers.NewHelper()

	albumUC := albums.NewAlbumUsecase()
	albumHandler := albums.NewAlbumHandler(*albumUC, *helper)

	artistRepo := artists.NewArtistRepository()
	artistUC := artists.NewArtistUsecase(*artistRepo, *helper)
	artistHandler := artists.NewArtistHandler(*artistUC, *helper, *common.NewCommonHandler(*helper, "Artists"))

	genreRepo := genres.NewGenreRepository(*helper)
	genreUC := genres.NewGenreUsecase(*genreRepo, *helper)
	genreHandler := genres.NewGenreHandler(*genreUC, *helper, *common.NewCommonHandler(*helper, "Genres"))

	playlistUC := usecases.NewPlaylistUsecase()
	playlistHandler := handlers.NewPlaylistHandler(*playlistUC, *helper)

	trackUC := usecases.NewTrackUsecase()
	trackHandler := handlers.NewTrackHandler(*trackUC, *helper)

	exe := NewDelivery(*albumHandler, *artistHandler, *genreHandler, *playlistHandler, *trackHandler, *helper)

	exe.Execution()
}
