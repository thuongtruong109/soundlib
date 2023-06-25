package cmd

import (
	"music-management/internal/handlers"
	"music-management/internal/usecases"

	"music-management/pkg/helpers"
	"music-management/internal/artists"
	"music-management/internal/genres"
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