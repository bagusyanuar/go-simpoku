package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Member struct {
	ID         uuid.UUID    `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	UserID     uuid.UUID    `gorm:"column:user_id;type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;not null" json:"user_id"`
	Name       string       `gorm:"type:varchar(255);not null;" json:"name"`
	Phone      string       `gorm:"type:varchar(16);not null;" json:"phone"`
	Avatar     string       `gorm:"type:text;not null;" json:"avatar"`
	CreatedAt  time.Time    `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt  time.Time    `gorm:"column:updated_at;not null" json:"updated_at"`
	User       User         `gorm:"foreignKey:UserID" json:"user"`
	Specialist []Specialist `gorm:"many2many:member_specialist;" json:"specialist"`
}

// type Member struct {
// 	BaseMember
// 	User       BaseUser         `gorm:"foreignKey:UserID"`
// 	Specialist []BaseSpecialist `gorm:"many2many:member_specialist;"`
// }

func (member *Member) BeforeCreate(tx *gorm.DB) (err error) {
	member.ID = uuid.New()
	member.CreatedAt = time.Now()
	member.UpdatedAt = time.Now()
	return
}

func (Member) TableName() string {
	return "members"
}
