package database

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/gorm"
)


var DB *gorm.DB

func Build() string {
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort, errParse := strconv.Atoi(os.Getenv("DB_PORT"))
	if errParse != nil {
		panic("Error : Failed Port Conversion")
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
}