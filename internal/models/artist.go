package models

import (
	"music-management/pkg/helpers"
)

type Artist struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Albums []Album `json:"albums"`
}

func NewArtist(helper helpers.Helper) *Artist {
	return &Artist{
		ID: helper.GenerateID(),
	}
}