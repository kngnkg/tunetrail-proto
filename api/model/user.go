package model

import (
	"regexp"
	"time"
	"unicode"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Id        int       `json:"id" db:"id"`
	UserName  string    `json:"userName" db:"user_name"`
	Name      string    `json:"name" db:"name"`
	Password  string    `json:"password" db:"password"`
	Email     string    `json:"email" db:"email"`
	IconUrl   string    `json:"iconUrl" db:"icon_url"`
	Bio       string    `json:"bio" db:"bio"`
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
	IconUrl  string `json:"iconUrl" binding:"required,url"`
	Bio      string `json:"bio" binding:"required,max=1000"`
}

// PasswordValidationFunctionはパスワードのバリデーションを行う関数
// TODO: テストを書く
func PasswordValidationFunction(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	// 8文字以上20文字以下である
	if len(password) < 8 || len(password) > 20 {
		return false
	}

	// 英数字のみである
	alphanumeric := regexp.MustCompile("^[a-zA-Z0-9]*$")
	if !alphanumeric.MatchString(password) {
		return false
	}

	hasUpper, hasLower, hasNumber := false, false, false

	for _, r := range password {
		switch {
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsNumber(r):
			hasNumber = true
		}
	}

	// 1文字以上の大文字が含まれている
	// 1文字以上の小文字が含まれている
	// 1文字以上の数字が含まれている
	return hasUpper && hasLower && hasNumber
}
