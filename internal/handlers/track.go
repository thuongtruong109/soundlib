package handlers

import (
	"music-management/pkg/helpers"
	
	"music-management/internal/usecases"
)

type TrackHandler struct {
	uc usecases.TrackUsecase
	helper helpers.Helper
}

func NewTrackHandler(uc usecases.TrackUsecase, helper helpers.Helper) *TrackHandler {
	return &TrackHandler{
		uc: uc,
		helper: helper,
	}
}

func (u *TrackHandler) GetTracks() {
	u.helper.OutputSuccess("GetTracks")
}

func (u *TrackHandler) GetTrack() {
	u.helper.OutputSuccess("GetTrack")
}

func (u *TrackHandler) CreateTrack() {
	u.helper.OutputSuccess("CreateTrack")
}

func (u *TrackHandler) DeleteTrack() {
	u.helper.OutputSuccess("DeleteTrack")
}

func (u *TrackHandler) UpdateTrack() {
	u.helper.OutputSuccess("UpdateTrack")
}