package repository

import (
	"go-simpoku/database"
	"go-simpoku/src/model"
)

type Specialist struct{}

func (Specialist) FindAll() (result *[]model.Specialist, err error) {
	var specialists *[]model.Specialist
	if err = database.DB.Debug().Find(&specialists).Error; err != nil {
		return nil, err
	}
	return specialists, nil
}

func (Specialist) Create(data *model.Specialist) (result *model.Specialist, err error) {
	if err = database.DB.Debug().Create(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
