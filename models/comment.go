package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Comment represents the model for a comment
type Comment struct {
	GormModel
	UserID  uint   `gorm:"foreignKey:UserID"`
	PhotoID uint   `gorm:"foreignKey:PhotoID" json:"photo_id" form:"photo_id" validate:"required"`
	Message string `gorm:"not_null" json:"message" form:"message" validate:"required"`
	Photo   *Photo
	User    *User
}

func (p *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}

func (p *Comment) BeforeUpdate(tx *gorm.DB) (err error) {

	_, errUpdate := govalidator.ValidateStruct(p)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
