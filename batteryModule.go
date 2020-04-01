package main

import (
	"io/ioutil"
)

// TODO: Move this to a central spot, like a config.h
const (
	batteryPath = "/sys/class/power_supply/BAT0/capacity"
	identifier  = "BAT"
)

// BatteryModule is the module that controls finding out the battery
// percentage, mainly useful in laptops
type BatteryModule struct {
}

func (_ BatteryModule) GetInfo() (string, error) {
	return readBatteryLevel(batteryPath)
}

func readBatteryLevel(path string) (string, error) {
	dat, err := ioutil.ReadFile(path)

	if err != nil {
		return "NULL", err
	}

	return (string(dat[:len(dat)-1]) + "%"), nil
}
