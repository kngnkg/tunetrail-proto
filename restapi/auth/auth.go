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
	"github.com/kngnkg/tunetrail/restapi/model"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . AuthProvider
type AuthProvider interface {
	SignUpWithContext(ctx context.Context, input *cognitoidentityprovider.SignUpInput, opts ...request.Option) (*cognitoidentityprovider.SignUpOutput, error)
	ConfirmSignUpWithContext(ctx context.Context, input *cognitoidentityprovider.ConfirmSignUpInput, opts ...request.Option) (*cognitoidentityprovider.ConfirmSignUpOutput, error)
	AdminInitiateAuthWithContext(ctx context.Context, input *cognitoidentityprovider.AdminInitiateAuthInput, opts ...request.Option) (*cognitoidentityprovider.AdminInitiateAuthOutput, error)
}

// Cognitoから返されるトークン

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
	ErrUserIdAlreadyExists = errors.New("auth: userId already exists")
	ErrEmailAlreadyExists  = errors.New("auth: email already exists")
	ErrInvalidPassword     = errors.New("auth: invalid password")
	ErrUserNotFound        = errors.New("auth: user not found")
	ErrNotAuthorized       = errors.New("auth: invalid email or password")
	ErrUserNotConfirmed    = errors.New("auth: user not confirmed")
	ErrUserSubIsNil        = errors.New("auth: user sub is nil")
	ErrCodeMismatch        = errors.New("auth: code mismatch")
	ErrCodeExpired         = errors.New("auth: code expired")
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

func (a *Auth) SignUp(ctx context.Context, userId model.UserID, email, password string) error {
	secretHash := a.getSecretHash(string(userId))

	param := &cognitoidentityprovider.SignUpInput{
		ClientId:   aws.String(a.cognitoClientId),
		SecretHash: aws.String(secretHash),
		Username:   aws.String(string(userId)),
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
			case *cognitoidentityprovider.UsernameExistsException:
				return fmt.Errorf("%w: %w", ErrUserIdAlreadyExists, awserr)
			case *cognitoidentityprovider.InvalidPasswordException:
				return fmt.Errorf("%w: %w", ErrInvalidPassword, awserr)
			default:
				return fmt.Errorf("error from aws: %w", awserr)
			}
		}
		return err
	}

	if res.UserSub == nil {
		return ErrUserSubIsNil
	}

	return nil
}

func (a *Auth) ConfirmSignUp(ctx context.Context, userId model.UserID, code string) error {
	secretHash := a.getSecretHash(string(userId))

	param := &cognitoidentityprovider.ConfirmSignUpInput{
		ClientId:         aws.String(a.cognitoClientId),
		SecretHash:       aws.String(secretHash),
		Username:         aws.String(string(userId)),
		ConfirmationCode: aws.String(code),
	}

	_, err := a.provider.ConfirmSignUpWithContext(ctx, param)
	if err != nil {
		if awserr, ok := err.(awserr.Error); ok {
			log.Println("awserr.Code(): " + awserr.Code())
			log.Println("awserr.Message(): " + awserr.Message())
			switch awserr.(type) {
			case *cognitoidentityprovider.AliasExistsException:
				return fmt.Errorf("%w: %w", ErrEmailAlreadyExists, awserr)
			case *cognitoidentityprovider.CodeMismatchException:
				return fmt.Errorf("%w: %w", ErrCodeMismatch, awserr)
			case *cognitoidentityprovider.ExpiredCodeException:
				return fmt.Errorf("%w: %w", ErrCodeExpired, awserr)
			default:
				return fmt.Errorf("error from aws: %w", awserr)
			}
		}
		return err
	}

	return nil
}

func (a *Auth) SignIn(ctx context.Context, userIdentifier, password string) (*model.Tokens, error) {
	secretHash := a.getSecretHash(userIdentifier)

	params := &cognitoidentityprovider.AdminInitiateAuthInput{
		ClientId:   aws.String(a.cognitoClientId),
		UserPoolId: aws.String(a.userPoolId),
		AuthFlow:   aws.String(cognitoidentityprovider.AuthFlowTypeAdminNoSrpAuth),
		AuthParameters: map[string]*string{
			"USERNAME":    aws.String(userIdentifier),
			"PASSWORD":    aws.String(password),
			"SECRET_HASH": aws.String(secretHash),
		},
	}

	res, err := a.provider.AdminInitiateAuthWithContext(ctx, params)
	if err != nil {
		if awserr, ok := err.(awserr.Error); ok {
			log.Println("awserr.Code(): " + awserr.Code())
			log.Println("awserr.Message(): " + awserr.Message())
			switch awserr.(type) {
			case *cognitoidentityprovider.UserNotFoundException:
				return nil, fmt.Errorf("%w: %w", ErrUserNotFound, awserr)
			case *cognitoidentityprovider.NotAuthorizedException:
				return nil, fmt.Errorf("%w: %w", ErrNotAuthorized, awserr)
			default:
				return nil, fmt.Errorf("error from aws: %w", awserr)
			}
		}
		return nil, err
	}

	if res == nil || res.AuthenticationResult == nil || res.AuthenticationResult.IdToken == nil {
		return nil, errors.New("failed to login")
	}

	tokens := &model.Tokens{
		Id:      *res.AuthenticationResult.IdToken,
		Access:  *res.AuthenticationResult.AccessToken,
		Refresh: *res.AuthenticationResult.RefreshToken,
	}

	return tokens, nil
}

func (a *Auth) RefreshToken(ctx context.Context, userIdentifier, refreshToken string) (*model.Tokens, error) {
	secretHash := a.getSecretHash(userIdentifier)

	params := &cognitoidentityprovider.AdminInitiateAuthInput{
		ClientId:   aws.String(a.cognitoClientId),
		UserPoolId: aws.String(a.userPoolId),
		AuthFlow:   aws.String(cognitoidentityprovider.AuthFlowTypeRefreshToken),
		AuthParameters: map[string]*string{
			"REFRESH_TOKEN": aws.String(refreshToken),
			"SECRET_HASH":   aws.String(secretHash),
		},
	}

	res, err := a.provider.AdminInitiateAuthWithContext(ctx, params)
	if err != nil {
		if awserr, ok := err.(awserr.Error); ok {
			log.Println("awserr.Code(): " + awserr.Code())
			log.Println("awserr.Message(): " + awserr.Message())
			switch awserr.(type) {
			case *cognitoidentityprovider.NotAuthorizedException:
				return nil, fmt.Errorf("%w: %w", ErrNotAuthorized, awserr)
			default:
				return nil, fmt.Errorf("error from aws: %w", awserr)
			}
		}
		return nil, err
	}

	if res == nil || res.AuthenticationResult == nil || res.AuthenticationResult.IdToken == nil {
		return nil, errors.New("failed to login")
	}

	tokens := &model.Tokens{
		Id:     *res.AuthenticationResult.IdToken,
		Access: *res.AuthenticationResult.AccessToken,
	}

	return tokens, nil
}

// Cognitoのユーザー名とクライアントID、クライアントシークレットからシークレットハッシュを生成する
func (a *Auth) getSecretHash(userIdentifier string) string {
	mac := hmac.New(sha256.New, []byte(a.cognitoClientSecret))
	mac.Write([]byte(userIdentifier + a.cognitoClientId))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
