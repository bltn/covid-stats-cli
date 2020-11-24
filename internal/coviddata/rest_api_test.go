package coviddata

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
	"time"
)

type mockRestClient struct {
	response http.Response
}

func (c mockRestClient) Get(_ string) (resp *http.Response, err error) {
	return &c.response, nil
}

func TestRestApi_GetData_ClientReturnsNon200StatusCode(t *testing.T) {
	url := "http://www.amireallyreal.com"
	client := mockRestClient{
		http.Response{
			StatusCode: 500,
			Body: ioutil.NopCloser(bytes.NewBufferString("Hello World")),
		},
	}
	api := NewCovidDataRestApi(url, client)

	data, err := api.getData(5)

	expectedErrorMsg := "Received non-200 status code 500"
	if err == nil || (expectedErrorMsg != err.Error() || len(data) > 0) {
		t.Fatalf("Expected err '%v' to be '%v' and data (len=%d) to be empty\n",
			err,
			expectedErrorMsg,
			len(data))
	}
}

func TestRestApi_GetData_ResponseHasDataWithMissingDate(t *testing.T) {
	url := "http://www.amireallyreal.com"
	client := mockRestClient{
		http.Response{
			StatusCode: 200,
			Body: ioutil.NopCloser(bytes.NewBufferString("{\"data\":[{\"deaths\": 5, \"cases\": 500}]}")),
		},
	}
	api := NewCovidDataRestApi(url, client)

	data, err := api.getData(5)

	expectedErrorMsg := "the covid data api is returning entries with no specified date"
	if err == nil || (expectedErrorMsg != err.Error() || len(data) > 0) {
		t.Fatalf("Expected err '%v' to be '%v' and data (len=%d) to be empty\n",
			err,
			expectedErrorMsg,
			len(data))
	}
}

func TestRestApi_GetData_ResponseHasDataWithMissingCases(t *testing.T) {
	yesterday := time.Now().Add(time.Hour * -24).Format("2006-01-02")
	url := "http://www.amireallyreal.com"
	client := mockRestClient{
		http.Response{
			StatusCode: 200,
			Body: ioutil.NopCloser(bytes.NewBufferString("{\"data\":[{\"date\":\""+yesterday+"\",\"deaths\":81}]}")),
		},
	}
	api := NewCovidDataRestApi(url, client)

	data, err := api.getData(5)

	if err != nil || data[0].cases != 0 {
		t.Fatalf("Expected err %v to be nil and cases %v to be 0", err, data[0].cases)
	}
}

func TestRestApi_GetData_ResponseHasDataWithMissingDeaths(t *testing.T) {
	yesterday := time.Now().Add(time.Hour * -24).Format("2006-01-02")
	url := "http://www.amireallyreal.com"
	client := mockRestClient{
		http.Response{
			StatusCode: 200,
			Body: ioutil.NopCloser(bytes.NewBufferString("{\"data\":[{\"date\":\""+yesterday+"\",\"cases\":81}]}")),
		},
	}
	api := NewCovidDataRestApi(url, client)

	data, err := api.getData(5)

	if err != nil || data[0].deaths != 0 {
		t.Fatalf("Expected err %v to be nil and deaths %v to be 0", err, data[0].deaths)
	}
}

func TestRestApi_GetData_ResponseHasNoData(t *testing.T) {
	url := "http://www.amireallyreal.com"
	client := mockRestClient{
		http.Response{
			StatusCode: 200,
			Body: ioutil.NopCloser(bytes.NewBufferString("{\"data\":[]}")),
		},
	}
	api := NewCovidDataRestApi(url, client)

	data, err := api.getData(5)

	expectedErrorMsg := "response {Data:[]} is empty"
	if err == nil || (expectedErrorMsg != err.Error() || len(data) > 0) {
		t.Fatalf("Expected err '%v' to be '%v' and data (len=%d) to be empty\n",
			err,
			expectedErrorMsg,
			len(data))
	}
}

func TestRestApi_GetData_EmptyJsonResponse(t *testing.T) {
	url := "http://www.amireallyreal.com"
	client := mockRestClient{
		http.Response{
			StatusCode: 200,
			Body: ioutil.NopCloser(bytes.NewBufferString("{}")),
		},
	}
	api := NewCovidDataRestApi(url, client)

	data, err := api.getData(5)

	expectedErrorMsg := "response {Data:[]} is empty"
	if err == nil || (expectedErrorMsg != err.Error() || len(data) > 0) {
		t.Fatalf("Expected err '%v' to be '%v' and data (len=%d) to be empty\n",
			err,
			expectedErrorMsg,
			len(data))
	}
}

