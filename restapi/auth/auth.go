package auth

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/google/uuid"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . AuthProvider
type AuthProvider interface {
	SignUpWithContext(ctx context.Context, input *cognitoidentityprovider.SignUpInput, opts ...request.Option) (*cognitoidentityprovider.SignUpOutput, error)
	AdminInitiateAuth(input *cognitoidentityprovider.AdminInitiateAuthInput) (*cognitoidentityprovider.AdminInitiateAuthOutput, error)
}

// Cognitoから返されるトークン
type Tokens struct {
	Id      string
	Access  string
	Refresh string
}

type authConfig struct {
	userPoolId          string
	cognitoClientId     string
	cognitoClientSecret string
}

type Auth struct {
	provider AuthProvider
	*authConfig
}

var (
	ErrEmailAlreadyExists = errors.New("auth: email already exists")
	ErrInvalidPassword    = errors.New("auth: invalid password")
	ErrUserNotConfirmed   = errors.New("auth: user not confirmed")
	ErrUserSubIsNil       = errors.New("auth: user sub is nil")
)

func NewAuth(region, userPoolId, cognitoClientId, cognitoClientSecret string) *Auth {
	sess := session.Must(session.NewSession())
	provider := cognitoidentityprovider.New(
		sess, aws.NewConfig().WithRegion(region),
	)

	authConfig := &authConfig{
		userPoolId:          userPoolId,
		cognitoClientId:     cognitoClientId,
		cognitoClientSecret: cognitoClientSecret,
	}

	return &Auth{
		provider:   provider,
		authConfig: authConfig,
	}
}

const MaxRetryCount = 3

func (a *Auth) SignUp(ctx context.Context, email, password string) (string, error) {
	var signUpWithRetry func(retryCount int) (string, error)
	signUpWithRetry = func(retryCount int) (string, error) {
		if retryCount > MaxRetryCount {
			return "", fmt.Errorf("maximum retry attempts exceeded")
		}

		cognitoUserName := generateCognitoUserName() // Cognito内でのみ使用するユーザー名
		secretHash := a.getSecretHash(cognitoUserName)

		param := &cognitoidentityprovider.SignUpInput{
			ClientId:   aws.String(a.cognitoClientId),
			SecretHash: aws.String(secretHash),
			Username:   aws.String(cognitoUserName),
			Password:   aws.String(password),
			UserAttributes: []*cognitoidentityprovider.AttributeType{
				{
					Name:  aws.String("email"),
					Value: aws.String(email),
				},
			},
		}

		res, err := a.provider.SignUpWithContext(ctx, param)
		if err != nil {
			if awserr, ok := err.(awserr.Error); ok {
				log.Println("awserr.Code(): " + awserr.Code())
				log.Println("awserr.Message(): " + awserr.Message())
				switch awserr.(type) {
				case *cognitoidentityprovider.InvalidPasswordException:
					return "", fmt.Errorf("%w: %w", ErrInvalidPassword, awserr)
				case *cognitoidentityprovider.AliasExistsException:
					return "", fmt.Errorf("%w: %w", ErrEmailAlreadyExists, awserr)
				case *cognitoidentityprovider.UsernameExistsException:
					// Cognitoのユーザー名が既に存在する場合は、Cognitoのユーザー名を変更して再度登録する
					return signUpWithRetry(retryCount + 1)
				default:
					return "", fmt.Errorf("error from aws: %w", awserr)
				}
			}
			return "", err
		}

		if res.UserSub == nil {
			return "", ErrUserSubIsNil
		}

		log.Print("usersub: " + *res.UserSub)
		// Cognitoから返されるUUIDを返す
		return *res.UserSub, nil
	}

	return signUpWithRetry(0)
}

func (a *Auth) Login(ctx context.Context, email, password string) (*Tokens, error) {
	params := &cognitoidentityprovider.AdminInitiateAuthInput{
		ClientId:   aws.String(a.cognitoClientId),
		UserPoolId: aws.String(a.userPoolId),
		AuthFlow:   aws.String(cognitoidentityprovider.AuthFlowTypeAdminNoSrpAuth),
		AuthParameters: map[string]*string{
			"EMAIL":    aws.String("email"),
			"PASSWORD": aws.String("password"),
		},
	}

	res, err := a.provider.AdminInitiateAuth(params)
	if err != nil {
		log.Print("AdminAuth Error")
		return nil, err
	}
	if res == nil || res.AuthenticationResult == nil || res.AuthenticationResult.IdToken == nil {
		return nil, errors.New("failed to login")
	}

	tokens := &Tokens{
		Id:      *res.AuthenticationResult.IdToken,
		Access:  *res.AuthenticationResult.AccessToken,
		Refresh: *res.AuthenticationResult.RefreshToken,
	}

	return tokens, nil
}

// Cognitoのユーザー名とクライアントID、クライアントシークレットからシークレットハッシュを生成する
func (a *Auth) getSecretHash(username string) string {
	mac := hmac.New(sha256.New, []byte(a.cognitoClientSecret))
	mac.Write([]byte(username + a.cognitoClientId))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

// Cognito内でのみ使用するユーザー名を生成する
func generateCognitoUserName() string {
	return uuid.New().String()
}
