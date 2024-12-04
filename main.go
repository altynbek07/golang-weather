package main

import (
	"demo/weather-app/geo"
	"demo/weather-app/weather"
	"flag"
	"fmt"
)

func main() {
	city := flag.String("city", "", "Город")
	format := flag.Int("format", 1, "Формат вывода погоды")
	flag.Parse()

	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}

	weatherData := weather.GetWeather(*geoData, *format)
	fmt.Println(weatherData)
}
