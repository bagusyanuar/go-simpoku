package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Event struct {
	ID          uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	Name        string         `gorm:"type:varchar(255);not null;" json:"name"`
	Description string         `gorm:"type:text;not null;" json:"deskription"`
	Slug        string         `gorm:"type:varchar(255);not null" json:"slug"`
	Image       string         `gorm:"type:text;not null;" json:"image"`
	StartAt     datatypes.Date `gorm:"type:date;not null;" json:"start_at"`
	FinishAt    datatypes.Date `gorm:"type:date;not null;" json:"finish_at"`
	Location    string         `gorm:"type:text;not null;" json:"location"`
	CreatedAt   time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
}

func (event *Event) BeforeCreate(tx *gorm.DB) (err error) {
	event.ID = uuid.New()
	event.CreatedAt = time.Now()
	event.UpdatedAt = time.Now()
	return
}

func (Event) TableName() string {
	return "events"
}
