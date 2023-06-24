package cmd

import (
	"music-management/internal/handlers"
	"music-management/internal/usecases"

	"music-management/pkg/helpers"
)

func App() {
	helper := helpers.NewHelper()

	albumUC := usecases.NewAlbumUsecase()
	albumHandler := handlers.NewAlbumHandler(*albumUC, *helper)

	artistUC := usecases.NewArtistUsecase(*helper)
	artistHandler := handlers.NewArtistHandler(*artistUC, *helper)

	genreUC := usecases.NewGenreUsecase(*helper)
	genreHandler := handlers.NewGenreHandler(*genreUC, *helper)

	playlistUC := usecases.NewPlaylistUsecase()
	playlistHandler := handlers.NewPlaylistHandler(*playlistUC, *helper)

	trackUC := usecases.NewTrackUsecase()
	trackHandler := handlers.NewTrackHandler(*trackUC, *helper)


	exe := NewDelivery(*albumHandler, *artistHandler, *genreHandler, *playlistHandler, *trackHandler, *helper)
	
	exe.Execution()
}