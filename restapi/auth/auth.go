package auth

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
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
	userPoolId      string
	cognitoClientId string
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

func NewAuth(region, userPoolId, cognitoClientId string) *Auth {
	sess := session.Must(session.NewSession())
	provider := cognitoidentityprovider.New(
		sess, aws.NewConfig().WithRegion(region),
	)

	authConfig := &authConfig{
		userPoolId:      userPoolId,
		cognitoClientId: cognitoClientId,
	}

	return &Auth{
		provider:   provider,
		authConfig: authConfig,
	}
}

func (a *Auth) SignUp(ctx context.Context, email, password string) (string, error) {
	param := &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(a.cognitoClientId),
		Username: aws.String(email), // emailをusernameとして登録する
		Password: aws.String(password),
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
			case *cognitoidentityprovider.UsernameExistsException:
				return "", fmt.Errorf("%w: %w", ErrEmailAlreadyExists, awserr)
			case *cognitoidentityprovider.InvalidPasswordException:
				return "", fmt.Errorf("%w: %w", ErrInvalidPassword, awserr)
			default:
				return "", fmt.Errorf("error from aws: %w", awserr)
			}
		}
		return "", err
	}

	if res.UserConfirmed == nil || !*res.UserConfirmed {
		return "", ErrUserNotConfirmed
	}
	if res.UserSub == nil {
		return "", ErrUserSubIsNil
	}

	// Cognitoから返されるUUIDを返す
	return *res.UserSub, nil
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
