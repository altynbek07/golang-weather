package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type GeoData struct {
	City string `json:"city"`
}

type CityPopulationResponse struct {
	Error bool `json:"error"`
}

func GetMyLocation(city string) (*GeoData, error) {
	if city != "" {
		isCity := checkCity(city)
		if !isCity {
			panic("Не существует города " + city)
		}
		return &GeoData{
			City: city,
		}, nil
	}

	// resp, err := http.Get("https://ipapi.co/json/")
	// if err != nil {
	// 	return nil, err
	// }

	req, err := http.NewRequest("GET", "https://ipapi.co/json/", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0")
	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("NOT200")
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var geo GeoData
	json.Unmarshal(body, &geo)
	return &geo, nil
}

func checkCity(city string) bool {
	postBody, _ := json.Marshal(map[string]string{"city": city})

	req, err := http.NewRequest("POST", "https://countriesnow.space/api/v0.1/countries/population/cities", bytes.NewBuffer(postBody))
	if err != nil {
		return false
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0")
	resp, err := new(http.Client).Do(req)
	if err != nil {
		return false
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}

	var cityPopulationResponse CityPopulationResponse
	json.Unmarshal(body, &cityPopulationResponse)
	return !cityPopulationResponse.Error
}
