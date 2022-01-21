package db

import "github.com/jinzhu/gorm"

type NamM struct {
	gorm.Model
	Host    string
	Service string
}
