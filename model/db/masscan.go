package db

import "github.com/jinzhu/gorm"

type MascM struct {
	gorm.Model
	Host  string
	Ports string
}
