package helpers

import (
	"github.com/thuongtruong109/gouse"
	"github.com/thuongtruong109/soundlib/pkg/constants"
)

type Helper struct {}

func NewHelper() *Helper {
	return &Helper{}
}

func (h *Helper) OutputColor(level int8) string {
	switch level {
	case constants.DEBUG:
		return gouse.GRAY_CONSOLE
	case constants.INFO:
		return gouse.GREEN_CONSOLE
	case constants.WARNING:
		return gouse.ORANGE_CONSOLE
	case constants.ERROR:
		return gouse.RED_CONSOLE
	case constants.FATAL:
		return gouse.PURPLE_CONSOLE
	case constants.LABEL:
		return gouse.CYAN_CONSOLE
	case constants.DESC:
		return gouse.YELLOW_CONSOLE
	case constants.INPUT:
		return gouse.CYAN_CONSOLE
	case constants.QUERY:
		return gouse.PINK_CONSOLE
	default:
		return gouse.WHITE_CONSOLE
	}
}

func (h *Helper) OutputSuccess(statusMsg string) {
	println(gouse.GREEN_CONSOLE + "\n::: Status: " + statusMsg + gouse.DEFAULT_CONSOLE)
}

func (h *Helper) OutputError(statusMsg string, err string) {
	println(gouse.RED_CONSOLE + "\n::: Status: " + statusMsg + "\n::: Message: " + err + gouse.DEFAULT_CONSOLE)
}

func (h *Helper) OutputDefault(label int8, msg string) {
	println(h.OutputColor(label) + msg + gouse.DEFAULT_CONSOLE)
}

func OutputTime(time string) {
	println(gouse.PURPLE_CONSOLE + "\n::: Query time: " + time + gouse.DEFAULT_CONSOLE)
}