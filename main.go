package main

/*
	Author: Jackson Taylor
	Date: 3/28/2020
	Description: This is Jackson's Awesome Status Bar for suckless', very nice, dwm.
*/

import (
	"fmt"
	// "os"
	"os/exec"
	"time"
)

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

	var value string

	for {
		value = ""
		for _, b := range main.Modules {
			p1, err := b.GetInfo()
			if err != nil {
				fmt.Println(err)
				continue
			}
			value += "[" + p1 + "]"
		}

		setBar(value)
	}
}

func getBattery() string {
	return "100 %"
}

func setBar(value string) (err error) {
	err = exec.Command("xsetroot", "-name", value).Run()
	return err
}
