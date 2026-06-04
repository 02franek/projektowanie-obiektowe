package models

import (
	"time"

	"gorm.io/gorm"
)

type Forecast struct {
	ID uint `gorm:"primarykey" json:"-"`
	WeatherID uint `json:"-"`
	Date string `json:"date"`
	MaxTemperature float64 `json:"max_temperature"`
	MinTemperature float64 `json:"min_temperature"`
	PrecipitationSum float64 `json:"precipitation_sum"`
}

type Weather struct {
	ID uint `gorm:"primarykey" json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	
	Location string `json:"location"`
	Temperature float64 `json:"temperature"`
	Precipitation float64 `json:"precipitation"`
	Description string `json:"description"`

	Forecasts []Forecast `gorm:"foreignKey:WeatherID" json:"forecast"`
}