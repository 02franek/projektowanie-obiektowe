package controllers

import (
	"net/http"
	"weather-app/db"
	"weather-app/models"
	"weather-app/services"

	"github.com/labstack/echo/v4"
)

type WeatherController struct{
	Proxy services.WeatherProxy
}

func (wc *WeatherController) HandleWeather(c echo.Context) error {
	locationParam := c.QueryParam("location")

	if locationParam != "" {
		weatherData, err := wc.Proxy.FetchWeather(locationParam)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": err.Error(),
			})
		}

		var existingWeather models.Weather

		result := db.DB.Where("location = ?", weatherData.Location).First(&existingWeather)

		if result.Error == nil {
			existingWeather.Temperature = weatherData.Temperature
			existingWeather.Precipitation = weatherData.Precipitation
			existingWeather.Description = weatherData.Description
			db.DB.Save(&existingWeather)

			db.DB.Model(&existingWeather).Association("Forecasts").Replace(weatherData.Forecasts)

			existingWeather.Forecasts = weatherData.Forecasts
			return c.JSON(http.StatusOK, existingWeather)
		} else {
			db.DB.Create(weatherData)
			return c.JSON(http.StatusOK, weatherData)
		}
	}

	var weathers []models.Weather
	db.DB.Preload("Forecasts").Find(&weathers)

	return c.JSON(http.StatusOK, weathers)
}