package track_in_album

import (
	"fmt"

	"github.com/thuongtruong109/gouse"
)

type TrackInAlbumRepository struct{
	dbPath string
}

func NewTrackInAlbumRepository(dbPath string) *TrackInAlbumRepository {
	return &TrackInAlbumRepository{
		dbPath: dbPath,
	}
}

func (r *TrackInAlbumRepository) GetTracksOfAlbum() ([]*TrackInAlbum, error) {
	tracks, err := gouse.ReadFileObj[*TrackInAlbum](r.dbPath)

	if err != nil {
		return nil, err
	}

	if tracks == nil {
		return nil, nil
	}

	return tracks, nil
}

func (r *TrackInAlbumRepository) AddTrackToAlbum(newTrack *TrackInAlbum) (*TrackInAlbum, error) {
	allGenres, _ := r.GetTracksOfAlbum()

	var genresInit []*TrackInAlbum

	if allGenres == nil {
		genresInit = make([]*TrackInAlbum, 0)
	} else {
		genresInit = make([]*TrackInAlbum, len(allGenres))
		copy(genresInit, allGenres)
	}

	genresInit = append(genresInit, newTrack)

	err2 := gouse.WriteFileObj(r.dbPath, genresInit)
	if err2 != nil {
		return nil, fmt.Errorf(gouse.DESC_CREATE_FAILED)
	}
	return newTrack, nil
}

func (r *TrackInAlbumRepository) DeleteTrackFromAlbum(trackID, albumID string) error {
	allTracks, _ := r.GetTracksOfAlbum()

	if allTracks == nil {
		return nil
	}

	var tracksInit []*TrackInAlbum

	for i, track := range allTracks {
		if track.TrackID == trackID && track.AlbumID == albumID {
			tracksInit = append(allTracks[:i], allTracks[i+1:]...)
			break
		}
	}

	err2 := gouse.WriteFileObj(r.dbPath, tracksInit)
	if err2 != nil {
		return fmt.Errorf(gouse.DESC_DELETE_FAILED)
	}

	return nil
}
