package repository

import (
	"fmt"
	"go-simpoku/database"
	"go-simpoku/src/lib"
	"go-simpoku/src/model"
)

type Auth struct {
	model.User
	Member model.Member `gorm:"foreignKey:UserID" json:"member"`
	Admin  model.Admin  `gorm:"foreignKey:UserID" json:"admin"`
}

func (auth *Auth) SignIn(role string) (user *Auth, err error) {
	password := *auth.User.Password
	preload := "Admin"
	table := "admins"
	if role == "member" {
		preload = "Member"
		table = "members"
	} else if role == "admin" {
		preload = "Admin"
		table = "admins"
	} else {
		return nil, lib.ErrInvalidRole
	}
	join := fmt.Sprintf("JOIN %s ON users.id = %s.user_id", table, table)
	if err = database.DB.Debug().
		Preload(preload).
		Joins(join).
		Where("username = ?", auth.User.Username).First(&auth).Error; err != nil {
		return nil, err
	}
	//validate password
	if !lib.IsPasswordValid(password, *auth.Password) {
		return auth, lib.ErrInvalidPassword
	}
	return auth, nil
}
