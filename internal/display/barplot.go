package display

import (
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func barPlot(plotables []Plotable) {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "This title",
		Subtitle: "This Subtitle",
	}))
	data := plotables[0]
	x, y := sanitize(data.Data())
	bar.SetXAxis(x).AddSeries(data.Title(), toBarData(y))
	f, _ := os.Create("bar.html")
	bar.Render(f)
}

func toBarData(data []int) []opts.BarData {
	yData := make([]opts.BarData, 0)
	for _, value := range data {
		yData = append(yData, opts.BarData{Value: value})
	}
	return yData
}
