package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kngnkg/tunetrail/restapi/auth"
	"github.com/kngnkg/tunetrail/restapi/clock"
	"github.com/kngnkg/tunetrail/restapi/config"
	"github.com/kngnkg/tunetrail/restapi/handler"
	"github.com/kngnkg/tunetrail/restapi/service"
	"github.com/kngnkg/tunetrail/restapi/store"
)

func SetupRouter(cfg *config.Config) (*gin.Engine, func(), error) {
	cl := clock.RealClocker{}

	db, cleanup, err := store.New(cfg)
	if err != nil {
		return nil, cleanup, err
	}

	a := auth.NewAuth(
		cfg.AWSRegion, cfg.CognitoUserPoolId, cfg.CognitoClientId, cfg.CognitoClientSecret,
	)

	j := auth.NewJWTer(cl, &auth.JWTerConfig{
		Region:          cfg.AWSRegion,
		UserPoolId:      cfg.CognitoUserPoolId,
		CognitoClientId: cfg.CognitoClientId,
	})

	r := &store.Repository{Clocker: cl}

	hh := &handler.HealthHandler{
		Service: &service.HealthService{DB: db, Repo: r},
	}
	uh := &handler.UserHandler{
		Service: &service.UserService{DB: db, Repo: r},
	}
	ah := &handler.AuthHandler{
		Service: &service.AuthService{
			DB:   db,
			Repo: r,
			Auth: a,
		},
	}

	router := gin.Default()

	router.Use(handler.CorsMiddleware())

	router.GET("/health", hh.HealthCheck)

	auth := router.Group("/auth")
	{
		auth.POST("/register", ah.RegisterUser)
		auth.PUT("/confirm", ah.ConfirmEmail)
		auth.POST("/signin", ah.SignIn)
		auth.POST("/refresh", ah.RefreshToken)
		// auth.POST("/signout", handler.AuthMiddleware(j), ah.SignOut) // サインアウト
	}

	user := router.Group("/user")
	{
		user.Use(handler.AuthMiddleware(j))
		user.GET("/:user_name", uh.GetUserByUserName) // ログインユーザ以外のユーザー情報取得
		// auth.GET("/me", uh.GetMe)                     // ログインユーザー情報取得
		user.PUT("/", uh.UpdateUser) // TODO: 改修予定
		// user.PUT("/:user_name", uh.UpdateUser)        // TODO: 改修予定
		user.DELETE("/:user_name", uh.DeleteUserByUserName)
	}

	return router, cleanup, nil
}
