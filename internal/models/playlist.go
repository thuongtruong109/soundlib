package models

import "github.com/thuongtruong109/soundlib/internal/tracks"

type PlayList struct {
	ID       string         `json:"id"`
	Name     string         `json:"name"`
	Tracks   []tracks.Track `json:"tracks"`
	Duration float32        `json:"duration"`
}
