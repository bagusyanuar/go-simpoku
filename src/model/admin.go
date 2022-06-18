package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admin struct {
	ID        uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	UserID    uuid.UUID `gorm:"column:user_id;type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;not null" json:"user_id"`
	Name      string    `gorm:"type:varchar(255);not null;" json:"name"`
	Phone     string    `gorm:"type:varchar(16);not null;" json:"phone"`
	Avatar    string    `gorm:"type:text;not null;" json:"avatar"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
	User      User      `gorm:"foreignKey:UserID"`
}

func (admin *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	admin.ID = uuid.New()
	admin.CreatedAt = time.Now()
	admin.UpdatedAt = time.Now()
	return
}

func (Admin) TableName() string {
	return "admins"
}
