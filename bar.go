package main

import "time"

// Bar is the struct that holds each of the modules and displays the data from them
type Bar struct {
	Modules     []Module
	RefreshRate time.Duration
}

func (b Bar) Display() string {
	var val string

	for _, mod := range b.Modules {
		s, _ := mod.GetInfo()
		val += s
		time.Sleep(b.RefreshRate)
	}

	return val
}
