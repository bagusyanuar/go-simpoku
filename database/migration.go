package database

import "go-simpoku/src/model"

func Migrate() {
	DB.AutoMigrate(
		&model.User{},
		&model.Admin{},
		&model.Member{},
		&model.Specialist{},
	)
}
