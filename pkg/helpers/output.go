package helpers

import (
	"github.com/thuongtruong109/soundlib/pkg/constants"
	"github.com/thuongtruong109/gouse/shared"
)

func (h *Helper) OutputColor(level int8) string {
	switch level {
	case constants.DEBUG:
		return constants.Gray
	case constants.INFO:
		return shared.GREEN_FG
	case constants.WARNING:
		return shared.ORANGE_FG
	case constants.ERROR:
		return shared.RED_FG
	case constants.FATAL:
		return shared.PURPLE_FG
	case constants.LABEL:
		return shared.CYAN_FG
	case constants.DESC:
		return shared.YELLOW_FG
	case constants.INPUT:
		return shared.CYAN_FG
	case constants.QUERY:
		return shared.PINK_FG
	default:
		return shared.WHITE_FG
	}
}

func (h *Helper) OutputSuccess(statusMsg string) {
	println(shared.GREEN_FG + "\n::: Status: " + statusMsg + shared.DEFAULT_FG)
}

func (h *Helper) OutputError(statusMsg string, err string) {
	println(shared.RED_FG + "\n::: Status: " + statusMsg + "\n::: Message: " + err + shared.DEFAULT_FG)
}

func (h *Helper) OutputNomal(label int8, msg string) {
	println(h.OutputColor(label) + msg + shared.DEFAULT_FG)
}

func OutputTime(time string) {
	println(shared.PURPLE_FG + "\n::: Query time: " + time + shared.DEFAULT_FG)
}