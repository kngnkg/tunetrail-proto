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
			DB:    db,
			Repo:  r,
			Auth:  a,
			JWTer: j,
		},
		AllowedDomain: cfg.AllowedDomain,
	}
	ph := &handler.PostHandler{
		Service: &service.PostService{DB: db, Repo: r},
	}

	router := gin.Default()

	router.Use(handler.CorsMiddleware(cfg.AllowedDomain))

	router.GET("/health", hh.HealthCheck)

	auth := router.Group("/auth")
	{
		auth.POST("/register", ah.RegisterUser)
		auth.PUT("/confirm", ah.ConfirmEmail)
		auth.POST("/signin", ah.SignIn)
		auth.POST("/refresh", ah.RefreshToken)
		// auth.POST("/signout", handler.AuthMiddleware(j), ah.SignOut) // サインアウト
	}

	// TODO: user_idで指定する
	users := router.Group("/users")
	{
		users.Use(handler.AuthMiddleware(j))

		// TODO: /users/by/username/:user_name に変更する
		users.GET("/:user_name", uh.GetUserByUserName) // ログインユーザ以外のユーザー情報取得
		users.GET("/me", uh.GetMe)                     // ログインユーザー情報取得
		users.PUT("", uh.UpdateUser)                   // TODO: 改修予定
		// user.PUT("/:user_name", uh.UpdateUser)        // TODO: 改修予定
		users.DELETE("/:user_name", uh.DeleteUserByUserName)

		users.GET("/timelines", ph.GetTimeline)

		follow := users.Group("/:user_name/follow")
		{
			follow.POST("", uh.FollowUser)
			follow.DELETE("", uh.UnfollowUser)
		}

		// userPosts := users.Group("/:user_id/posts")
		// {
		// 	userPosts.GET("", ph.GetPostsByUserId)
		// }

		// posts := users.Group("/:user_name/posts")
		// {
		// 	posts.GET("", ph.GetPostsByUserName)
		// }
	}

	posts := router.Group("/posts")
	{
		posts.Use(handler.AuthMiddleware(j))

		posts.POST("", ph.AddPost)

		// 暫定的にここに置く
		posts.GET("/:user_id", ph.GetPostsByUserId)
	}

	return router, cleanup, nil
}
