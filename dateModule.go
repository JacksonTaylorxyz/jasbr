package main

import (
	"time"
)

type DateModule struct {
}

// getDate formats a string for the status bar
func (_ DateModule) GetInfo() (string, error) {
	now := time.Now()
	// return string([]rune(now.Weekday().String()[0:3]))
	// return now.Format(time.Stamp)
	return string([]rune(now.Month().String())) + " " + now.Format(time.Kitchen), nil
}
