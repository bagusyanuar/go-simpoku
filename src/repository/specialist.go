package repository

import (
	"go-simpoku/database"
	"go-simpoku/src/model"
)

type Specialist struct {
	model.Specialist
}

func (Specialist) FindAll(param string) (result []Specialist, err error) {
	var specialists []Specialist
	query := database.DB.Debug()
	if param != "" {
		query = query.Where("name LIKE ?", "%"+param+"%")
	}
	if err = query.Find(&specialists).Error; err != nil {
		return nil, err
	}
	return specialists, nil
}

func (s *Specialist) Create() (result *Specialist, err error) {
	if err = database.DB.Debug().Create(&s).Error; err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Specialist) Find(id string) (d *Specialist, err error)  {
	if err = database.DB.Debug().Where("id = ?", id).First(&s).Error; err != nil {
		return nil, err
	}
	return s, nil
}
