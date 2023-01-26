package auth

import (
	"net/http"

	"github.com/rubengomes8/golang-personal-finances/internal/http/models"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := ValidateToken(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
				ErrorMsg: "Unauthorized",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
