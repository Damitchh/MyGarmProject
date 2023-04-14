package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Photo represents the model for a photo
type Photo struct {
	GormModel
	Title    string `gorm:"title" json:"title" form:"title" validate:"required"`
	Caption  string `gorm:"caption" json:"caption" form:"caption"`
	PhotoUrl string `gorm:"photo_url" json:"photo_url" form:"photo_url" validate:"required"`
	UserID   uint   `gorm:"foreignKey:UserID" json:"user_id"`
	User     *User
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {

	_, errUpdate := govalidator.ValidateStruct(p)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