func TestRestApi_GetData_MalformedJsonResponse(t *testing.T) {
	url := "http://www.amireallyreal.com"
	client := mockRestClient{
		http.Response{
			StatusCode: 200,
			Body: ioutil.NopCloser(bytes.NewBufferString("[]")),// starting a json doc with '[]' is invalid syntax
		},
	}
	api := NewCovidDataRestApi(url, client)

	data, err := api.getData(5)

	expectedErrorMsg := "json: cannot unmarshal array into Go value of type coviddata.response"
	if err == nil || (expectedErrorMsg != err.Error() || len(data) > 0) {
		t.Fatalf("Expected err '%v' to be '%v' and data (len=%d) to be empty\n",
			err,
			expectedErrorMsg,
			len(data))
	}
}

func TestRestApi_GetData_FiltersOutDataWithTodaysDate(t *testing.T) {
	url := "http://www.amireallyreal.com"

	today := time.Now()
	yesterday := today.Add(time.Hour * -24)

	response := "{\"data\":["
	response += "{\"deaths\":10,\"cases\":81,\"date\":\"" + yesterday.Format("2006-01-02") + "\"},"
	response += "{\"deaths\":5,\"cases\":40,\"date\":\"" + today.Format("2006-01-02") + "\"}"
	response += "]}"
	client := mockRestClient{
		http.Response{
			StatusCode: 200,
			Body: ioutil.NopCloser(bytes.NewBufferString(response)),
		},
	}
	api := NewCovidDataRestApi(url, client)

	data, _ := api.getData(1)

	if len(data) != 1 {
		t.Fatalf("Expected data %+v to have a length of 1 (today's should be filtered out)", data)
	}

	if !isSameDay(data[0].date, yesterday) || data[0].cases != 81 || data[0].deaths != 10 {
		t.Fatalf("data %+v does not have the expected fields", data)
	}
}

func TestRestApi_GetData(t *testing.T) {
	url := "http://www.amireallyreal.com"

	today := time.Now()

	fourDaysAgo := data{date: today.Add(time.Hour * -96), cases: 1000, deaths: 50}
	threeDaysAgo := data{date: today.Add(time.Hour * -72), cases: 400, deaths: 4}
	twoDaysAgo := data{date: today.Add(time.Hour * -48), cases: 400, deaths: 4}
	oneDayAgo := data{date: today.Add(time.Hour * -24), cases: 81, deaths: 10}
	dataForToday := data{date: today, cases: 555, deaths: 8}

	response := "{\"data\":["
	response += asJson(fourDaysAgo) + ","
	response += asJson(threeDaysAgo) + ","
	response += asJson(twoDaysAgo) + ","
	response += asJson(oneDayAgo) + ","
	response += asJson(dataForToday)
	response += "]}"
	client := mockRestClient{
		http.Response{
			StatusCode: 200,
			Body: ioutil.NopCloser(bytes.NewBufferString(response)),
		},
	}
	api := NewCovidDataRestApi(url, client)

	actual, _ := api.getData(3)

	expected := make([]data, 0)
	expected = append(expected, oneDayAgo)
	expected = append(expected, twoDaysAgo)
	expected = append(expected, threeDaysAgo)

	if !deepEqual(actual, expected) {
		t.Fatalf("Expected data %+v to match %+v", actual, expected)
	}
}

func deepEqual(a []data, b []data) bool {
	if len(a) != len(b) {
		return false
	}

	for _, data := range a {
		if !contains(b, data) {
			return false
		}
	}

	return true
}

func contains(dd []data, d data) bool {
	for _, data := range dd {
		if isSameDay(data.date, d.date) && data.deaths == d.deaths && data.cases == d.cases {
			return true
		}
	}

	return false
}

func asJson(data data) string {
	asJson := "{"
	asJson += "\"deaths\":" + strconv.Itoa(data.deaths) + ","
	asJson += "\"cases\":" + strconv.Itoa(data.cases) + ","
	asJson += "\"date\":\"" + data.date.Format("2006-01-02") + "\""
	asJson += "}"

	return asJson
}
