package db

import "github.com/jinzhu/gorm"

type ProbTmpM struct {
	gorm.Model
	Host     string `gorm:"type:text"`
	Ports    string `gorm:"type:text"`
	Services string `gorm:"type:text"`
	Fingers  string `gorm:"type:text"`
}
