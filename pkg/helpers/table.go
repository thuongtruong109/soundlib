package helpers

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/thuongtruong109/gouse/types"
)

func TableOutput[H, R, F, F2, F3 any](header H, rows []R, footer ...F) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	if !types.IsNil(header) {
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

	if !types.IsNil(footer) {
		t.AppendFooter(table.Row{footer[0], footer[1], footer[2]})
	}

	t.SetStyle(table.StyleColoredBright)
	t.Render()
}

func TableNoOutput[H, F, F2, F3 any](header H, footer ...F) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	if !types.IsNil(header) {
		t.AppendHeader(table.Row{header})
	}

	if !types.IsNil(footer) {
		t.AppendFooter(table.Row{footer[0], footer[1], footer[2]})
	}

	t.SetStyle(table.StyleColoredBright)
	t.Render()
}
