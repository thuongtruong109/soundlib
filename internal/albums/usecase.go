package albums

import (
	"fmt"

	"github.com/thuongtruong109/gouse"
	"github.com/thuongtruong109/soundlib/pkg/helpers"
)

type AlbumUsecase struct {
	repo   AlbumRepository
	helper helpers.Helper
}

func NewAlbumUsecase(repo AlbumRepository, helper helpers.Helper) *AlbumUsecase {
	return &AlbumUsecase{
		repo:   repo,
		helper: helper,
	}
}

func (a *AlbumUsecase) GetAlbums() ([]string, string, error) {
	result, time, err := helpers.QueryTimeTwoOutput[[]*Album](a.repo.GetAlbums)()
	if err != nil {
		return nil, time, err
	}

	if result == nil {
		return nil, time, fmt.Errorf(gouse.DESC_NOT_FOUND_DATA)
	}

	var output []string
	for _, v := range result {
		output = append(output, fmt.Sprintf("ID: %s, Name: %s, Cover: %s, CreatedAt: %s, Owner ID: %s", v.ID, v.Name, v.Cover, v.CreatedAt, v.OwnerID))
	}

	return output, time, nil
}

func (a *AlbumUsecase) GetAlbum() ([]string, string, error) {
	var id string
	fmt.Print("» Enter ID: ")
	fmt.Scanln(&id)

	result, time, err := helpers.QueryTimeTwoOutputWithParams[*Album, string](a.repo.GetAlbum)(id)
	if err != nil {
		return nil, time, err
	}

	if result == nil {
		return nil, time, fmt.Errorf(gouse.DESC_NOT_FOUND_DATA)
	}

	output := []string{fmt.Sprintf("ID: %s, Name: %s, Cover: %s, CreatedAt: %s, Owner ID: %s", result.ID, result.Name, result.Cover, result.CreatedAt, result.OwnerID)}

	return output, time, nil
}

func (a *AlbumUsecase) CreateAlbum() ([]string, string, error) {
	var name, cover, ownerId string
	fmt.Print("» Enter name: ")
	fmt.Scanln(&name)

	fmt.Print("» Enter cover url: ")
	fmt.Scanln(&cover)

	fmt.Print("» Enter owner id: ")
	fmt.Scanln(&ownerId)

	album := &Album{
		ID:        gouse.UUID(),
		Name:  name,
		Cover: cover,
		CreatedAt: gouse.ISODate(),
		OwnerID: ownerId,
	}

	result, time, err := helpers.QueryTimeTwoOutputWithParams[*Album, *Album](a.repo.CreateAlbum)(album)
	if err != nil {
		return nil, time, err
	}

	if result == nil {
		return nil, time, fmt.Errorf(gouse.DESC_NOT_FOUND_DATA)
	}

	output := []string{fmt.Sprintf("ID: %s", result.ID)}

	return output, time, nil
}

func (a *AlbumUsecase) DeleteAlbum() (string, error) {
	var id string
	fmt.Print("» Enter ID: ")
	fmt.Scanln(&id)

	time, err := helpers.QueryTimeErrorWithOneParam[error](a.repo.DeleteAlbum)(id)
	if err != nil {
		return time, err
	}

	return time, nil
}

func (a *AlbumUsecase) UpdateAlbum() ([]string, string, error) {
	var id, name, cover, owner_id string
	fmt.Print("» Enter ID: ")
	fmt.Scanln(&id)

	fmt.Print("» Enter name: ")
	fmt.Scanln(&name)

	fmt.Print("» Enter cover url: ")
	fmt.Scanln(&cover)

	fmt.Print("» Enter owner id: ")
	fmt.Scanln(&owner_id)

	newAlbum := &Album{
		ID:        id,
		Name:  name,
		Cover: cover,
		CreatedAt: gouse.ISODate(),
		OwnerID: owner_id,
	}

	result, time, err := helpers.QueryTimeTwoOutputWithParams[*Album, *Album](a.repo.UpdateAlbum)(newAlbum)
	if err != nil {
		return nil, time, err
	}

	if result == nil {
		return nil, time, fmt.Errorf(gouse.DESC_NOT_FOUND_DATA)
	}

	output := []string{fmt.Sprintf("ID: %s", result.ID)}

	return output, time, nil
}