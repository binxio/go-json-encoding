package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
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

type WeatherDate struct {
	Location     string `json:"location"`
	Weather      string `json:"weather"`
	Temperature  int    `json:"temperature"`
	Celsius      bool   `json:"celsius"`
	Date         Date    `json:"date"`
	TempForecast []int  `json:"temp_forecast"`
	Wind		  Wind  `json:"wind"`
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

type Date struct {
	value time.Time
}

func GetWeatherDataTime() WeatherDate {
	t, err := time.Parse(time.RFC3339, "2018-06-22T15:04:05Z")
	if err != nil {
		panic(err)
	}
	return WeatherDate{
		Location:     "Amsterdam",
		Weather:      "sunny",
		Temperature:  22,
		Celsius:      true,
		Date:         Date { value: t },
		TempForecast: []int { 25, 26, 24, 20, 21, 22},
		Wind: Wind {
			Direction: "SE",
			Speed: 15,
		},
	}
}

func JsonWithStringDate() {
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

func JsonWithTime() {
	bytes, _ := json.Marshal(GetWeatherDataTime())
	fmt.Println(string(bytes))
	data, err := ioutil.ReadFile("weather.json")
	if err != nil {
		panic(err)
	}
	weather := WeatherDate{}
	err2 := json.Unmarshal(data, &weather) // pointer to weather
	if err != nil {
		panic(err2)
	}
	fmt.Print(weather)
}

func (w *Date) UnmarshalJSON(b []byte) error {
	x := string("")
	e := json.Unmarshal(b, &x)
	if e != nil {
		return e
	}
	t, err := time.Parse(time.RFC3339, x)
	if err != nil {
		return err
	}
	w.value = t
	return nil
}

func (w Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(w.value.Format(time.RFC3339))
}

func main() {
	JsonWithStringDate()
	JsonWithTime()

}