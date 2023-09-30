package model

import (
	"time"

	"github.com/google/uuid"
)

type UserID string

// ユーザーIDを生成する
func NewUserId() UserID {
	return UserID(uuid.New().String())
}

type User struct {
	Id       UserID `json:"id" db:"id" binding:"required,uuid4"`
	UserName string `json:"userName" db:"user_name" binding:"required,min=3,max=20"`
	Name     string `json:"name" db:"name" binding:"required,min=3,max=20"`
	IconUrl  string `json:"iconUrl" db:"icon_url" binding:"required,url"`
	Bio      string `json:"bio" db:"bio" binding:"required,max=1000"`
	Password string `json:"password" binding:"password"`
	Email    string `json:"email" binding:"email"`
	FollowInfo
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

type AuthInfo struct {
	Id    UserID `json:"id" binding:"required,uuid4"`
	Email string `json:"email" binding:"email"`
}

type UserRegistrationData struct {
	UserName string `json:"userName" binding:"required,min=3,max=20"`
	Name     string `json:"name" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,password"`
	Email    string `json:"email" binding:"required,email"`
}

type UserSignInData struct {
	UserName string `json:"userName" binding:"omitempty,min=3,max=20"`
	Email    string `json:"email" binding:"omitempty,email"`
	Password string `json:"password" binding:"required,password"`
}

type UserUpdateData struct {
	Id       UserID `json:"id" db:"id" binding:"required"`
	UserName string `json:"userName" db:"user_name" binding:"required,min=3,max=20"`
	Name     string `json:"name" db:"name" binding:"required,min=3,max=20"`
	IconUrl  string `json:"iconUrl" db:"icon_url" binding:"required,url"`
	Bio      string `json:"bio" db:"bio" binding:"required,max=1000"`
	Password string `json:"password" db:"password" binding:"required,password"`
	Email    string `json:"email" db:"email" binding:"required,email"`
}
