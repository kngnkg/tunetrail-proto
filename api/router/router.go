package router

import (
	"log"
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
	log.Print("router: start router.setupRouter")
	log.Print("router: setup gin router")
	router := gin.Default()

	log.Print("router: setup db")
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

	log.Print("router: setup cors middleware")
	router.Use(CorsMiddleware())

	log.Print("router: setup endpoints")
	router.GET("/health", hh.HealthCheck)
	user := router.Group("/user")
	{
		user.POST("/register", uh.RegisterUser)
		user.GET("/:user_name", uh.GetUserByUserName)
		user.PUT("/update", uh.UpdateUser)
		user.DELETE("/:user_name", uh.DeleteUserByUserName)
	}

	log.Print("router: end router.setupRouter")
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
