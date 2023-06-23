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

	artistUC := usecases.NewArtistUsecase()
	artistHandler := handlers.NewArtistHandler(*artistUC, *helper)

	genreUC := usecases.NewGenreUsecase(*helper)
	genreHandler := handlers.NewGenreHandler(*genreUC, *helper)

	playlistUC := usecases.NewPlaylistUsecase()
	playlistHandler := handlers.NewPlaylistHandler(*playlistUC, *helper)

	songUC := usecases.NewSongUsecase()
	songHandler := handlers.NewSongHandler(*songUC, *helper)


	exe := NewDelivery(*albumHandler, *artistHandler, *genreHandler, *playlistHandler, *songHandler, *helper)
	
	exe.Execution()
}