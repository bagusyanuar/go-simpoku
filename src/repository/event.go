package repository

import (
	"go-simpoku/database"
	"go-simpoku/src/model"
)

type Event struct {
	model.Event
}

func (event *Event) Create() (e *Event, err error) {
	if err = database.DB.Debug().Create(&event).Error; err != nil {
		return nil, err
	}
	return event, nil
}

func (Event) FindAll() (d []Event, err error) {
	var data []Event
	if err = database.DB.Debug().Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (event Event) FindBySlug() (d *Event, err error)  {
	if err = database.DB.Debug().Where("slug = ?", event.Slug).First(&event).Error; err != nil {
		return nil, err
	}
	return &event, nil
}
