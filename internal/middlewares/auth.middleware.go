package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/vite-va/go-ecommerce-backend-api/pkg/response"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorResponse(ctx, response.ErrInvalidToken, "")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
