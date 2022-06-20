package repository

import (
	"go-simpoku/database"
	"go-simpoku/src/lib"
	"go-simpoku/src/model"
)


type AdminAuth struct {
	model.User
	Admin model.Admin `gorm:"foreignKey:UserID" json:"admin"`
}

func (auth *AdminAuth) SignIn() (admin *AdminAuth, err error) {
	password := *auth.User.Password
	if err = database.DB.Debug().
	Preload("Admin").
	Joins("JOIN admins ON users.id = admins.user_id").
	Where("username = ?", auth.User.Username).First(&auth).Error; err != nil {
		return nil, err
	}
	//validate password
	if !lib.IsPasswordValid(password, *auth.Password) {
		
		return auth, lib.ErrInvalidPassword
	}
	return auth, nil
}
