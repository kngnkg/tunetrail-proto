package model

import (
	"time"
)

type User struct {
	Id        int       `json:"id" db:"id" binding:"required"`
	UserName  string    `json:"userName" db:"user_name" binding:"required,min=3,max=20"`
	Name      string    `json:"name" db:"name" binding:"required,min=3,max=20"`
	Password  string    `json:"password" db:"password" binding:"required,password"`
	Email     string    `json:"email" db:"email" binding:"required,email"`
	IconUrl   string    `json:"iconUrl" db:"icon_url" binding:"required,url"`
	Bio       string    `json:"bio" db:"bio" binding:"required,max=1000"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

type Users []User

// UserRegisterRequestはユーザー登録時のリクエストをバインドする構造体
type UserRegisterRequest struct {
	UserName string `json:"userName" binding:"required,min=3,max=20"`
	Name     string `json:"name" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,password"`
	Email    string `json:"email" binding:"required,email"`
}
