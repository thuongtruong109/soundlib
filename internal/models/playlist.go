package models

type PlayList struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Tracks []Track `json:"tracks"`
	Duration float32 `json:"duration"`
}