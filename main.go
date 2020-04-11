package main

/*
	Author: Jackson Taylor
	Date: 3/28/2020
	Description: This is Jackson's Awesome Status Bar for suckless', very nice, dwm.
*/

import (
	"fmt"
	"os/exec"
	"time"
)

// Module is the interface for you to extend what your bar does.
type Module interface {
	GetInfo() (string, error)
}

// Modules that are displayed in the bar
var BarModules = []Module{
	DateModule{},
	BatteryModule{},
}

func main() {
	main := Bar{
		Modules:     BarModules,
		RefreshRate: time.Second * 1,
	}
	fmt.Println("Bar Started")

	for {
		display(main.Display())
		time.Sleep(main.RefreshRate)
		fmt.Println("Updating bar")
	}
}

func display(value string) (err error) {
	err = exec.Command("xsetroot", "-name", value).Run()
	return err
}
