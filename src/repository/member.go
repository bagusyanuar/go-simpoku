package repository

import (
	"go-simpoku/database"
	"go-simpoku/src/model"
)


type Member struct {
	model.Member
	User model.User `gorm:"foreignKey:UserID" json:"user"`
	Specialist []model.Specialist `gorm:"many2many:member_specialist;" json:"specialist"`
}

func (m *Member) Find(identifier string) (d *Member, err error) {
	if err = database.DB.Debug().
		Preload("User").
		Preload("Specialist").
		Where("id = ?", identifier).First(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

// func (Member) FindAll() (results []member, err error) {
// 	var member []member
// 	if err = database.DB.Debug().
// 		Preload("Specialist").
// 		Preload("User").
// 		Find(&member).
// 		Error; err != nil {
// 		return member, err
// 	}
// 	return member, nil
// }

// func (Member) Find(user_id uuid.UUID) (result *member, err error) {
// 	var data member
// 	if err = database.DB.Debug().
// 		Preload("Specialist").
// 		Preload("User").
// 		Where("id = ?", user_id).First(&data).Error; err != nil {
// 		return nil, err
// 	}
// 	return &data, nil
// }
