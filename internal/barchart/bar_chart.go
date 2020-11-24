package barchart

import (
	"errors"
	"strconv"
	"strings"
)

type BarChart struct {
	title string
	bars []Bar
}

func NewBarChart(title string, bars []Bar) (BarChart, error) {
	if len(bars) < 1 {
		return BarChart{}, errors.New("there are no bars in the bar chart")
	}

	return BarChart{title: title, bars: bars}, nil
}


func (b BarChart) Plot(scaleFactor float64) string {
	var plotted string

	highestCount := 0
	for _, bar := range b.bars {
		if bar.count > highestCount {
			highestCount = bar.count
		}
	}

	xAxis := int(float64(highestCount) * scaleFactor) + 1

	plotted += "\n"
	plotted += "----- " + b.title + " -----\n"
	plotted += "\n"

	for _, bar := range b.bars {
		scaledCount := int(float64(bar.count) * scaleFactor)

		padding := len(strconv.Itoa(highestCount)) - len(strconv.Itoa(bar.count))
		yAxisLabel := bar.label + " (" + strconv.Itoa(bar.count) + ") " + strings.Repeat(" ", padding) + "| "
		plotted += yAxisLabel

		for x := 0; x < scaledCount; x++ {
			plotted += "*"
		}
		for x := scaledCount; x <= xAxis; x++ {
			plotted += " "
		}
		plotted += "\n"
	}

	plotted += "\n"

	return plotted
}

