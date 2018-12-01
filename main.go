package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// https://mholt.github.io/json-to-go/
type Weather struct {
	Location     string `json:"location"`
	Weather      string `json:"weather"`
	Temperature  int    `json:"temperature"`
	Celsius      bool   `json:"celsius"`
	Date         string `json:"date"`
	TempForecast []int  `json:"temp_forecast"`
	Wind		  Wind  `json:"wind"`
}
type Wind struct {
	Direction string `json:"direction"`
	Speed     int    `json:"speed"`
}

func GetWeatherData() Weather {
	return Weather{
		Location:     "Amsterdam",
		Weather:      "sunny",
		Temperature:  22,
		Celsius:      true,
		Date:         "2018-06-22",
		TempForecast: []int { 25, 26, 24, 20, 21, 22},
		Wind: Wind {
			Direction: "SE",
			Speed: 15,
		},
	}
}

func main() {
	bytes, _ := json.Marshal(GetWeatherData())
	fmt.Println(string(bytes))
	data, err := ioutil.ReadFile("weather.json")
	if err != nil {
		panic(err)
	}
	weather := Weather{}
	err2 := json.Unmarshal(data, &weather) // pointer to weather
	if err != nil {
		panic(err2)
	}
	fmt.Print(weather)
}