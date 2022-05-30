package database

import "go-simpoku/src/model"

type MemberToUser struct {
	model.Member
	User model.User `gorm:"foreignKey:UserID"`
}
type AdminToUser struct {
	model.Admin
	User model.User `gorm:"foreignKey:UserID"`
}