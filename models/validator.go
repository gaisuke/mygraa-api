package models

import "gorm.io/gorm"

type Validator interface {
	BeforeCreate(tx *gorm.DB) error
	BeforeUpdate(tx *gorm.DB) error
}
