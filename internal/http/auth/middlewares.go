package auth

import (
	"net/http"

	httpModels "github.com/rubengomes8/golang-personal-finances/internal/models/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := ValidateToken(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, httpModels.ErrorResponse{
				ErrorMsg: "Unauthorized",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
