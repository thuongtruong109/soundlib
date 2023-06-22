package models

import (
	"music-management/pkg/helpers"
)

type PlayList struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Tracks []Track `json:"tracks"`
	Duration float32 `json:"duration"`
}

func NewPlayList() *PlayList {
	return &PlayList{
		ID: helpers.GenerateID(),
		Name: "",
		Tracks: []Track{},
		Duration: 0,
	}
}