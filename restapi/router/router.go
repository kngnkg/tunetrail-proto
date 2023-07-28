package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kngnkg/tunetrail/restapi/auth"
	"github.com/kngnkg/tunetrail/restapi/clock"
	"github.com/kngnkg/tunetrail/restapi/config"
	"github.com/kngnkg/tunetrail/restapi/handler"
	"github.com/kngnkg/tunetrail/restapi/service"
	"github.com/kngnkg/tunetrail/restapi/store"
)

func SetupRouter(cfg *config.Config) (*gin.Engine, func(), error) {
	db, cleanup, err := store.New(cfg)
	if err != nil {
		return nil, cleanup, err
	}

	a := auth.NewAuth(
		cfg.AWSRegion, cfg.CognitoUserPoolId, cfg.CognitoClientId, cfg.CognitoClientSecret,
	)

	cl := clock.RealClocker{}

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

	router.Use(CorsMiddleware())

	router.GET("/health", hh.HealthCheck)

	auth := router.Group("/auth")
	{
		auth.POST("/register", ah.RegisterUser)
		auth.PUT("/confirm", ah.ConfirmEmail)
		// auth.POST("/login", ah.Login)           // ログイン
		// auth.POST("/logout", ah.Logout)         // ログアウト
	}

	user := router.Group("/user")
	{
		user.GET("/:user_name", uh.GetUserByUserName) // ログインユーザ以外のユーザー情報取得
		// auth.GET("/me", uh.GetMe)                     // ログインユーザー情報取得
		user.PUT("/", uh.UpdateUser) // TODO: 改修予定
		// user.PUT("/:user_name", uh.UpdateUser)        // TODO: 改修予定
		user.DELETE("/:user_name", uh.DeleteUserByUserName)
	}

	return router, cleanup, nil
}

// CORSの設定
func CorsMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		// 許可したいHTTPメソッドの一覧
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		// 許可したいHTTPリクエストヘッダの一覧
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		// 許可したいアクセス元の一覧
		AllowOrigins: []string{
			"http://localhost:3000",
			"http://www.tune-trail.com",
			"https://www.tune-trail.com",
			"http://tune-trail.com",
			"https://tune-trail.com",
		},
		// preflight requestで許可した後の接続可能時間
		MaxAge: 24 * time.Hour,
	})
}
