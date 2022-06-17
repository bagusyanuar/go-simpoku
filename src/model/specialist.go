package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseSpecialist struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Slug      string    `gorm:"type:varchar(255);not null" json:"slug"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

func (specialist *BaseSpecialist) BeforeCreate(tx *gorm.DB) (err error) {
	specialist.CreatedAt = time.Now()
	specialist.UpdatedAt = time.Now()
	return
}

func (BaseSpecialist) TableName() string {
	return "specialists"
}


type Specialist struct {
	BaseSpecialist
	Member []Member `gorm:"many2many:member_specialist;"`
}