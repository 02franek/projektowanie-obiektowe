package controllers

import (
	"errors"
	"net/http"
	"weather-app/db"
	"weather-app/models"
	"weather-app/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type WeatherController struct{
	Proxy services.WeatherProxy
}

type MultipleCitiesRequest struct {
	Locations []string `json:"locations"`
}

func (wc *WeatherController) HandleWeather(c echo.Context) error {
	method := c.Request().Method

	if method == http.MethodGet {
		locationParam := c.QueryParam("location")

		if locationParam != "" {
			weatherData, err := wc.Proxy.FetchWeather(locationParam)
			if err != nil {
				return c.JSON(http.StatusNotFound, map[string]string{
					"error": err.Error(),
				})
			}

			savedWeather := saveOrUpdateWeather(weatherData)
			return c.JSON(http.StatusOK, savedWeather)
		}

		var weathers []models.Weather
		db.DB.Preload("Forecasts").Find(&weathers)
		return c.JSON(http.StatusOK, weathers)
	}

	if method == http.MethodPost {
		var req MultipleCitiesRequest

		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Incorrect JSON data format",
			})
		}

		var results []models.Weather

		for _, location := range req.Locations {
			weatherData, err := wc.Proxy.FetchWeather(location)
			if err == nil {
				savedWeather := saveOrUpdateWeather(weatherData)
				results = append(results, savedWeather)
			}
		}

		return c.JSON(http.StatusOK, results)
	}

	return c.JSON(http.StatusMethodNotAllowed, map[string]string{
		"error": "Method not allowed",
	})
		
}

func saveOrUpdateWeather(weatherData *models.Weather) models.Weather {
	var existingWeather models.Weather

	result := db.DB.Where("location = ?", weatherData.Location).First(&existingWeather)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				db.DB.Create(weatherData)
			return *weatherData
		}
	}

	existingWeather.Temperature = weatherData.Temperature
	existingWeather.Precipitation = weatherData.Precipitation
	existingWeather.Description = weatherData.Description
	db.DB.Save(&existingWeather)

	db.DB.Model(&existingWeather).Association("Forecasts").Replace(weatherData.Forecasts)
	existingWeather.Forecasts = weatherData.Forecasts

	return existingWeather
}