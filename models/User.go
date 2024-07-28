package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gaisuke/mygram-api/helpers"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username" gorm:"unique;varchar(100)" valid:"required~Your username is required"`
	Email     string    `json:"email" gorm:"unique;varchar(100)" valid:"email,required~Your email is required,email~Invalid email format"`
	Password  string    `json:"password" valid:"required~Your password is required,minstringlength(6)~Minimum password length is 6"`
	Age       int       `json:"age" valid:"required~Your age is required,range(9|100)~Your age must be minimum of 9"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(u)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
