package playlists

import (
	"fmt"

	"github.com/thuongtruong109/gouse"
	"github.com/thuongtruong109/soundlib/pkg/helpers"
)

type PlaylistUsecase struct {
	repo   PlaylistRepository
	helper helpers.Helper
}

func NewPlaylistUsecase(repo PlaylistRepository, helper helpers.Helper) *PlaylistUsecase {
	return &PlaylistUsecase{
		repo:   repo,
		helper: helper,
	}
}

func (a *PlaylistUsecase) GetPlaylists() ([]string, string, error) {
	result, time, err := helpers.QueryTimeTwoOutput[[]*Playlist](a.repo.GetPlaylists)()
	if err != nil {
		return nil, time, err
	}

	if result == nil {
		return nil, time, fmt.Errorf(gouse.DESC_NOT_FOUND_DATA)
	}

	var output []string
	for _, v := range result {
		output = append(output, fmt.Sprintf("ID: %s, Name: %s, Description: %s, Duration: %s CreatedAt: %s, Creator ID: %s", v.ID, v.Name, v.Description, v.Duration, v.CreatedAt, v.CreatorID))
	}

	return output, time, nil
}

func (a *PlaylistUsecase) GetPlaylist() ([]string, string, error) {
	var id string
	fmt.Print("» Enter ID: ")
	fmt.Scanln(&id)

	result, time, err := helpers.QueryTimeTwoOutputWithParams[*Playlist, string](a.repo.GetPlaylist)(id)
	if err != nil {
		return nil, time, err
	}

	if result == nil {
		return nil, time, fmt.Errorf(gouse.DESC_NOT_FOUND_DATA)
	}

	output := []string{fmt.Sprintf("ID: %s, Name: %s, Description: %s, Duration: %s CreatedAt: %s, Creator ID: %s", result.ID, result.Name, result.Description, result.Duration, result.CreatedAt, result.CreatorID)}

	return output, time, nil
}

func (a *PlaylistUsecase) CreatePlaylist() ([]string, string, error) {
	var name, description, duration, creatorId string
	fmt.Print("» Enter name: ")
	fmt.Scanln(&name)

	fmt.Print("» Enter description: ")
	fmt.Scanln(&description)

	fmt.Print("» Enter duration: ")
	fmt.Scanln(&duration)

	fmt.Print("» Enter creator id: ")
	fmt.Scanln(&creatorId)

	playlist := &Playlist{
		ID:        gouse.UUID(),
		Name:  name,
		Description: description,
		Duration: duration,
		CreatedAt: gouse.ISODate(),
		CreatorID: creatorId,
	}

	result, time, err := helpers.QueryTimeTwoOutputWithParams[*Playlist, *Playlist](a.repo.CreatePlaylist)(playlist)
	if err != nil {
		return nil, time, err
	}

	if result == nil {
		return nil, time, fmt.Errorf(gouse.DESC_NOT_FOUND_DATA)
	}

	output := []string{fmt.Sprintf("ID: %s", result.ID)}

	return output, time, nil
}

func (a *PlaylistUsecase) DeletePlaylist() (string, error) {
	var id string
	fmt.Print("» Enter ID: ")
	fmt.Scanln(&id)

	time, err := helpers.QueryTimeErrorWithOneParam[error](a.repo.DeletePlaylist)(id)
	if err != nil {
		return time, err
	}

	return time, nil
}

func (a *PlaylistUsecase) UpdatePlaylist() ([]string, string, error) {
	var id, name, description, duration, creatorId string
	fmt.Print("» Enter name: ")
	fmt.Scanln(&name)

	fmt.Print("» Enter description: ")
	fmt.Scanln(&description)

	fmt.Print("» Enter duration: ")
	fmt.Scanln(&duration)

	fmt.Print("» Enter creator id: ")
	fmt.Scanln(&creatorId)

	newPlaylist := &Playlist{
		ID:        id,
		Name:  name,
		Description: description,
		Duration: duration,
		CreatedAt: gouse.ISODate(),
		CreatorID: creatorId,
	}

	result, time, err := helpers.QueryTimeTwoOutputWithParams[*Playlist, *Playlist](a.repo.UpdatePlaylist)(newPlaylist)
	if err != nil {
		return nil, time, err
	}

	if result == nil {
		return nil, time, fmt.Errorf(gouse.DESC_NOT_FOUND_DATA)
	}

	output := []string{fmt.Sprintf("ID: %s", result.ID)}

	return output, time, nil
}