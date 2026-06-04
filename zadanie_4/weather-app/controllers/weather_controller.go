package controllers

import (
	"net/http"
	"weather-app/db"
	"weather-app/models"

	"github.com/labstack/echo/v4"
)

type WeatherController struct{}

func (wc *WeatherController) HandleWeather(c echo.Context) error {
	var weathers []models.Weather

	db.DB.Find(&weathers)

	return c.JSON(http.StatusOK, weathers)
}