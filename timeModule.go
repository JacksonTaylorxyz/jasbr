package main

import "time"

type TimeModule struct {
}

func (_ TimeModule) GetInfo() (string, error) {
	return time.Now().Format(time.Kitchen), nil
}
