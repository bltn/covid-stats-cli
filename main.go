package main

import (
	"bufio"
	"covid-stats-cli/internal/coviddata"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	api := coviddata.NewCovidDataRestApi(covidApiUrl(), http.DefaultClient)
	covidDataHandler := coviddata.NewHandler(api)

	printIntroTitle()

	userInput := make(chan string)
	go listenForInput(userInput)

	for {
		fmt.Println("Type:")
		fmt.Println()
		fmt.Println("- d for deaths")
		fmt.Println("- c for cases")
		fmt.Println()
		fmt.Print("> ")

		select {
		case input := <-userInput:
			if input == "d" {
				printDeathsMenu()
				fmt.Print("> ")
				select {
				case furtherInput := <-userInput:
					fmt.Println()
					printDeathStats(furtherInput, covidDataHandler)
				}
			} else if input == "c" {
				printCasesMenu()
				fmt.Print("> ")
				select {
				case furtherInput := <-userInput:
					fmt.Println()
					printCaseStats(furtherInput, covidDataHandler)
				}
			} else {
				fmt.Printf("'%s' isn't really something I offered, is it? :) \n\n", input)
			}
		}
	}
}

func printIntroTitle() {
	fmt.Println()
	fmt.Println("Ready for some anxiety? Awesome! Anxiety for everyone!❤️")
	fmt.Println()
	fmt.Println("Here you can check all the latest COVID trends....")
	fmt.Println(".... never again question whether it's time to start hating on the Tories and your fellow man")
	fmt.Println()
	fmt.Println("**NOTE** this tool is for viewing COVID trends, not the raw figures. You'll find no raw data" +
		" in these parts.")
	fmt.Println()
	fmt.Println("<------- Enjoy :) :) -------> ")
	fmt.Println()
}

func listenForInput(input chan<- string) {
	for {
		buf := bufio.NewReader(os.Stdin)
		bytes, err := buf.ReadBytes('\n')
		if err != nil {
			panic(err)
		}

		input <- strings.TrimSpace(string(bytes))
	}
}

func printCaseStats(input string, handler *coviddata.Handler) {
	switch input {
	case "w":
		fmt.Println("Fetching cases for the last 1 weeks")
		stats, err := handler.GetCasesChart(1)
		if err != nil {
			fmt.Printf("Error fetching the case stats: %+v\n", err)
		} else {
			fmt.Println(stats)
		}
	case "ww":
		fmt.Println("Fetching cases for the last 2 weeks")
		stats, err := handler.GetCasesChart(2)
		if err != nil {
			fmt.Printf("Error fetching the case stats: %+v\n", err)
		} else {
			fmt.Println(stats)
		}
	case "www":
		fmt.Println("Fetching cases for the last 3 weeks")
		stats, err := handler.GetCasesChart(3)
		if err != nil {
			fmt.Printf("Error fetching the case stats: %+v\n", err)
		} else {
			fmt.Println(stats)
		}
	case "m":
		fmt.Println("Fetching cases for the last 4 weeks")
		stats, err := handler.GetCasesChart(4)
		if err != nil {
			fmt.Printf("Error fetching the case stats: %+v\n", err)
		} else {
			fmt.Println(stats)
		}
	case "mm":
		fmt.Println("Fetching cases for the last 8 weeks")
		stats, err := handler.GetCasesChart(8)
		if err != nil {
			fmt.Printf("Error fetching the case stats: %+v\n", err)
		} else {
			fmt.Println(stats)
		}
	case "mmm":
		fmt.Println("Fetching cases for the last 12 weeks")
		stats, err := handler.GetCasesChart(12)
		if err != nil {
			fmt.Printf("Error fetching the case stats: %+v\n", err)
		} else {
			fmt.Println(stats)
		}
	default:
		fmt.Printf("'%s' is not a valid option mmmm'kay.....\n", input)
	}
}

func printDeathStats(input string, handler *coviddata.Handler) {
	switch input {
	case "w":
		fmt.Println("Fetching deaths stats for the last 1 weeks...")
		stats, err := handler.GetDeathsChart(1)
		if err != nil {
			fmt.Printf("Error fetching the death stats: %+v\n", err)
		} else {
			fmt.Println(stats)
		}
	case "ww":
		fmt.Println("Fetching deaths stats for the last 2 weeks...")
		stats, err := handler.GetDeathsChart(2)
		if err != nil {
			fmt.Printf("Error fetching the death stats: %+v\n", err)
		} else {
			fmt.Println(stats)
		}
	case "www":
		fmt.Println("Fetching deaths stats for the last 3 weeks...")
		stats, err := handler.GetDeathsChart(3)
		if err != nil {
			fmt.Printf("Error fetching the death stats: %+v\n", err)
		} else {
			fmt.Println(stats)
		}
	case "m":
		fmt.Println("Fetching deaths stats for the last 4 weeks...")
		stats, err := handler.GetDeathsChart(4)
		if err != nil {
			fmt.Printf("Error fetching the death stats: %+v\n", err)
		} else {
			fmt.Println(stats)
		}
	case "mm":
		fmt.Println("Fetching deaths stats for the last 8 weeks...")
		stats, err := handler.GetDeathsChart(8)
		if err != nil {
			fmt.Printf("Error fetching the death stats: %+v\n", err)
		} else {
			fmt.Println(stats)
		}
	case "mmm":
		fmt.Println("Fetching deaths stats for the last 12 weeks...")
		stats, err := handler.GetDeathsChart(12)
		if err != nil {
			fmt.Printf("Error fetching the death stats: %+v\n", err)
		} else {
			fmt.Println(stats)
		}
	default:
		fmt.Printf("'%s' is not a valid option mmmm'kay.....\n", input)
	}
}

func printCasesMenu() {
	fmt.Println()
	fmt.Println("- w for the last weeks' cases")
	fmt.Println("- ww for the last twos' cases")
	fmt.Println("- www for the last three weeks' cases")
	fmt.Println("- m for the last four weeks' cases")
	fmt.Println("- mm for the last eight weeks' cases")
	fmt.Println("- mmm for the last twelve weeks' cases")
	fmt.Println()
}

func printDeathsMenu() {
	fmt.Println()
	fmt.Println("- w for the last weeks' deaths")
	fmt.Println("- ww for the last twos' deaths")
	fmt.Println("- www for the last three weeks' deaths")
	fmt.Println("- m for the last four weeks' deaths")
	fmt.Println("- mm for the last eight weeks' deaths")
	fmt.Println("- mmm for the last twelve weeks' deaths")
	fmt.Println()
}

func covidApiUrl() string {
	return "https://api.coronavirus.data.gov.uk/v1/data?filters=areaType=nation;areaName=england" +
		"&structure={\"date\":\"date\",\"cases\":\"newCasesByPublishDate\",\"deaths\":\"newDeaths28DaysByPublishDate\"}"
}
