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

type member struct {
	model.Member
	User       model.User         `gorm:"foreignKey:UserID" json:"user"`
	Specialist []model.Specialist `gorm:"many2many:member_specialist;" json:"specialists"`
}

func (Member) FindAll() (results []member, err error) {
	var member []member
	if err = database.DB.Debug().
		Preload("Specialist").
		Preload("User").
		Find(&member).
		Error; err != nil {
		return member, err
	}
	return member, nil
}

func (Member) Find(user_id uuid.UUID) (result *member, err error) {
	var data member
	if err = database.DB.Debug().
		Preload("Specialist").
		Preload("User").
		Where("id = ?", user_id).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
