package coviddata

import (
	"covid-stats-cli/internal/rest"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

type response struct {
	Data []struct {
		Date *string
		Cases *int
		Deaths *int
	}
}

type restApi interface {
	getData(previousDays int) ([]data, error)
}

type restApiImpl struct {
	url string
	client rest.Client
}

func NewCovidDataRestApi(url string, client rest.Client) restApi {
	return restApiImpl{url, client}
}

func (api restApiImpl) getData(previousDays int) ([]data, error) {
	resp, err := api.client.Get(api.url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("Received non-200 status code " + strconv.Itoa(resp.StatusCode))
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response response
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return nil, err
	}

	if len(response.Data) == 0 {
		return nil, errors.New(fmt.Sprintf("response %+v is empty", response))
	}

	from := time.Now().Add(time.Duration(-previousDays*24) * time.Hour)
	var covidData []data
	for _, responseData := range response.Data {
		if responseData.Date == nil {
			return nil, errors.New("the covid data api is returning entries with no specified date")
		}

		date, err := time.Parse("2006-01-02", *responseData.Date)
		if err != nil {
			return nil, err
		}

		if isOnOrAfter(from, date) && !isSameDay(time.Now(), date) {
			var cases int
			var deaths int

			if responseData.Cases == nil {
				fmt.Printf("WARNING: the api has no case data for %v. This might skew the results.", date)
				cases = 0
			} else {
				cases = *responseData.Cases
			}

			if responseData.Deaths == nil {
				fmt.Printf("WARNING: the api has no death data for %v. This might skew the results.", date)
				deaths = 0
			} else {
				deaths = *responseData.Deaths
			}

			covidData = append(covidData, data{
				date:   date,
				cases:  cases,
				deaths: deaths,
			})
		}
	}

	return covidData, nil
}

func isSameDay(t1 time.Time, t2 time.Time) bool {
	return t1.Year() == t2.Year() && t1.YearDay() == t2.YearDay()
}

func isOnOrAfter(t1 time.Time, t2 time.Time) bool {
	if t1.Year() < t2.Year() {
		return true
	}

	return t1.Year() == t2.Year() && t1.YearDay() <= t2.YearDay()
}