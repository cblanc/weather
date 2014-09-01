package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type dailyForecast struct {
	Dt   int64
	Temp struct {
		Min float32
		Max float32
		Day float32
	}
	Weather  []weatherDescription
	Pressure float32
	Humidity int16
	Clouds   int16
	Speed    float32
}

type weatherDescription struct {
	Main        string
	Description string
}

type forecast struct {
	City struct {
		Name  string
		Coord struct {
			Lon float32
			Lat float32
		}
	}
	List []dailyForecast
}

func main() {
	var location string

	if len(os.Args) < 2 {
		fmt.Printf("Please provide a location for a weather forecast. E.g. $ weather london")
		os.Exit(1)
	} else {
		location = os.Args[1]
		fmt.Printf("Looking up: %s\n", location)
	}

	response, err := http.Get("http://api.openweathermap.org/data/2.5/forecast/daily?mode=json&units=metric&cnt=5&q=" + location)

	defer response.Body.Close()

	if err != nil {
		fmt.Printf("Unable to retrieve data due to error: %s", err)
		response.Body.Close()
		os.Exit(1)
	}

	var f forecast

	if err := json.NewDecoder(response.Body).Decode(&f); err != nil {
		fmt.Printf("Unable to parse result due to error: %s", err)
		os.Exit(1)
	}

	fmt.Printf("5 day forecast for %s (%.5f, %.5f)\n\n", f.City.Name, f.City.Coord.Lon, f.City.Coord.Lat)

	for _, daily := range f.List {
		fmt.Printf("Date: %s\n", time.Unix(daily.Dt, 0).Format(time.RFC850))
		fmt.Printf("Description: %s (%s)\n", daily.Weather[0].Main, daily.Weather[0].Description)
		fmt.Printf("Temperate (Range): %0.fC (%.0fC-%.0fC)\n", daily.Temp.Day, daily.Temp.Min, daily.Temp.Max)
		fmt.Printf("Cloud Coverage: %d%%\n", daily.Clouds)
		fmt.Printf("\n")
	}

	os.Exit(0)

}
