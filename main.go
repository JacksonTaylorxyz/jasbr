package main

/*
	Author: Jackson Taylor
	Date: 3/28/2020
	Description: This is Jackson's Awesome Status Bar for suckless', very nice, dwm.
*/

import (
	"fmt"
)

func main() {
	main := Bar{
		Modules: BarModules,
	}

	// Display the bar info to stdout
	fmt.Println(main.Display())
}
