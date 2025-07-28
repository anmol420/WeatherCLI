package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Response struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

type Location struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

type Current struct {
	Temperature float64   `json:"temp_c"`
	Humidity    int       `json:"humidity"`
	Condition   Condition `json:"condition"`
}

type Condition struct {
	Text string `json:"text"`
}

func CurrentWeather(location string) (Response, error) {
	err := godotenv.Load()
	if err != nil {
		return Response{}, fmt.Errorf("Error loading .env file: %w", err)
	}
	apiKey := os.Getenv("WEATHER_API_KEY")
	apiUrl := "http://api.weatherapi.com/v1/current.json?key=" + apiKey + "&q=" + location
	res, err := http.Get(apiUrl)
	if err != nil {
		return Response{}, fmt.Errorf("Error fetching data: %w", err)
	}
	defer res.Body.Close()
	resData, err := io.ReadAll(res.Body)
	if err != nil {
		return Response{}, fmt.Errorf("Error reading response body: %w", err)
	}
	var response Response
	err = json.Unmarshal(resData, &response)
	if err != nil {
		return Response{}, fmt.Errorf("Error unmarshalling response: %w", err)
	}
	return response, nil
}