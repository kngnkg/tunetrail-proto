package service

// var (
// 	invalidPassword        = "invalidPassword"
// 	mismatchCode           = "mismatch"
// 	expiredCode            = "expired"
// 	emailAlreadyExistsCode = "emailAlreadyExists"
// )

// type AuthServiceTestSuite struct {
// 	suite.Suite
// 	as         *AuthService        // テスト対象のサービス
// 	fc         *clock.FixedClocker // テスト用の時刻を固定する
// 	dummyUsers []*model.User       // テスト用のダミーユーザー
// }

// func TestAuthServiceTestSuite(t *testing.T) {
// 	t.Parallel()
// 	suite.Run(t, new(AuthServiceTestSuite))
// }

// func (s *AuthServiceTestSuite) SetupTest() {
// 	fc := &clock.FixedClocker{}

// 	// テスト用のダミーユーザーを作成
// 	// これらのユーザーだけがDBに存在することを想定している
// 	dummyUsers := fixture.CreateUsers(2)

// 	moqRepo := s.setupRepoMock()
// 	moqAuth := s.setupAuthMock()

// 	s.as = &AuthService{
// 		DB:   &DBConnectionMock{},
// 		Repo: moqRepo,
// 		Auth: moqAuth,
// 	}
// 	s.fc = fc
// 	s.dummyUsers = dummyUsers
// }

// func (s *AuthServiceTestSuite) TestRegisterUser() {
// 	type args struct {
// 		ctx  context.Context
// 		data *model.UserRegistrationData
// 	}

// 	tests := []struct {
// 		name     string
// 		args     args
// 		wantUser *model.User
// 		wantErr  error
// 	}{
// 		{
// 			"ok",
// 			args{
// 				ctx: context.Background(),
// 				data: &model.UserRegistrationData{
// 					UserName: "test",
// 					Name:     "test",
// 					Password: "test",
// 					Email:    "test@example.com",
// 				},
// 			},
// 			&model.User{
// 				UserName:  "test",
// 				Name:      "test",
// 				IconUrl:   "",
// 				Bio:       "",
// 				CreatedAt: s.fc.Now(),
// 				UpdatedAt: s.fc.Now(),
// 			},
// 			nil,
// 		},
// 		{
// 			// ユーザー名が重複している場合
// 			"errUserNameAlreadyExists",
// 			args{
// 				ctx: context.Background(),
// 				data: &model.UserRegistrationData{
// 					UserName: s.dummyUsers[0].UserName, // ダミーユーザーのユーザー名と重複させる
// 					Name:     "test",
// 					Password: "test",
// 					Email:    "test@example.com",
// 				},
// 			},
// 			nil,
// 			ErrUserNameAlreadyExists,
// 		},
// 		{
// 			// メールアドレスが重複している場合
// 			"errEmailAlreadyExists",
// 			args{
// 				ctx: context.Background(),
// 				data: &model.UserRegistrationData{
// 					UserName: "test",
// 					Name:     "test",
// 					Password: "test",
// 					Email:    s.dummyUsers[0].Email, // ダミーユーザーのメールアドレスと重複させる
// 				},
// 			},
// 			nil,
// 			ErrEmailAlreadyExists,
// 		},
// 		// パスワードが不正な場合
// 	}

// 	for _, tt := range tests {
// 		s.Run(tt.name, func() {
// 			got, err := s.as.RegisterUser(tt.args.ctx, tt.args.data)
// 			if !errors.Is(err, tt.wantErr) {
// 				s.T().Errorf("UserService.RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
// 			}

// 			if tt.wantErr != nil {
// 				// 異常系のテストの場合はここで終了
// 				return
// 			}

// 			// idが設定されていることを確認
// 			assert.NotEmpty(s.T(), got.Id)

// 			// それ以外のフィールドは正しい値が設定されていることを確認
// 			wantv := reflect.ValueOf(tt.wantUser).Elem()
// 			gotv := reflect.ValueOf(got).Elem()
// 			for i := 0; i < wantv.NumField(); i++ {
// 				if field := wantv.Type().Field(i); field.Name == "Id" {
// 					continue
// 				}
// 				assert.Equal(s.T(), wantv.Field(i).Interface(), gotv.Field(i).Interface())
// 			}
// 		})
// 	}
// }

// func (s *AuthServiceTestSuite) TestConfirmEmail() {
// 	type args struct {
// 		ctx      context.Context
// 		userName string
// 		code     string
// 	}

// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr error
// 	}{
// 		{
// 			"success",
// 			args{
// 				ctx:      context.Background(),
// 				userName: s.dummyUsers[0].UserName,
// 				code:     "test",
// 			},
// 			nil,
// 		},
// 		{
// 			"user not found",
// 			args{
// 				ctx:      context.Background(),
// 				userName: "not_found",
// 				code:     "test",
// 			},
// 			ErrUserNotFound,
// 		},
// 		{
// 			"code mismatch",
// 			args{
// 				ctx:      context.Background(),
// 				userName: s.dummyUsers[0].UserName,
// 				code:     mismatchCode,
// 			},
// 			auth.ErrCodeMismatch,
// 		},
// 		{
// 			"code expired",
// 			args{
// 				ctx:      context.Background(),
// 				userName: s.dummyUsers[0].UserName,
// 				code:     expiredCode,
// 			},
// 			auth.ErrCodeExpired,
// 		},
// 		{
// 			"email already exists",
// 			args{
// 				ctx:      context.Background(),
// 				userName: s.dummyUsers[0].UserName,
// 				code:     emailAlreadyExistsCode,
// 			},
// 			auth.ErrEmailAlreadyExists,
// 		},
// 	}

