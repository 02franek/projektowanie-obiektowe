package main

import (
	"weather-app/controllers"
	"weather-app/db"
	"weather-app/services"

	"github.com/labstack/echo/v4"
)


func main() {
	db.InitDB()

	e := echo.New()
	wController := &controllers.WeatherController{
		Proxy: services.WeatherProxy{},
	}

	e.GET("/weather", wController.HandleWeather)
	e.POST("/weather", wController.HandleWeather)

	e.Logger.Fatal(e.Start(":8080"))
}