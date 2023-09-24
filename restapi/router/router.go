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
	lh := &handler.LikeHandler{
		Service: &service.LikeService{DB: db, Repo: r},
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

	users := router.Group("/users")
	{
		users.Use(handler.AuthMiddleware(j))

		users.GET("/by/username/:user_name", uh.GetUserByUserName)
		users.GET("/me", uh.GetMe) // ログインユーザー情報取得
		users.GET("/timelines", ph.GetTimeline)

		id := users.Group("/:user_id")
		{
			// id.GET("", uh.GetUserByUserId) // TODO: 改修予定
			// id.PUT("", uh.UpdateUser) // TODO: 改修予定
			// id.DELETE("", uh.DeleteUserByUserName) // TODO: 改修予定
			id.GET("/posts", ph.GetPostsByUserId)
			id.GET("/likes", ph.GetLikedPostsByUserId)
			id.GET("/followees", uh.GetFollowees)
			id.GET("/followers", uh.GetFollowers)

			follow := id.Group("/follows")
			{
				follow.POST("", uh.FollowUser)
				follow.DELETE("/:followee_user_id", uh.UnfollowUser)
			}
		}
	}

	posts := router.Group("/posts")
	{
		posts.Use(handler.AuthMiddleware(j))

		posts.POST("", ph.AddPost)

		postId := posts.Group("/:post_id")
		{
			postId.GET("", ph.GetPostById)
			postId.DELETE("", ph.DeletePost)
			postId.GET("/replies", ph.GetReplies)

			likes := postId.Group("/likes")
			{
				likes.POST("", lh.AddLike)
				likes.DELETE("", lh.DeleteLike)
			}
		}
	}

	return router, cleanup, nil
}
