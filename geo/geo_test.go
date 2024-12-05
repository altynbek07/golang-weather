package geo_test

import (
	"demo/weather-app/geo"
	"testing"
)

func TestGetMyLocation(t *testing.T) {
	// Arrange - подготовка, expected результат, данные для функции
	city := "London"
	expected := geo.GeoData{
		City: "London",
	}

	// Act - выполняем функцию
	got, err := geo.GetMyLocation(city)

	// Assert - проверка результата с expected
	if err != nil {
		t.Error(err)
	}

	if got.City != expected.City {
		t.Errorf("Ожидалось %v, получили %v", expected, got)
	}
}

func TestGetMyLocationNoCity(t *testing.T) {
	// Arrange - подготовка, expected результат, данные для функции
	city := "Londondfsdews"

	// Act - выполняем функцию
	_, err := geo.GetMyLocation(city)

	// Assert - проверка результата с expected
	if err != geo.ErrorNoCity {
		t.Errorf("Ожидалось %v, получили %v", geo.ErrorNoCity, err)
	}
}
