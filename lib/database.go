package lib

import (
	"fmt"
	"github.com/binganao/Taio/common"
	db2 "github.com/binganao/Taio/model/db"
	"github.com/binganao/Taio/pkg/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := common.DATABASE_DRIVER
	host := common.DATABASE_HOST
	port := common.DATABASE_PORT
	database := common.DATABASE_DATABASE
	databaseUser := common.DATABASE_USERNAME
	databasePwd := common.DATABASE_PASSWORD
	charset := common.DATABASE_CHARSET
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		databaseUser,
		databasePwd,
		host,
		port,
		database,
		charset)

	db, err := gorm.Open(driverName, args)
	if err != nil {
		logger.Error("Database connect failed!")
		os.Exit(0)
	}

	db.AutoMigrate(&db2.MascM{})
	db.AutoMigrate(&db2.NamM{})

	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
