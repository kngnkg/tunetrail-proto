package handler

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kngnkg/tunetrail/restapi/auth"
)

const (
	AccessTokenKey  = "accessToken"
	IdTokenKey      = "idToken"
	RefreshTokenKey = "refreshToken"
	UserIdKey       = "userId"
)

// CORSの設定
func CorsMiddleware(allowedDomain string) gin.HandlerFunc {
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
		},
		// 許可したいアクセス元の一覧
		AllowOrigins: []string{
			"https://" + allowedDomain + ":3000",
			"https://www." + allowedDomain + ":3000",
		},
		AllowCredentials: true,
		// preflight requestで許可した後の接続可能時間
		MaxAge: 24 * time.Hour,
	})
}

// 認証ミドルウェア
func AuthMiddleware(j *auth.JWTer) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cookieからアクセストークンを取得
		token, err := c.Cookie(AccessTokenKey)
		if err != nil {
			c.Error(err)
			errorResponse(c, http.StatusUnauthorized, NotAuthorizedCode)
			return
		}

		// JWTの検証
		if err := j.Verify(c, token); err != nil {
			if err == auth.ErrTokenExpired {
				// エラーレスポンスではなくリダイレクトさせたい
				errorResponse(c, http.StatusUnauthorized, TokenExpiredCode)
				return
			}

			c.Error(err)
			errorResponse(c, http.StatusUnauthorized, NotAuthorizedCode)
			return
		}

		it, err := c.Cookie(IdTokenKey)
		if err != nil {
			c.Error(err)
			errorResponse(c, http.StatusUnauthorized, NotAuthorizedCode)
			return
		}

		// JWTの検証
		ai, err := j.ParseIdToken(c, it)
		if err != nil {
			if err == auth.ErrTokenExpired {
				// TODO: 考える
				errorResponse(c, http.StatusUnauthorized, TokenExpiredCode)
				return
			}

			c.Error(err)
			errorResponse(c, http.StatusUnauthorized, NotAuthorizedCode)
			return
		}

		// リクエストにユーザーIDをセット
		c.Set(UserIdKey, ai.Id)

		c.Next()
	}
}
