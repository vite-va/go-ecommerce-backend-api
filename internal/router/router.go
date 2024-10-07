package router

import (
	"github.com/gin-gonic/gin"
	c "github.com/vite-va/go-ecommerce-backend-api/internal/controller"
	"github.com/vite-va/go-ecommerce-backend-api/internal/middlewares"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// use the middlewares
	r.Use(middlewares.AuthenMiddleware())
	v1 := r.Group("v1")
	v1.GET("/ping", c.NewPongController().Pong)
	return r
}

// func Pong(ctx *gin.Context) {
// 	name := ctx.DefaultQuery("name", "vite-va")
// 	uid := ctx.Query("uid")
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"message": "ping...pong! " + name,
// 		"uid":     uid,
// 		"users":   []string{"cr7", "messi", "anonystick"},
// 	})
// }
