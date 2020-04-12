package main

import (
	"fmt"
	"strconv"
	"time"
)

type DateModule struct {
}

// getDate formats a string for the status bar
func (_ DateModule) GetInfo() (string, error) {
	now := time.Now()
	// return now.Weekday().String() + " " + now.Month().String()[0:3] + " " + now.Year(), nil
	return fmt.Sprintf("%s %s %s %s", now.Weekday(), now.Month().String()[0:3], strconv.Itoa(now.Day()), strconv.Itoa(now.Year())), nil
}
