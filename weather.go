package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// type dailyForecast {

// }

type forecast struct {
	City struct {
		Name  string
		Coord struct {
			Lon float64
			Lat float64
		}
	}
	// List []dailyForecast
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

	fmt.Printf("Location: %s\n", f.City.Name)
	fmt.Printf("Geolocation: %.5f,%.5f\n", f.City.Coord.Lon, f.City.Coord.Lat)

	os.Exit(0)

}
