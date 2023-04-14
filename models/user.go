package models

import (
	"MygarmProject/helpers"
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// User represents the model for a user
type User struct {
	GormModel
	Username string `gorm:"unique;not null" json:"username" form:"username" validate:"required,unique" valid:"required~Your username is required" `
	Email    string `gorm:"unique;not null" json:"email" form:"email" validate:"required,email,unique" valid:"required~Your email is required,email~Invalid email format"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have minimum 6 characters"'`
	Age      uint   `gorm:"not null" json:"age" form:"age" validate:"required,gt=8" valid:"required~Your age is required" `
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}

func (p *User) BeforeUpdate(tx *gorm.DB) (err error) {

	_, errUpdate := govalidator.ValidateStruct(p)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
