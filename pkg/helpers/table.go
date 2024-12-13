package helpers

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/thuongtruong109/gouse"
)

func TableOutput[H, R any, S, D, T string](header H, rows []R, sum S, description D, time T) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	if !gouse.IsNil(header) {
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

	if !gouse.IsNil(sum) && !gouse.IsNil(description) && !gouse.IsNil(time) {
		t.AppendFooter(table.Row{sum, description, time})
	}

	t.SetStyle(table.StyleColoredBright)
	t.Render()
}

func TableNoOutput[H any, D, T string](header H, description D, time T) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	if !gouse.IsNil(header) {
		t.AppendHeader(table.Row{header})
	}

	if !gouse.IsNil(description) && !gouse.IsNil(time) {
		t.AppendFooter(table.Row{description, time})
	}

	t.SetStyle(table.StyleColoredBright)
	t.Render()
}
