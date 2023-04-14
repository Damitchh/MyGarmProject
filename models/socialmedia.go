package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// SocialMedia represents the model for a socialmedia
type SocialMedia struct {
	GormModel
	Name           string `gorm:"not_null" json:"name" form:"name" validate:"required"`
	SocialMediaUrl string `gorm:"not_null" json:"social_media_url" form:"social_media_url" validate:"required"`
	UserID         uint   `gorm:"unique;foreignKey:UserID"`
	User           *User
}

func (p *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}

func (p *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {

	_, errUpdate := govalidator.ValidateStruct(p)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
