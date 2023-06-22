package cmd

import (
	"music-management/repositories"
	"music-management/usecases"
)


func Execute() {
	albumRepo := repositories.NewAlbumRepository()
	albumUC := usecases.NewAlbumUsecases(*albumRepo)

	handler := NewDelivery(*albumUC)

	handler.HandlerExecution()
}