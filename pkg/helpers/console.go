package helpers

import (
	"os"
	"os/exec"
	"runtime"
	"reflect"
    "github.com/jedib0t/go-pretty/v6/table"
)

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

func TableOutput[H, R, F any](header H, rows []R, footer F) {
	t := table.NewWriter()
    t.SetOutputMirror(os.Stdout)

	if !isNil(header) {
		t.AppendHeader(table.Row{header})
	}

	if rows != nil {
		if len(rows) == 1 {
			for _, row := range rows {
				t.AppendRow([]interface{}{row})
			}
		} else {
			for i, row := range rows {
				t.AppendRow([]interface{}{i, row})
				t.AppendSeparator()
			}
		}
	}

	if !isNil(footer) {
		t.AppendFooter(table.Row{footer})
	}


	t.SetStyle(table.StyleColoredBright)
    t.Render()
}

func isNil(value interface{}) bool {
	if value == nil {
		return true
	}

	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		return v.IsNil()
	}

	return false
}