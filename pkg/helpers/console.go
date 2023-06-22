package helpers

import (
	"os"
	"os/exec"
	"runtime"
	"music-management/pkg/constants"
)

func (h *Helper) OutputColor(level int8) string {
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

func (h *Helper) Output(level int8, msg string) {
	println(h.OutputColor(level) + msg + constants.Reset)
}

func (h *Helper) ClearConsole() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}