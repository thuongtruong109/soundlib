package track_in_playlist

import (
	"fmt"

	"github.com/thuongtruong109/gouse"
)

type TrackInPlaylistRepository struct{
	dbPath string
}

func NewTrackInPlaylistRepository(dbPath string) *TrackInPlaylistRepository {
	return &TrackInPlaylistRepository{
		dbPath: dbPath,
	}
}

func (r *TrackInPlaylistRepository) GetTracksOfPlaylist() ([]*TrackInPlaylist, error) {
	tracks, err := gouse.ReadFileObj[*TrackInPlaylist](r.dbPath)

	if err != nil {
		return nil, err
	}

	if tracks == nil {
		return nil, nil
	}

	return tracks, nil
}

func (r *TrackInPlaylistRepository) AddTrackToAlbum(newTrack *TrackInPlaylist) (*TrackInPlaylist, error) {
	allGenres, _ := r.GetTracksOfPlaylist()

	var genresInit []*TrackInPlaylist

	if allGenres == nil {
		genresInit = make([]*TrackInPlaylist, 0)
	} else {
		genresInit = make([]*TrackInPlaylist, len(allGenres))
		copy(genresInit, allGenres)
	}

	genresInit = append(genresInit, newTrack)

	err2 := gouse.WriteFileObj(r.dbPath, genresInit)
	if err2 != nil {
		return nil, fmt.Errorf(gouse.DESC_CREATE_FAILED)
	}
	return newTrack, nil
}

func (r *TrackInPlaylistRepository) DeleteTrackFromAlbum(trackID, playlistID string) error {
	allTracks, _ := r.GetTracksOfPlaylist()

	if allTracks == nil {
		return nil
	}

	var tracksInit []*TrackInPlaylist

	for i, track := range allTracks {
		if track.TrackID == trackID && track.PlaylistID == playlistID {
			tracksInit = append(allTracks[:i], allTracks[i+1:]...)
			break
		}
	}

	err2 := gouse.WriteFileObj[[]*TrackInPlaylist](r.dbPath, tracksInit)
	if err2 != nil {
		return fmt.Errorf(gouse.DESC_DELETE_FAILED)
	}

	return nil
}
