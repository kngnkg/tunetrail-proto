package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kwtryo/tunetrail/api/clock"
	"github.com/kwtryo/tunetrail/api/config"
	"github.com/kwtryo/tunetrail/api/handler"
	"github.com/kwtryo/tunetrail/api/service"
	"github.com/kwtryo/tunetrail/api/store"
)

func SetupRouter(cfg *config.Config) (*gin.Engine, func(), error) {
	router := gin.Default()

	db, cleanup, err := store.New(cfg)
	if err != nil {
		return nil, cleanup, err
	}

	cl := clock.RealClocker{}
	r := &store.Repository{Clocker: cl}

	hh := &handler.HealthHandler{
		Service: &service.HealthService{DB: db, Repo: r},
	}
	uh := &handler.UserHandler{
		Service: &service.UserService{DB: db, Repo: r},
	}

	router.Use(CorsMiddleware())

	router.GET("/health", hh.HealthCheck)
	user := router.Group("/user")
	{
		user.POST("/register", uh.RegisterUser)
		user.GET("/:user_name", uh.GetUserByUserName)
		user.PUT("/update", uh.UpdateUser)
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
		},
		// preflight requestで許可した後の接続可能時間
		MaxAge: 24 * time.Hour,
	})
}
