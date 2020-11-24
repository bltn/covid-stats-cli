package coviddata

import (
	"covid-stats-cli/internal/barchart"
	"sort"
)

type Handler struct {
	api restApi
}

func NewHandler(api restApi) *Handler {
	return &Handler{api}
}

func (h Handler) GetCasesChart(previousWeeks int) (string, error) {
	covidData, err := h.api.getData(previousWeeks * 7)
	if err != nil {
		return "", err
	}

	// sort oldest -> newest
	sort.Slice(covidData, func(i, j int) bool {
		return covidData[i].date.Before(covidData[j].date)
	})

	var bars []barchart.Bar
	for _, d := range covidData {
		bars = append(bars, barchart.NewBar(d.date.Format("02/01"), d.cases))
	}

	chart, err := barchart.NewBarChart("New cases", bars)
	if err != nil {
		return "", err
	}

	scaleFactor := barchart.CalculateScaleFactor(bars, 100.0)

	return chart.Plot(scaleFactor), nil
}

func (h Handler) GetDeathsChart(previousWeeks int) (string, error) {
	covidData, err := h.api.getData(previousWeeks * 7)
	if err != nil {
		return "", err
	}

	// sort oldest -> newest
	sort.Slice(covidData, func(i, j int) bool {
		return covidData[i].date.Before(covidData[j].date)
	})

	var bars []barchart.Bar
	for _, d := range covidData {
		bars = append(bars, barchart.NewBar(d.date.Format("02/01"), d.deaths))
	}

	chart, err := barchart.NewBarChart("New deaths", bars)
	if err != nil {
		return "", err
	}

	scaleFactor := barchart.CalculateScaleFactor(bars, 100.0)

	return chart.Plot(scaleFactor), nil
}