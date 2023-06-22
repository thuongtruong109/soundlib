package helpers

import (
	"music-management/pkg/constants"
)

func OutputColor(level int8) string {
	switch level {
	case constants.DEBUG:
		return constants.Gray
	case constants.INFO:
		return constants.Green
	case constants.WARNING:
		return constants.Yellow
	case constants.ERROR:
		return constants.Red
	case constants.FATAL:
		return constants.Purple
	case constants.LABEL:
		return constants.Cyan
	case constants.DESC:
		return constants.Yellow
	case constants.INPUT:
		return constants.Orange
	case constants.QUERY:
		return constants.Pink
	default:
		return constants.White
	}
}

func Output(level int8, msg string) {
	println(OutputColor(level) + msg + constants.Reset)
}
