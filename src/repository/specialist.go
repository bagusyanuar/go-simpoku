package repository

import (
	"go-simpoku/database"
	"go-simpoku/src/model"
)

type Specialist struct{}

func (Specialist) FindAll() (result *[]model.BaseSpecialist, err error) {
	var specialists *[]model.BaseSpecialist
	if err = database.DB.Debug().Find(&specialists).Error; err != nil {
		return nil, err
	}
	return specialists, nil
}

func (Specialist) Create(data *model.BaseSpecialist) (result *model.BaseSpecialist, err error) {
	if err = database.DB.Debug().Create(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
