package gorm

import "github.com/jinzhu/gorm"

type TemperatureReading struct {
	gorm.Model

	Location       string
	TemperatureInF float32
	TemperatureInC float32
}

func (t *TemperatureReading) Create() error {
	err := DB.Create(&t).Error
	return err
}
