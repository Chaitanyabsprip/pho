package display

import (
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/montanaflynn/stats"
)

func boxPlot(plotables []Plotable) {
	box := charts.NewBoxPlot()
	box.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "This title",
		Subtitle: "This Subtitle",
	}))
	data := plotables[0]
	y := make([]int, 0)
	for k, v := range data.Data() {
		for i := 0; i < v; i++ {
			y = append(y, k)
		}
	}
	box.SetXAxis([]string{data.Title()}).AddSeries(data.Title(), toBoxData(y))
	f, err := os.Create("box.html")
	if err != nil {
		return
	}
	box.Render(f)
}

func createBoxPlotData(data []int) []float64 {
	dataAsFloat := make([]float64, 0)
	for _, v := range data {
		dataAsFloat = append(dataAsFloat, float64(v))
	}
	min, _ := stats.Min(dataAsFloat)
	max, _ := stats.Max(dataAsFloat)
	q, _ := stats.Quartile(dataAsFloat)
	return []float64{
		min,
		q.Q1,
		q.Q2,
		q.Q3,
		max,
	}
}

func toBoxData(data []int) []opts.BoxPlotData {
	yData := make([]opts.BoxPlotData, 0)
	yData = append(yData, opts.BoxPlotData{Value: createBoxPlotData(data)})
	return yData
}
