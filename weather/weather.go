package weather

import (
	"demo/weather-app/geo"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var ErrorWrongFormat = errors.New("WRONG_FORMAT")

func GetWeather(geo geo.GeoData, format int) (string, error) {
	if format < 1 || format > 4 {
		return "", ErrorWrongFormat
	}
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("ERROR_URL")
	}

	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", baseUrl.String(), nil)
	if err != nil {
		return "", errors.New("ERROR_INIT_REQUEST")
	}
	req.Header.Set("User-Agent", "Mozilla/5.0")
	resp, err := new(http.Client).Do(req)
	if err != nil {
		return "", errors.New("ERROR_HTTP")
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("ERROR_READ_BODY")
	}

	return string(body), nil
}
