package api

import (
	"net/http"
)

func SendApiWeatherRequest() (*http.Response, error) {
	resp, err := http.Get(OpenMeteoURL)
	if err != nil {
		panic(err)
		return nil, err
	}
	return resp, nil
}
