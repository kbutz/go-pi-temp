package gorm

import (
	"github.com/golang/glog"
)

//Migrate : Initializes all models in db
func Migrate() error {
	glog.Info("Running object Migrations...")

	/*
			===========================================
			Keep these alphabetical for easy search
		  ===========================================
	*/

	glog.Info("Creating TemperatureReading Table")
	err := DB.AutoMigrate(&TemperatureReading{}).Error
	if err != nil {
		glog.Info(err)
		return err
	}

	return err
}