// 	for _, tt := range tests {
// 		s.Run(tt.name, func() {
// 			err := s.as.ConfirmEmail(tt.args.ctx, tt.args.userName, tt.args.code)
// 			if !errors.Is(err, tt.wantErr) {
// 				s.T().Errorf("UserService.ConfirmEmail() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func (s *AuthServiceTestSuite) TestSignIn() {
// 	type args struct {
// 		ctx  context.Context
// 		data *model.UserSignInData
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    bool
// 		wantErr error
// 	}{
// 		{
// 			"success (userName)",
// 			args{
// 				ctx: context.Background(),
// 				data: &model.UserSignInData{
// 					UserName: s.dummyUsers[0].UserName,
// 					Password: s.dummyUsers[0].Password,
// 				},
// 			},
// 			true,
// 			nil,
// 		},
// 		{
// 			"success (email)",
// 			args{
// 				ctx: context.Background(),
// 				data: &model.UserSignInData{
// 					Email:    s.dummyUsers[0].Email,
// 					Password: s.dummyUsers[0].Password,
// 				},
// 			},
// 			true,
// 			nil,
// 		},
// 		{
// 			"invalid password",
// 			args{
// 				ctx: context.Background(),
// 				data: &model.UserSignInData{
// 					Email:    s.dummyUsers[0].Email,
// 					Password: "invalid",
// 				},
// 			},
// 			false,
// 			ErrNotAuthorized,
// 		},
// 		{
// 			"user not found (userName)",
// 			args{
// 				ctx: context.Background(),
// 				data: &model.UserSignInData{
// 					UserName: "notFound",
// 					Password: s.dummyUsers[0].Password,
// 				},
// 			},
// 			false,
// 			ErrUserNotFound,
// 		},
// 		{
// 			"user not found (email)",
// 			args{
// 				ctx: context.Background(),
// 				data: &model.UserSignInData{
// 					Email:    "notFound@example.com",
// 					Password: s.dummyUsers[0].Password,
// 				},
// 			},
// 			false,
// 			ErrUserNotFound,
// 		},
// 	}
// 	for _, tt := range tests {
// 		s.Run(tt.name, func() {
// 			got, err := s.as.SignIn(tt.args.ctx, tt.args.data)
// 			if !errors.Is(err, tt.wantErr) {
// 				s.T().Errorf("UserService.SignIn() error = %v, wantErr %v", err, tt.wantErr)
// 			}

// 			assert.Equal(s.T(), tt.want, got != nil)
// 		})
// 	}
// }

// func (s *AuthServiceTestSuite) setupRepoMock() *UserRepositoryMock {
// 	mock := &UserRepositoryMock{
// 		WithTransactionFunc: func(ctx context.Context, db store.Beginner, fn func(tx *sqlx.Tx) error) error {
// 			tx := &sqlx.Tx{}
// 			err := fn(tx)
// 			if err != nil {
// 				return fmt.Errorf("failed to execute mock transaction: %w", err)
// 			}
// 			return nil
// 		},
// 		GetUserByUserNameFunc: func(ctx context.Context, db store.Queryer, userName string) (*model.User, error) {
// 			for _, u := range s.dummyUsers {
// 				if userName == u.UserName {
// 					return u, nil
// 				}
// 			}
// 			return nil, store.ErrUserNotFound
// 		},
// 		UserExistsByUserNameFunc: func(ctx context.Context, db store.Queryer, userName string) (bool, error) {
// 			for _, u := range s.dummyUsers {
// 				if userName == u.UserName {
// 					return true, nil
// 				}
// 			}
// 			return false, nil
// 		},
// 		RegisterUserFunc: func(ctx context.Context, db store.Execer, u *model.User) error {
// 			// ダミーの値を設定
// 			u.CreatedAt = s.fc.Now()
// 			u.UpdatedAt = s.fc.Now()
// 			return nil
// 		},
// 	}
// 	return mock
// }

// func (s *AuthServiceTestSuite) setupAuthMock() *AuthMock {
// 	mock := &AuthMock{
// 		SignUpFunc: func(ctx context.Context, userId model.UserID, email, password string) error {
// 			if password == invalidPassword {
// 				return auth.ErrInvalidPassword
// 			}
// 			for _, u := range s.dummyUsers {
// 				if userId == u.Id {
// 					return auth.ErrUserIdAlreadyExists
// 				}
// 				if email == u.Email {
// 					return auth.ErrEmailAlreadyExists
// 				}
// 			}
// 			return nil
// 		},
// 		ConfirmSignUpFunc: func(ctx context.Context, userId model.UserID, code string) error {
// 			if code == mismatchCode {
// 				return auth.ErrCodeMismatch
// 			}
// 			if code == expiredCode {
// 				return auth.ErrCodeExpired
// 			}
// 			if code == emailAlreadyExistsCode {
// 				return auth.ErrEmailAlreadyExists
// 			}
// 			return nil
// 		},
// 		SignInFunc: func(ctx context.Context, userIdentifier, password string) (*model.Tokens, error) {
// 			for _, u := range s.dummyUsers {
// 				if userIdentifier == string(u.Id) || userIdentifier == u.Email {
// 					if password != u.Password {
// 						return nil, auth.ErrNotAuthorized
// 					}
// 					return &model.Tokens{
// 						Id:      uuid.New().String(),
// 						Access:  uuid.New().String(),
// 						Refresh: uuid.New().String(),
// 					}, nil
// 				}
// 			}
// 			return nil, auth.ErrUserNotFound
// 		},
// 	}
// 	return mock
// }
