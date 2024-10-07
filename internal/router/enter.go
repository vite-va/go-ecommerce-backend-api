package router

import (
	"github.com/vite-va/go-ecommerce-backend-api/internal/router/manage"
	"github.com/vite-va/go-ecommerce-backend-api/internal/router/user"
)

type RouterGroup struct {
	User   user.UserRouterGroup
	Manage manage.ManageRouterGroup
}

var RouterGroupApp = new(RouterGroup)
