package cmd

import (
	"fmt"

	"github.com/thuongtruong109/gouse"
	"github.com/thuongtruong109/soundlib/pkg/constants"
)

func (d *Delivery) Execution() {
	loop := func() {
		gouse.Banner(gouse.DOUBLE_ALPHA, constants.APP_NAME)

		option := d.Run()
		d.HandleOption(option)
	}
	loop()

	for {
		fmt.Print(d.helper.OutputColor(constants.QUERY) + constants.AGAIN_TEXT)
		var input string
		fmt.Scanln(&input)
		if input == "y" {
			gouse.Cls()
			loop()
		} else {
			gouse.Cls()
			d.helper.OutputDefault(constants.LABEL, constants.THANKYOU_TEXT)
			break
		}
	}
}
