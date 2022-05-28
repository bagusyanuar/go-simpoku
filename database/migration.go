package database

import "go-simpoku/src/model"

func Migrate() {
	DB.AutoMigrate(&model.User{})
}