package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"weather-app/models"
)

type WeatherProxy struct{}

type geoResponse struct {
	Results []struct {
		Latitude float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Name string `json:"name"`
	} `json:"results"`
}

type weatherResponse struct {
	Current struct {
		Temperature float64 `json:"temperature_2m"`
		Precipiration float64 `json:"precipitation"`
	} `json:"current"`
	Daily struct {
		Time []string `json:"time"`
		TemperatureMax []float64 `json:"temperature_2m_max"`
		TemperatureMin []float64 `json:"temperature_2m_min"`
		PrecipitationSum []float64 `json:"precipitation_sum"`
	} `json:"daily"`
}

func (wp *WeatherProxy) FetchWeather(city string) (*models.Weather, error) {
	geoURL := fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1", city)
	res, err := http.Get(geoURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var geo geoResponse
	if err := json.NewDecoder(res.Body).Decode(&geo); err != nil {
		return nil, err
	}

	if len(geo.Results) == 0 {
		return nil, fmt.Errorf("Did not find coordinates for location: %s", city)
	}

	lat := geo.Results[0].Latitude
	lon := geo.Results[0].Longitude
	resolvedName := geo.Results[0].Name

	weatherURL := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current=temperature_2m,precipitation&daily=temperature_2m_max,temperature_2m_min,precipitation_sum&timezone=auto&forecast_days=3", lat, lon)
	wRes, err := http.Get(weatherURL)
	if err != nil {
		return nil, err
	}
	defer wRes.Body.Close()
	
	var wData weatherResponse
	if err := json.NewDecoder(wRes.Body).Decode(&wData); err != nil {
		return nil, err
	}

	var forecasts []models.Forecast
	for i := range wData.Daily.Time {
		forecasts = append(forecasts, models.Forecast{
			Date: wData.Daily.Time[i],
			MaxTemperature: wData.Daily.TemperatureMax[i],
			MinTemperature: wData.Daily.TemperatureMin[i],
			PrecipitationSum: wData.Daily.PrecipitationSum[i],
		})
	}

	weather := &models.Weather{
		Location: resolvedName,
		Temperature: wData.Current.Temperature,
		Precipitation: wData.Current.Precipiration,
		Description: "Provided by Open-Meteo",
		Forecasts: forecasts,
	}
	
	return weather, nil
}