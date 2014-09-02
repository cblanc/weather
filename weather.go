package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"net/http"
	"net/url"
	"os"
	"strings"
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

func extractLocation(args []string) (string, error) {
	location := ""

	if len(args) < 2 {
		return location, errors.New("Please provide a location for a weather forecast. E.g. $ weather london")
	}
	for i := 1; i < len(args); i++ {
		location += args[i] + " "
	}
	return strings.Trim(location, " "), nil
}

func getForecast(location string) (*forecast, error) {
	var f forecast

	response, err := http.Get("http://api.openweathermap.org/data/2.5/forecast/daily?mode=json&units=metric&cnt=5&q=" + url.QueryEscape(location))

	defer response.Body.Close()

	if err != nil {
		return &f, err
	}

	if err := json.NewDecoder(response.Body).Decode(&f); err != nil {
		return &f, err
	}

	return &f, nil
}

func prettyPrintForecast(f *forecast) {
	fmt.Printf("\n5 day forecast for %s (%.5f, %.5f)\n\n", f.City.Name, f.City.Coord.Lon, f.City.Coord.Lat)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Day", "Forecast", "Temp (Range)", "Cloud Coverage (%)", "Wind Speed (m/s)", "Humidity (%)"})

	layout := "Monday, 01/02"

	for _, daily := range f.List {
		data := []string{}
		data = append(data, time.Unix(daily.Dt, 0).Format(layout))
		data = append(data, fmt.Sprintf("%s (%s)", daily.Weather[0].Main, daily.Weather[0].Description))
		data = append(data, fmt.Sprintf("%0.fC (%.0fC-%.0fC)\n", daily.Temp.Day, daily.Temp.Min, daily.Temp.Max))
		data = append(data, fmt.Sprintf("%d%%", daily.Clouds))
		data = append(data, fmt.Sprintf("%0.1f", daily.Speed))
		data = append(data, fmt.Sprintf("%d", daily.Humidity))
		table.Append(data)
	}

	table.Render()
}

func logError(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func main() {
	location, err := extractLocation(os.Args)

	if err != nil {
		logError(err)
	}

	f, err := getForecast(location)

	if err != nil {
		logError(err)
	}

	prettyPrintForecast(f)

	os.Exit(0)
}
