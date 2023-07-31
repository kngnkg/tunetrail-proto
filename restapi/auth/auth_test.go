package auth

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/kngnkg/tunetrail/restapi/model"
	"github.com/kngnkg/tunetrail/restapi/testutil/fixture"
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
		VALID_PASSWORD = "password"
	)

	type args struct {
		ctx      context.Context
		email    string
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr error
	}{
		{
			name: "success",
			args: args{
				ctx:      context.Background(),
				email:    "email@example.com",
				password: VALID_PASSWORD,
			},
			want:    true,
			wantErr: nil,
		},
		{
			name: "password is invalid",
			args: args{
				ctx:      context.Background(),
				email:    "email@example.com",
				password: "invalid",
			},
			want:    false,
			wantErr: ErrInvalidPassword,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apm := &AuthProviderMock{}
			// モックの設定
			apm.SignUpWithContextFunc = func(ctx context.Context, input *cognitoidentityprovider.SignUpInput, opts ...request.Option) (*cognitoidentityprovider.SignUpOutput, error) {
				if *input.Password != VALID_PASSWORD {
					awserr := &cognitoidentityprovider.InvalidPasswordException{
						Message_: aws.String("mock"),
					}
					return nil, awserr
				}
				output := &cognitoidentityprovider.SignUpOutput{
					UserSub: aws.String("test-userSub"),
				}
				return output, nil
			}

			a := createAuthFortest(t, apm)

			got, err := a.SignUp(tt.args.ctx, tt.args.email, tt.args.password)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Auth.SignUp() error = %v, wantErr %v", err, tt.wantErr)
			}

			assert.Equal(t, tt.want, got != "")
		})
	}
}

func TestAuth_ConfirmSignUp(t *testing.T) {
	var (
		MISMATCH_CODE = "000000"
		EXPIRED_CODE  = "111111"
		EMAIL_EXISTS  = "email already exists"
	)

	type args struct {
		ctx             context.Context
		cognitoUserName string
		code            string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "success",
			args: args{
				ctx:             context.Background(),
				cognitoUserName: "test-userName",
				code:            "123456",
			},
			wantErr: nil,
		},
		{
			name: "code mismatch",
			args: args{
				ctx:             context.Background(),
				cognitoUserName: "test-userName",
				code:            MISMATCH_CODE,
			},
			wantErr: ErrCodeMismatch,
		},
		{
			name: "code is expired",
			args: args{
				ctx:             context.Background(),
				cognitoUserName: "test-userName",
				code:            EXPIRED_CODE,
			},
			wantErr: ErrCodeExpired,
		},
		{
			name: "email already exists",
			args: args{
				ctx:             context.Background(),
				cognitoUserName: "test-userName",
				code:            EMAIL_EXISTS,
			},
			wantErr: ErrEmailAlreadyExists,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apm := &AuthProviderMock{}
			apm.ConfirmSignUpWithContextFunc = func(ctx context.Context, input *cognitoidentityprovider.ConfirmSignUpInput, opts ...request.Option) (*cognitoidentityprovider.ConfirmSignUpOutput, error) {
				if *input.ConfirmationCode == MISMATCH_CODE {
					awserr := &cognitoidentityprovider.CodeMismatchException{
						Message_: aws.String("mock"),
					}
					return nil, awserr
				}
				if *input.ConfirmationCode == EXPIRED_CODE {
					awserr := &cognitoidentityprovider.ExpiredCodeException{
						Message_: aws.String("mock"),
					}
					return nil, awserr
				}
				if *input.ConfirmationCode == EMAIL_EXISTS {
					awserr := &cognitoidentityprovider.AliasExistsException{
						Message_: aws.String("mock"),
					}
					return nil, awserr
				}
				output := &cognitoidentityprovider.ConfirmSignUpOutput{}
				return output, nil
			}

			a := createAuthFortest(t, apm)

			err := a.ConfirmSignUp(tt.args.ctx, tt.args.cognitoUserName, tt.args.code)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Auth.ConfirmSignUp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAuth_SignIn(t *testing.T) {
	// このユーザーが登録されていることを前提とする
	u := fixture.User(&model.User{})

	type args struct {
		ctx      context.Context
		email    string
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    *Tokens
		wantErr error
	}{
		{
			name: "success",
			args: args{
				ctx:      context.Background(),
				email:    u.Email,
				password: u.Password,
			},
			want:    &Tokens{},
			wantErr: nil,
		},
		{
			name: "password mismatch",
			args: args{
				ctx:      context.Background(),
				email:    u.Email,
				password: "invalid",
			},
			want:    nil,
			wantErr: ErrUserNotFound,
		},
		{
			name: "email not found",
			args: args{
				ctx:      context.Background(),
				email:    "invalid@example.com",
				password: u.Password,
			},
			want:    nil,
			wantErr: ErrUserNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apm := &AuthProviderMock{}
			// モックの設定
			apm.AdminInitiateAuthWithContextFunc = func(ctx context.Context, input *cognitoidentityprovider.AdminInitiateAuthInput, opts ...request.Option) (*cognitoidentityprovider.AdminInitiateAuthOutput, error) {
				if *input.AuthParameters["PASSWORD"] != u.Password {
					awserr := &cognitoidentityprovider.NotAuthorizedException{
						Message_: aws.String("mock"),
					}
					return nil, awserr
				}
				if *input.AuthParameters["EMAIL"] != u.Email {
					awserr := &cognitoidentityprovider.NotAuthorizedException{
						Message_: aws.String("mock"),
					}
					return nil, awserr
				}
				output := &cognitoidentityprovider.AdminInitiateAuthOutput{}
				return output, nil
			}

			a := createAuthFortest(t, apm)

			got, err := a.SignIn(tt.args.ctx, tt.args.email, tt.args.password)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Auth.SignIn() error = %v, wantErr %v", err, tt.wantErr)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
