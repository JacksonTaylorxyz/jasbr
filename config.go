package main

// This is the main configuration file for jasbr
// You can comment out any modules you don't want to use
// you can also add any new ones here that you make
// Order does matter

// Modules that are displayed in the bar
var BarModules = []Module{
	DateModule{},
	BatteryModule{},
	TimeModule{},
}
