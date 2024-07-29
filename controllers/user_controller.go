package controllers

import (
	"net/http"

	"github.com/gaisuke/mygram-api/dto"
	"github.com/gaisuke/mygram-api/helpers"
	"github.com/gaisuke/mygram-api/services"
	"github.com/gin-gonic/gin"
)

func UserRegister(ctx *gin.Context) {
	var userRegisterRequestDto dto.RegisterUserRequestDto

	if err := ctx.ShouldBind(&userRegisterRequestDto); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, err := services.UserRegister(userRegisterRequestDto)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	user, err := services.GetUserById(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	res := dto.RegisterUserResponseDto{
		ID:       user.ID,
		Email:    user.Email,
		Age:      user.Age,
		Username: user.Username,
	}

	ctx.JSON(http.StatusCreated, res)
}

func UserLogin(ctx *gin.Context) {
	var userLoginRequestDto dto.LoginUserRequestDto

	if err := ctx.ShouldBind(&userLoginRequestDto); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	passwordReq := userLoginRequestDto.Password

	user, err := services.LoginUser(userLoginRequestDto.Email)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return

	}

	comparePass, err := helpers.ComparePass(user.Password, passwordReq)

	if !comparePass {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token := helpers.GenerateToken(user.ID, user.Email)

	res := dto.LoginUserResponseDto{
		Token: token,
	}

	ctx.JSON(http.StatusOK, res)
}
