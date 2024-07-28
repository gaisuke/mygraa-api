package dto

type RegisterUserRequestDto struct {
	Age      int    `json:"age" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type RegisterUserResponseDto struct {
	ID       uint   `json:"id"`
	Age      int    `json:"age" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required"`
}
