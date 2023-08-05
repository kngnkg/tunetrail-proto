package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/kngnkg/tunetrail/restapi/clock"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
)

type JWTerConfig struct {
	Region          string
	UserPoolId      string
	CognitoClientId string
}

type JWTer struct {
	clocker clock.Clocker
	*JWTerConfig
}

var (
	ErrTokenExpired = errors.New("token is expired")
)

func NewJWTer(clocker clock.Clocker, config *JWTerConfig) *JWTer {
	return &JWTer{
		clocker:     clocker,
		JWTerConfig: config,
	}
}

func (j *JWTer) Verify(ctx context.Context, tokenString string) error {
	// キャッシュされるか確認
	url := fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json", j.Region, j.UserPoolId)
	keySet, err := jwk.Fetch(ctx, url)
	if err != nil {
		return err
	}

	token, err := jwt.Parse([]byte(tokenString), jwt.WithKeySet(keySet))
	if err != nil {
		return err
	}

	// tokenの検証

	// iss (issuer) クレームは、ユーザープールのURLである必要がある
	if token.Issuer() != fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s", j.Region, j.UserPoolId) {
		return errors.New("invalid issuer")
	}

	clientId, _ := token.Get("client_id")
	if clientId.(string) != j.CognitoClientId {
		return errors.New("invalid client_id")
	}

	tokenUse, _ := token.Get("token_use")
	if tokenUse.(string) != "access" {
		return errors.New("invalid token_use")
	}

	// scope クレームに必要なスコープが含まれていることを確認する
	scope, _ := token.Get("scope")
	if scope.(string) != "aws.cognito.signin.user.admin" {
		return errors.New("invalid scope")
	}

	// exp (expiration time) クレームは、現在の時刻よりも未来である必要がある
	if !token.Expiration().After(j.clocker.Now()) {
		return ErrTokenExpired
	}

	return nil
}
