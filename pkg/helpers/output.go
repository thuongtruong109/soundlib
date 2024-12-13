package helpers

import (
	"github.com/thuongtruong109/gouse"
	"github.com/thuongtruong109/soundlib/pkg/constants"
)

func (h *Helper) OutputColor(level int8) string {
	switch level {
	case constants.DEBUG:
		return constants.Gray
	case constants.INFO:
		return gouse.GREEN_FG
	case constants.WARNING:
		return gouse.ORANGE_FG
	case constants.ERROR:
		return gouse.RED_FG
	case constants.FATAL:
		return gouse.PURPLE_FG
	case constants.LABEL:
		return gouse.CYAN_FG
	case constants.DESC:
		return gouse.YELLOW_FG
	case constants.INPUT:
		return gouse.CYAN_FG
	case constants.QUERY:
		return gouse.PINK_FG
	default:
		return gouse.WHITE_FG
	}
}

func (h *Helper) OutputSuccess(statusMsg string) {
	println(gouse.GREEN_FG + "\n::: Status: " + statusMsg + gouse.DEFAULT_FG)
}

func (h *Helper) OutputError(statusMsg string, err string) {
	println(gouse.RED_FG + "\n::: Status: " + statusMsg + "\n::: Message: " + err + gouse.DEFAULT_FG)
}

func (h *Helper) OutputNomal(label int8, msg string) {
	println(h.OutputColor(label) + msg + gouse.DEFAULT_FG)
}

func OutputTime(time string) {
	println(gouse.PURPLE_FG + "\n::: Query time: " + time + gouse.DEFAULT_FG)
}