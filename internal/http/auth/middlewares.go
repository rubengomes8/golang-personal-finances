package auth

import (
	"net/http"

	"github.com/rubengomes8/golang-personal-finances/internal/http/models"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := validateToken(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, models.ErrorResponse{
				ErrorMsg: "Unauthorized",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
