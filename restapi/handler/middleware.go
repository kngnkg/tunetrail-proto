package handler

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kngnkg/tunetrail/restapi/auth"
)

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

// 認証ミドルウェア
func AuthMiddleware(j *auth.JWTer) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Authorizationヘッダーの値を取得
		headerValue := c.GetHeader("Authorization")
		if headerValue == "" {
			errorResponse(c, http.StatusUnauthorized, NotAuthorizedCode)
			return
		}

		// Bearerプレフィックスを削除
		token := strings.TrimPrefix(headerValue, "Bearer ")

		// JWTの検証
		if err := j.Verify(c, token); err != nil {
			c.Error(err)
			errorResponse(c, http.StatusUnauthorized, NotAuthorizedCode)
			return
		}

		c.Next()
	}
}
