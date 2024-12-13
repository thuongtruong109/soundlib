package tracks

import (
	"github.com/thuongtruong109/gouse"
	"github.com/thuongtruong109/soundlib/pkg/constants"
)

type TrackRepository struct{}

func NewTrackRepository() *TrackRepository {
	return &TrackRepository{}
}

func (tr *TrackRepository) GetTracks() ([]*Track, error) {
	allTracks, err := gouse.ReadFileObj[*Track](constants.TRACK_PATH)
	if err != nil {
		return nil, err
	}

	if allTracks == nil {
		return nil, nil
	}

	return allTracks, nil
}

func (tr *TrackRepository) GetTrackById(trackId string) (*Track, error) {
	allTracks, err := tr.GetTracks()
	if err != nil {
		return nil, err
	}

	if allTracks == nil {
		return nil, nil
	}

	for _, v := range allTracks {
		if v.ID == trackId {
			return v, nil
		}
	}

	return nil, nil
}

func (tr *TrackRepository) CreateTrack(newTrack *Track) (*Track, error) {
	allTracks, _ := tr.GetTracks()

	var trackInit []*Track

	if allTracks == nil {
		trackInit = make([]*Track, 0)
	} else {
		trackInit = make([]*Track, len(allTracks))
		copy(trackInit, allTracks)
	}

	trackInit = append(trackInit, newTrack)

	err := gouse.WriteFileObj(constants.TRACK_PATH, trackInit)
	if err != nil {
		return nil, err
	}

	return newTrack, nil
}

func (tr *TrackRepository) UpdateTrack(updatedTrack *Track) (*Track, error) {
	allTracks, _ := tr.GetTracks()

	if allTracks == nil {
		return nil, nil
	}

	for i, v := range allTracks {
		if v.ID == updatedTrack.ID {
			allTracks[i] = updatedTrack
			break
		}
	}

	return updatedTrack, nil
}

func (tr *TrackRepository) DeleteTrack(trackId string) error {
	allTracks, _ := tr.GetTracks()

	if allTracks == nil {
		return nil
	}

	var trackResult []*Track

	for i, v := range allTracks {
		if v.ID == trackId {
			trackResult = append(allTracks[:i], allTracks[i+1:]...)
			break
		}
	}

	err := gouse.WriteFileObj(constants.TRACK_PATH, trackResult)
	if err != nil {
		return err
	}

	return nil
}
