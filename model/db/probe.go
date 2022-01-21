package db

import "github.com/jinzhu/gorm"

type ProbM struct {
	gorm.Model
	Host     string
	Ports    string
	Services string
	Fingers  string
}
