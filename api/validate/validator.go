package validate

import (
	"fmt"
	"regexp"
	"unicode"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitValidation() error {
	// カスタムバリデーションルールを登録
	validate := binding.Validator.Engine().(*validator.Validate)
	err := validate.RegisterValidation("password", passwordValidation)
	if err != nil {
		return fmt.Errorf("failed to register custom validation: %w", err)
	}
	return nil
}

// passwordValidationはパスワードのバリデーションを行う関数
// TODO: テストを書く
func passwordValidation(fl validator.FieldLevel) bool {
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
