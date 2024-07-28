package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	ID             uint   `json:"id"`
	Name           string `json:"name" valid:"required~Your social media name is required"`
	SocialMediaUrl string `json:"social_media_url" valid:"required~Your social media url is required"`
	UserId         uint   `json:"user_id"`
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (s *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(s)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
