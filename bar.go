package main

import (
	"fmt"
	"github.com/HelixBePraised/jasbr/modules"
)

// Bar is the struct that holds each of the modules and displays the data from them
type Bar struct {
	Modules []modules.Module
}

func (b Bar) Display() string {
	var val string

	for _, mod := range b.Modules {
		s, err := mod.GetInfo()

		// TODO: Handle error differently
		if err != nil {
			fmt.Println(err)
		}
		val += "[" + s + "]"
	}

	return val
}
