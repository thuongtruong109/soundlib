package models

import (
	"music-management/pkg/helpers"
)

type PlayList struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Tracks []Track `json:"tracks"`
	Duration float32 `json:"duration"`
}

func NewPlayList(helper helpers.Helper) *PlayList {
	return &PlayList{
		ID: helper.GenerateID(),
	}
}