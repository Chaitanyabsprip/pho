// Package display provides display  
package display

import (
	"sort"
)

// Plotable interface  
type Plotable interface {
	Data() map[int]int
	Title() string
}

// Display function  
func Display(method string, plotables []Plotable) {
	switch method {
	case "bar":
		barPlot(plotables)
	case "box":
		boxPlot(plotables)
	case "line":
		linePlot(plotables)
	}
}

func sanitize(data map[int]int) ([]int, []int) {
	keys := make([]int, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	x := make([]int, len(keys))
	y := make([]int, len(keys))
	for i, key := range keys {
		x[i] = key
		y[i] = data[key]
	}

	return x, y
}
