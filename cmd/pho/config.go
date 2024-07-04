package main

import "slices"

// Config struct  
type Config struct {
	DisplayMethod string
	Filepaths     []string
	Recurse       bool
}

// Equals method  
func (c Config) Equals(other Config) bool {
	return c.DisplayMethod == other.DisplayMethod &&
		slices.Equal(c.Filepaths, other.Filepaths) &&
		c.Recurse == other.Recurse
}
