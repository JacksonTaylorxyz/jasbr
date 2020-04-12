package main

// Module is the interface for you to extend what your bar does.
type Module interface {
	GetInfo() (string, error)
}
