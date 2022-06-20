package database

import "go-simpoku/src/model"

//table user
type User struct {
	model.User
}

//table admin
type Admin struct {
	model.Admin
	User model.User `gorm:"foreignKey:UserID"`
}

//table member
type Member struct {
	model.Member
	User       model.User         `gorm:"foreignKey:UserID"`
	Specialist []model.Specialist `gorm:"many2many:member_specialist;"`
}

//tabel specialist
type Specialist struct {
	model.Specialist
	Member []model.Member `gorm:"many2many:member_specialist;"`
}

//tabel event
type Event struct {
	model.Event
}

func Migrate() {
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Member{})
	DB.AutoMigrate(&Admin{})
	DB.AutoMigrate(&Specialist{})
	DB.AutoMigrate(&Event{})
}
