package helpers

import (
	"os"
	"os/exec"
	"runtime"
    "github.com/jedib0t/go-pretty/v6/table"
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

func (h *Helper) TableOutput(header interface{}, rows []interface{}, footer interface{}) {
	t := table.NewWriter()
    t.SetOutputMirror(os.Stdout)

	if header != nil {
		t.AppendHeader(table.Row{header})
	}

	if rows != nil {
		for i, row := range rows {
			t.AppendRow([]interface{}{i, row})
			t.AppendSeparator()
		}
	}

	if footer != nil {
		t.AppendFooter(table.Row{footer})
	}

	t.SetStyle(table.StyleColoredBright)
    t.Render()
}