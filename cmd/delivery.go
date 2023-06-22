package cmd

import (
	"music-management/pkg/helpers"
	"music-management/pkg/constants"

	"music-management/usecases"
)

type Delivery struct {
	albumUC usecases.AlbumUsecases
}

func NewDelivery(albumUC usecases.AlbumUsecases) *Delivery {
	return &Delivery{
		albumUC: albumUC,
	}
}

func (h *Delivery) DisplayOptions() {
	helpers.Output(constants.DESC, "1. Create new album")
	helpers.Output(constants.DESC, "2. Get all albums")
	helpers.Output(constants.DESC, "3. Get album by id")
	helpers.Output(constants.DESC, "4. Delete album by id")

	helpers.Output(constants.DESC, "5. Create new artist")
	helpers.Output(constants.DESC, "6. Get all artists")
	helpers.Output(constants.DESC, "7. Get artist by id")
	helpers.Output(constants.DESC, "8. Get all albums of artist")
	helpers.Output(constants.DESC, "9. Get all songs of artist")
	helpers.Output(constants.DESC, "10. Delete artist by id")

	helpers.Output(constants.DESC, "11. Create new genre")
	helpers.Output(constants.DESC, "12. Get all genres")
	helpers.Output(constants.DESC, "13. Get genre by id")
	helpers.Output(constants.DESC, "14. Get all albums of genre")
	helpers.Output(constants.DESC, "15. Get all songs of genre")
	helpers.Output(constants.DESC, "16. Delete genre by id")

	helpers.Output(constants.DESC, "17. Create new song")
	helpers.Output(constants.DESC, "18. Get all songs")
	helpers.Output(constants.DESC, "19. Get song by id")
	helpers.Output(constants.DESC, "20. Delete song by id")
}

func (h *Delivery) HandleOption(option int) {
	switch option {
	case 1:
		h.albumUC.CreateAlbum()
	case 2:
		h.albumUC.GetAlbums()
	case 3:
		h.albumUC.GetAlbum()
	case 4:
		h.albumUC.DeleteAlbum()
	case 5:
		helpers.Output(constants.DESC, "\n ::: Creating new artist")
	case 6:
		helpers.Output(constants.DESC, "\n ::: Getting all artists")
	case 7:
		helpers.Output(constants.DESC, "\n ::: Getting artist by id")
	case 8:
		helpers.Output(constants.DESC, "\n ::: Getting all albums of artist")
	case 9:
		helpers.Output(constants.DESC, "\n ::: Getting all songs of artist")
	case 10:
		helpers.Output(constants.DESC, "\n ::: Deleting artist by id")
	case 11:
		helpers.Output(constants.DESC, "\n ::: Creating new genre")
	case 12:
		helpers.Output(constants.DESC, "\n ::: Getting all genres")
	case 13:
		helpers.Output(constants.DESC, "\n ::: Getting genre by id")
	case 14:
		helpers.Output(constants.DESC, "\n ::: Getting all albums of genre")
	case 15:
		helpers.Output(constants.DESC, "\n ::: Getting all songs of genre")
	case 16:
		helpers.Output(constants.DESC, "\n ::: Deleting genre by id")
	case 17:
		helpers.Output(constants.DESC, "\n ::: Creating new song")
	case 18:
		helpers.Output(constants.DESC, "\n ::: Getting all songs")
	case 19:
		helpers.Output(constants.DESC, "\n ::: Getting song by id")
	case 20:
		helpers.Output(constants.DESC, "\n ::: Deleting song by id")
	default:
		helpers.Output(constants.DESC, "\n ::: Invalid option\n")
	}
}