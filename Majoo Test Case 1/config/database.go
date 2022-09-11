package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB_HOST = "localhost"
	DB_USER = "majoo"
	DB_PASS = "majoo"
	DB_PORT = "3306"
	DB_NAME = "majoo"
	DB      *gorm.DB
)

func SetupDatabase() {
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)
	var err error

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	DB, err = gorm.Open(mysql.Open(URL), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic(err.Error())
	}

}
