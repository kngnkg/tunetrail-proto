package validate

import (
	"fmt"
	"regexp"
	"unicode"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

const (
	MinPasswordLength = 8
	MaxPasswordLength = 64
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

	if len(password) < MinPasswordLength || len(password) > MaxPasswordLength {
		return false
	}

	symbol := regexp.MustCompile(`[!@#\$%\^&\*\(\)-_=\+{}\[\]:;"'<>,\.\?/\\~\|]`)

	hasUpper, hasLower, hasNumber, hasSymbol := false, false, false, false

	for _, r := range password {
		switch {
		case unicode.IsUpper(r):
			hasUpper = true // 1文字以上の大文字が含まれている
		case unicode.IsLower(r):
			hasLower = true // 1文字以上の小文字が含まれている
		case unicode.IsNumber(r):
			hasNumber = true // 1文字以上の数字が含まれている
		case symbol.MatchString(string(r)):
			hasSymbol = true // 1文字以上の記号が含まれている
		}
	}

	return hasUpper && hasLower && hasNumber && hasSymbol
}
