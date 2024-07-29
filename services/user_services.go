package services

import (
	"github.com/gaisuke/mygram-api/database"
	"github.com/gaisuke/mygram-api/dto"
	"github.com/gaisuke/mygram-api/models"
)

func UserRegister(userRegisterRequestDto dto.RegisterUserRequestDto) (id uint, err error) {
	db := database.GetDB()
	user := models.User{
		Email:    userRegisterRequestDto.Email,
		Password: userRegisterRequestDto.Password,
		Age:      userRegisterRequestDto.Age,
		Username: userRegisterRequestDto.Username,
	}

	if err := db.Create(&user).Error; err != nil {
		return user.ID, err
	}

	return user.ID, nil
}

func LoginUser(email string) (user models.User, err error) {
	db := database.GetDB()

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func GetUserById(id uint) (user models.User, err error) {
	db := database.GetDB()

	if err := db.First(&user, id).Error; err != nil {
		return user, err
	}

	return user, nil
}
