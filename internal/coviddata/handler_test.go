package coviddata

import (
	"errors"
	"testing"
	"time"
)

func TestHandler_GetDeathsChart_ApiReturnsError(t *testing.T) {
	apiErr := errors.New("our data centre went bye bye")
	mockApi := givenApiThatReturns(nil, apiErr)
	handler := NewHandler(mockApi)

	chart, err := handler.GetDeathsChart(5)

	if chart != "" || err != apiErr {
		t.Fatalf("Expected error %v to equal %v and chart '%s' to be empty",
			err, apiErr, chart)
	}
}

func TestHandler_GetDeathsChart_NewBarChartReturnsError(t *testing.T) {
	emptyData := make([]data, 0) // barchart.NewBarChart will return an error when no data given
	mockApi := givenApiThatReturns(emptyData, nil)
	handler := NewHandler(mockApi)

	chart, err := handler.GetDeathsChart(5)

	expectedErrorMsgFromBarChart := "there are no bars in the bar chart"
	if chart != "" || (err == nil || err.Error() != expectedErrorMsgFromBarChart) {
		t.Fatalf("got %s and %v", chart, err)
	}
}

func TestHandler_GetDeathsChart_PlotsChartWithDatesSortedOldToNew(t *testing.T) {
	today := time.Now()
	oneDayAgo := today.Add(time.Hour * -24)
	twoDaysAgo := today.Add(time.Hour * -48)
	threeDaysAgo := today.Add(time.Hour * -72)
	fourDaysAgo := today.Add(time.Hour * -96)
	fiveDaysAgo := today.Add(time.Hour * -120)

	deathsData := make([]data, 0)
	deathsData = append(deathsData, data{
		date:  twoDaysAgo,
		deaths: 6,
	})
	deathsData = append(deathsData, data{
		date:  oneDayAgo,
		deaths: 5,
	})
	deathsData = append(deathsData, data{
		date:  threeDaysAgo,
		deaths: 7,
	})
	deathsData = append(deathsData, data{
		date:  fourDaysAgo,
		deaths: 12,
	})
	deathsData = append(deathsData, data{
		date:  fiveDaysAgo,
		deaths: 10,
	})

	mockApi := givenApiThatReturns(deathsData, nil)
	handler := NewHandler(mockApi)

	expectedChart := "\n----- New deaths -----\n\n" +
	fiveDaysAgo.Format("02/01") + " (10) | **********    \n" +
	fourDaysAgo.Format("02/01") + " (12) | ************  \n" +
	threeDaysAgo.Format("02/01") + " (7)  | *******       \n" +
	twoDaysAgo.Format("02/01") + " (6)  | ******        \n" +
	oneDayAgo.Format("02/01") + " (5)  | *****         \n\n"
	chart, err := handler.GetDeathsChart(5)

	if chart != expectedChart || err != nil {
		t.Fatalf("expected chart '%s' and nil err, but got chart '%s' and err '%v'", expectedChart, chart, err)
	}
}

func TestHandler_GetCasesChart_ApiReturnsError(t *testing.T) {
	apiErr := errors.New("our data centre went bye bye")
	mockApi := givenApiThatReturns(nil, apiErr)
	handler := NewHandler(mockApi)

	chart, err := handler.GetCasesChart(5)

	if chart != "" || err != apiErr {
		t.Fatalf("Expected error %v to equal %v and chart '%s' to be empty",
			err, apiErr, chart)
	}
}

func TestHandler_GetCasesChart_NewBarChartReturnsError(t *testing.T) {
	emptyData := make([]data, 0) // barchart.NewBarChart will return an error when no data given
	mockApi := givenApiThatReturns(emptyData, nil)
	handler := NewHandler(mockApi)

	chart, err := handler.GetCasesChart(5)

	expectedErrorMsgFromBarChart := "there are no bars in the bar chart"
	if chart != "" || (err == nil || err.Error() != expectedErrorMsgFromBarChart) {
		t.Fatalf("got %s and %v", chart, err)
	}
}

func TestHandler_GetCasesChart_PlotsChartWithDatesSortedOldToNew(t *testing.T) {
	today := time.Now()
	oneDayAgo := today.Add(time.Hour * -24)
	twoDaysAgo := today.Add(time.Hour * -48)
	threeDaysAgo := today.Add(time.Hour * -72)
	fourDaysAgo := today.Add(time.Hour * -96)
	fiveDaysAgo := today.Add(time.Hour * -120)

	caseData := make([]data, 0)
	caseData = append(caseData, data{
		date:  twoDaysAgo,
		cases: 6,
	})
	caseData = append(caseData, data{
		date:  oneDayAgo,
		cases: 5,
	})
	caseData = append(caseData, data{
		date:  threeDaysAgo,
		cases: 7,
	})
	caseData = append(caseData, data{
		date:  fourDaysAgo,
		cases: 12,
	})
	caseData = append(caseData, data{
		date:  fiveDaysAgo,
		cases: 10,
	})

	mockApi := givenApiThatReturns(caseData, nil)
	handler := NewHandler(mockApi)

	expectedChart := "\n----- New cases -----\n\n" +
	fiveDaysAgo.Format("02/01") + " (10) | **********    \n" +
	fourDaysAgo.Format("02/01") + " (12) | ************  \n" +
	threeDaysAgo.Format("02/01") + " (7)  | *******       \n" +
	twoDaysAgo.Format("02/01") + " (6)  | ******        \n" +
	oneDayAgo.Format("02/01") + " (5)  | *****         \n\n"
	chart, err := handler.GetCasesChart(5)

	if chart != expectedChart || err != nil {
		t.Fatalf("expected chart '%s' and nil err, but got chart '%s' and err '%v'", expectedChart, chart, err)
	}
}

type mockRestApi struct {
	mockGetData func(previousDays int) ([]data, error)
}

func (m mockRestApi) getData(previousDays int) ([]data, error) {
	return m.mockGetData(previousDays)
}

func givenApiThatReturns(d []data, e error) mockRestApi {
	mf := func(i int) ([]data, error) {
		return d, e
	}
	return mockRestApi{mf}
}
