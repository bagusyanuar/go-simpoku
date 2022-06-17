package database

import "go-simpoku/src/model"

func Migrate() {
	DB.AutoMigrate(&model.User{})
	// DB.AutoMigrate(&model.Admin{})
	DB.AutoMigrate(&model.Member{})
	DB.AutoMigrate(&model.Specialist{})
}
