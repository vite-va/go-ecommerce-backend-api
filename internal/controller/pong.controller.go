package controller

import "github.com/gin-gonic/gin"

type PongController struct{}

func NewPongController() *PongController {
	return &PongController{}
}

func (p *PongController) Pong(c *gin.Context) {

}
