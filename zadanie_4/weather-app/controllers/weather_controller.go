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

		return c.JSON(http.StatusOK, weatherData)
	}

	var weathers []models.Weather
	db.DB.Find(&weathers)

	return c.JSON(http.StatusOK, weathers)
}