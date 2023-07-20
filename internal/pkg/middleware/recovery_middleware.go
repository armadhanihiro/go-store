package middleware

import (
	"gostore/internal/pkg/reason"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			err := recover()
			if err != nil {
				log.Error(err)
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"code":    http.StatusInternalServerError,
					"status":  http.StatusText(http.StatusInternalServerError),
					"message": reason.InternalServerError,
				})
			}
		}()

		ctx.Next()
	}
}
