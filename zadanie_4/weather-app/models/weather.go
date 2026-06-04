package models

import "gorm.io/gorm"

type Weather struct {
	gorm.Model
	Location string `json:"location"`
	Temperature float64 `json:"temperature"`
	Precipitation float64 `json:"precipitation"`
	Description string `json:"description"`
}