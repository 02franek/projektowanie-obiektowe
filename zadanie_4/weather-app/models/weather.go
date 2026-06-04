package models

import (
	"time"

	"gorm.io/gorm"
)

type Weather struct {
	ID uint `gorm:"primarykey" json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	
	Location string `json:"location"`
	Temperature float64 `json:"temperature"`
	Precipitation float64 `json:"precipitation"`
	Description string `json:"description"`
}