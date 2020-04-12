package main

/*
	Author: Jackson Taylor
	Date: 3/28/2020
	Description: This is Jackson's Awesome Status Bar for suckless', very nice, dwm.
*/

import (
	"fmt"
)

// Module is the interface for you to extend what your bar does.
type Module interface {
	GetInfo() (string, error)
}

// Modules that are displayed in the bar
var BarModules = []Module{
	DateModule{},
	BatteryModule{},
	TimeModule{},
}

func main() {
	main := Bar{
		Modules: BarModules,
	}

	// Display the bar info to stdout
	fmt.Println(main.Display())
}
