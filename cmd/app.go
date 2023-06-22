package cmd

import (
	"fmt"

	"music-management/pkg/helpers"
	"music-management/pkg/constants"
)

func (d *Delivery) HandlerExecution() {
	loop := func() {
		helpers.Output(constants.LABEL, "\n===== MUSIC MANAGEMENT =====")
		d.DisplayOptions()
	
		fmt.Print(helpers.OutputColor(constants.INPUT) + "Choose action: ")
	
		var option int
		fmt.Scanln(&option)
	
		d.HandleOption(option)
	}
	loop()

	for {
		fmt.Print(helpers.OutputColor(constants.QUERY) + "Do you want to continue? (y/n): ")
		var input string
		fmt.Scanln(&input)
		if input == "y" {
			loop()
		} else {
			break
		}
	}
}