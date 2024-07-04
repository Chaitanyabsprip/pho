package display

import (
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func linePlot(plotables []Plotable) {
	bar := charts.NewLine()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "This title",
		Subtitle: "This Subtitle",
	}))
	data := plotables[0]
	x, y := sanitize(data.Data())
	bar.SetXAxis(x).AddSeries(data.Title(), toLineData(y))
	f, _ := os.Create("line.html")
	bar.Render(f)
}

func toLineData(data []int) []opts.LineData {
	yData := make([]opts.LineData, 0)
	for _, value := range data {
		yData = append(yData, opts.LineData{Value: value})
	}
	return yData
}
