package db

import "github.com/jinzhu/gorm"

type FingM struct {
	gorm.Model
	Url    string
	WebApp string
	Server string
}
