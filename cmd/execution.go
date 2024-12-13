package cmd

import (
	"fmt"

	gu_console "github.com/thuongtruong109/gouse/console"
	gu_shared "github.com/thuongtruong109/gouse/shared"
	"github.com/thuongtruong109/soundlib/pkg/constants"
)

func (d *Delivery) Execution() {
	loop := func() {
		gu_console.Banner(gu_shared.DOUBLE_ALPHA, "SOUNDLIB")

		option := d.Run()
		d.HandleOption(option)
	}
	loop()

	for {
		fmt.Print(d.helper.OutputColor(constants.QUERY) + constants.AGAIN_TEXT)
		var input string
		fmt.Scanln(&input)
		if input == "y" {
			gu_console.Clear()
			loop()
		} else {
			gu_console.Clear()
			d.helper.OutputNomal(constants.LABEL, constants.THANKYOU_TEXT)
			break
		}
	}
}
