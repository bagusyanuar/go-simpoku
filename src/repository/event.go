package repository

import (
	"go-simpoku/database"
	"go-simpoku/src/model"
)


type Event struct {
	model.Event
}

func (event *Event) Create() (e *Event, err error)   {
	if err = database.DB.Debug().Create(&event).Error; err != nil {
		return nil, err
	}
	return event, nil
}