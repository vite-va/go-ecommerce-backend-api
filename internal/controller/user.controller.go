package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/vite-va/go-ecommerce-backend-api/internal/service"
	"github.com/vite-va/go-ecommerce-backend-api/pkg/response"
)

// type UserController struct {
// 	userService *service.UserService
// }

// func NewUserController() *UserController {
// 	return &UserController{
// 		userService: service.NewUserService(),
// 	}
// }

// // uc: user controller
// // controller --> service --> repo --> models --> dbs
// func (uc *UserController) GetUserByID(c *gin.Context) {

// 	response.SucccessResponse(c, 20001, []string{"tipjs", "m10", "vite"})
// }

// INTERFACE_VERSION

type UserController struct {
	userService service.IUserService
}

func NewUserController(
	userService service.IUserService,
) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	result := uc.userService.Register("", "")
	response.SucccessResponse(c, result, nil)
}
