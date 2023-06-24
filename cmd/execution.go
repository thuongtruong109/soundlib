package cmd

import (
	"fmt"
	"music-management/pkg/constants"
)

func (d *Delivery) Execution() {
	loop := func() {
		d.helper.OutputNomal(constants.LABEL, "\n===== MUSIC MANAGEMENT =====")
		d.DisplayOptions()
	
		fmt.Print(d.helper.OutputColor(constants.INPUT) + "\n» Choose action: ")
	
		var option int
		fmt.Scanln(&option)
	
		d.HandleOption(option)
	}
	loop()

	for {
		fmt.Print(d.helper.OutputColor(constants.QUERY) + "\n ¤ Do you want to continue? (y/n): ")
		var input string
		fmt.Scanln(&input)
		if input == "y" {
			d.helper.ClearConsole()
			loop()
		} else {
			d.helper.ClearConsole()
			d.helper.OutputNomal(constants.LABEL, "\n --- ♥ Thank you for using our service!")
			break
		}
	}
}