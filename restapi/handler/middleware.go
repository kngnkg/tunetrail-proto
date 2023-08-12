package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kngnkg/tunetrail/restapi/auth"
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
			"Authorization", // 不要になるかも
		},
		// 許可したいアクセス元の一覧
		AllowOrigins: []string{
			"https://" + allowedDomain,
			"https://www." + allowedDomain,
			"https://api." + allowedDomain,
		},
		// preflight requestで許可した後の接続可能時間
		MaxAge: 24 * time.Hour,
	})
}

// 認証ミドルウェア
func AuthMiddleware(j *auth.JWTer) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cookieからアクセストークンを取得
		token, err := c.Cookie("accessToken")
		if err != nil {
			c.Error(err)
			errorResponse(c, http.StatusUnauthorized, NotAuthorizedCode)
			return
		}

		log.Println("token: ", token)

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

		c.Next()
	}
}
