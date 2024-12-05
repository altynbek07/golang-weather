package weather_test

import (
	"demo/weather-app/geo"
	"demo/weather-app/weather"
	"strings"
	"testing"
)

func TestGetWeather(t *testing.T) {
	// Arrange - подготовка, expected результат, данные для функции
	expected := "London"
	geoData := geo.GeoData{
		City: expected,
	}
	format := 3

	// Act - выполняем функцию
	result, err := weather.GetWeather(geoData, format)

	// Assert - проверка результата с expected
	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(result, expected) {
		t.Errorf("Ожидалось %v, получили %v", expected, result)
	}
}

var testCases = []struct {
	name   string
	format int
}{
	{name: "0 format", format: 0},
	{name: "Big format", format: 147},
	{name: "Minus format", format: -1},
}

func TestGetWeatherWrongFormat(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange - подготовка, expected результат, данные для функции
			expected := "London"
			geoData := geo.GeoData{
				City: expected,
			}

			// Act - выполняем функцию
			_, err := weather.GetWeather(geoData, tc.format)

			// Assert - проверка результата с expected
			if err != weather.ErrorWrongFormat {
				t.Errorf("Ожидалось %v, получили %v", weather.ErrorWrongFormat, err)
			}
		})
	}
}
