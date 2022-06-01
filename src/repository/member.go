package repository

import (
	"go-simpoku/database"
	"go-simpoku/src/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Member struct {
	TX *gorm.DB
}

func (Member) FindAll() (results []model.Member, err error) {
	var member []model.Member
	if err = database.DB.Debug().
		Preload("Specialist").
		Preload("User").
		Find(&member).
		Error; err != nil {
		return member, err
	}
	return member, nil
}
func (Member) Find(user_id uuid.UUID) (result *model.Member, err error) {
	var member model.Member
	if err = database.DB.Debug().
		Preload("Specialist").
		Preload("User").
		Where("id = ?", user_id).First(&member).Error; err != nil {
		return nil, err
	}
	return &member, nil
}
