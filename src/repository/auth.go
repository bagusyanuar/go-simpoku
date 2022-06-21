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

type AuthMember struct {
	model.User
	Member model.Member `gorm:"foreignKey:UserID" json:"member"`
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

func (auth *AuthMember) SignUp() (token *string, err error) {
	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err = tx.Debug().Create(&auth).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	jwt := lib.JWT{}
	claim := lib.JWTClaims{
		Unique:     auth.ID,
		Identifier: auth.Member.ID,
		Username:   auth.Username,
		Email:      auth.Email,
		Role:       "member",
	}
	accessToken, errorTokenize := jwt.GenerateToked(claim)
	if errorTokenize != nil {
		tx.Rollback()
		return nil, errorTokenize
	}
	tx.Commit()
	return &accessToken, nil
}
