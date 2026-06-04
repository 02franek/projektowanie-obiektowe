package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type WeatherController struct{}

func (wc *WeatherController) HandleWeather(c echo.Context) error {
	method := c.Request().Method

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Weather controller test message",
		"method": method,
		"status": "THis is mock reponse",
	})
}

func main() {
	e := echo.New()

	wController := &WeatherController{}

	e.GET("/weather", wController.HandleWeather)
	e.POST("/weather", wController.HandleWeather)

	e.Logger.Fatal(e.Start(":8080"))
}