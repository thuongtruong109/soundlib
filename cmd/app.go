package cmd

import (
	"github.com/thuongtruong109/soundlib/internal/handlers"
	"github.com/thuongtruong109/soundlib/internal/usecases"

	"github.com/thuongtruong109/soundlib/pkg/helpers"
	"github.com/thuongtruong109/soundlib/internal/artists"
	"github.com/thuongtruong109/soundlib/internal/genres"
)

func App() {
	helper := helpers.NewHelper()

	albumUC := usecases.NewAlbumUsecase()
	albumHandler := handlers.NewAlbumHandler(*albumUC, *helper)

	artistRepo := artists.NewArtistRepository(*helper)
	artistUC := artists.NewArtistUsecase(*artistRepo, *helper)
	artistHandler := artists.NewArtistHandler(*artistUC, *helper)

	genreRepo := genres.NewGenreRepository(*helper)
	genreUC := genres.NewGenreUsecase(*genreRepo, *helper)
	genreHandler := genres.NewGenreHandler(*genreUC, *helper)

	playlistUC := usecases.NewPlaylistUsecase()
	playlistHandler := handlers.NewPlaylistHandler(*playlistUC, *helper)

	trackUC := usecases.NewTrackUsecase()
	trackHandler := handlers.NewTrackHandler(*trackUC, *helper)


	exe := NewDelivery(*albumHandler, *artistHandler, *genreHandler, *playlistHandler, *trackHandler, *helper)
	
	exe.Execution()
}