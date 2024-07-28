package models

import "time"

type User struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email" valid:"email,required~Your email is required,email~Invalid email format"`
	Password  string    `json:"password"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func BeforeCreate(user *User) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
}
