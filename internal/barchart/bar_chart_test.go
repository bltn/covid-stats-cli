package barchart

import (
	"testing"
)

func TestNewBarChartThrowsErrorIfThereAreNoBars(t *testing.T) {
	_, err := NewBarChart("title", make([]Bar, 0))

	if err == nil {
		t.Fatalf("NewBarChart() should throw an error if there are no bars")
	}
}

func TestBarChartPlotsForMultipleBarsWhereScaleFactorIsLessThanOne(t *testing.T) {
	bars := make([]Bar, 0)
	bars = append(bars, NewBar("25th Dec", 10))
	bars = append(bars, NewBar("26th Dec", 25))
	bars = append(bars, NewBar("27th Dec", 50))
	bars = append(bars, NewBar("28th Dec", 30))
	bars = append(bars, NewBar("29th Dec", 1))

	chart, err := NewBarChart("Data about something or other", bars)
	if err != nil {
		t.Fatal(err)
	}

	e := "\n----- Data about something or other -----\n\n" +
		"25th  (10) | *****                      \n" +
		"26th  (25) | ************               \n" +
		"27th  (50) | *************************  \n" +
		"28th  (30) | ***************            \n" +
		"29th  (1)  |                            \n\n"

	plotted := chart.Plot(0.5)

	if plotted != e {
		t.Fatalf("Expected: %s\n\n Got: %s\n\n", e, plotted)
	}
}

func TestBarChartPlotsForMultipleBarsWhereScaleFactorEqualsOne(t *testing.T) {
	bars := make([]Bar, 0)
	bars = append(bars, NewBar("25th Dec", 10))
	bars = append(bars, NewBar("26th Dec", 25))
	bars = append(bars, NewBar("27th Dec", 50))
	bars = append(bars, NewBar("28th Dec", 30))
	bars = append(bars, NewBar("29th Dec", 1))

	chart, err := NewBarChart("Data about something or other", bars)
	if err != nil {
		t.Fatal(err)
	}

	e := "\n----- Data about something or other -----\n\n" +
		"25th  (10) | **********                                          \n" +
		"26th  (25) | *************************                           \n" +
		"27th  (50) | **************************************************  \n" +
		"28th  (30) | ******************************                      \n" +
		"29th  (1)  | *                                                   \n\n"

	plotted := chart.Plot(1.0)

	if plotted != e {
		t.Fatalf("Expected: %s\n\n Got: %s\n\n", e, plotted)
	}
}

func TestBarChartPlotsForOneBar(t *testing.T) {
	bar := NewBar("1st Jan", 100)
	bars := make([]Bar, 0)
	bars = append(bars, bar)

	chart, err := NewBarChart("Data about something or other", bars)
	if err != nil {
		t.Fatal(err)
	}

	expected := "\n----- Data about something or other -----\n\n" +
		"1st J (100) | " +
		"****************************************************************************************************  \n\n"
	plotted := chart.Plot(1.0)

	if plotted != expected {
		t.Fatalf("Expected: %s\n\n Got: %s\n\n", expected, plotted)
	}
}