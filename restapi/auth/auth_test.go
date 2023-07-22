package auth

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/stretchr/testify/assert"
)

func createAuthFortest(t *testing.T, apm *AuthProviderMock) *Auth {
	return &Auth{
		provider: apm,
		authConfig: &authConfig{
			userPoolId:      "test-userPoolId",
			cognitoClientId: "test-cognitoClientId",
		},
	}
}

func TestAuth_SignUp(t *testing.T) {
	var (
		VALID_USER_SUB  = "test-userSub"
		DUPLICATE_EMAIL = "duplicate@example.com"
		VALID_PASSWORD  = "password"
	)

	apm := &AuthProviderMock{}
	// モックの設定
	apm.SignUpWithContextFunc = func(ctx context.Context, input *cognitoidentityprovider.SignUpInput, opts ...request.Option) (*cognitoidentityprovider.SignUpOutput, error) {
		if *input.Username == DUPLICATE_EMAIL {
			awserr := &cognitoidentityprovider.UsernameExistsException{
				Message_: aws.String("mock"),
			}
			return nil, awserr
		}
		if *input.Password != VALID_PASSWORD {
			awserr := &cognitoidentityprovider.InvalidPasswordException{
				Message_: aws.String("mock"),
			}
			return nil, awserr
		}
		output := &cognitoidentityprovider.SignUpOutput{
			UserSub:       aws.String(VALID_USER_SUB),
			UserConfirmed: aws.Bool(true),
		}
		return output, nil
	}

	type args struct {
		ctx      context.Context
		email    string
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			name: "success",
			args: args{
				ctx:      context.Background(),
				email:    "email@example.com",
				password: VALID_PASSWORD,
			},
			want:    VALID_USER_SUB,
			wantErr: nil,
		},
		{
			name: "email already exists",
			args: args{
				ctx:      context.Background(),
				email:    DUPLICATE_EMAIL,
				password: VALID_PASSWORD,
			},
			want:    "",
			wantErr: ErrEmailAlreadyExists,
		},
		{
			name: "password is invalid",
			args: args{
				ctx:      context.Background(),
				email:    "email@example.com",
				password: "invalid",
			},
			want:    "",
			wantErr: ErrInvalidPassword,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := createAuthFortest(t, apm)

			got, err := a.SignUp(tt.args.ctx, tt.args.email, tt.args.password)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Auth.SignUp() error = %v, wantErr %v", err, tt.wantErr)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
