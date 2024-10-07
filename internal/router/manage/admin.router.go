package manage

import "github.com/gin-gonic/gin"

type AdminRouter struct {
}

func (ar *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	// public router
	adminRouterPublic := Router.Group("/admin")
	{
		adminRouterPublic.POST("/login")
	}
	// private router
	amdinRouterPrivate := Router.Group("/admin/user")
	// amdinRouterPrivate.Use(Limiter())
	// amdinRouterPrivate.Use(Authen())
	// amdinRouterPrivate.Use(Permission())
	{
		amdinRouterPrivate.POST("/active_user")
	}
}
