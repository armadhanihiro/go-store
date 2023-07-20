package middleware

import (
	"gostore/internal/pkg/handler"
	"gostore/internal/pkg/reason"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AccessTokenVerifier interface {
	VerifyAccessToken(tokenString string) (sub string, err error)
}

func AuthMiddleware(tokenCreator AccessTokenVerifier) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken := tokenFromHeader(ctx)
		if accessToken == "" {
			handler.ResponseError(ctx, http.StatusUnauthorized, reason.Unauthorized)
			ctx.Abort()
			return
		}

		// verify token
		sub, err := tokenCreator.VerifyAccessToken(accessToken)
		if err != nil {
			handler.ResponseError(ctx, http.StatusUnauthorized, reason.Unauthorized)
			ctx.Abort()
			return
		}

		// attach sub to request (attach to ctx)
		ctx.Set("user_id", sub)

		// continue
		ctx.Next()
	}
}

func tokenFromHeader(ctx *gin.Context) string {
	var accessToken string
	bearerToken := ctx.Request.Header.Get("Authorization")
	field := strings.Fields(bearerToken)

	if len(field) != 0 && field[0] == "Bearer" {
		accessToken = field[1]
	}

	return accessToken
}
