package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MemberController interface {
	Get()
}

type memberController struct {
}

func NewMemberController() *memberController {
	return &memberController{}
}

func (controller *memberController) Get(ctx *gin.Context) {
	fmt.Println("here in the controller...")
	ctx.JSON(http.StatusOK, "woowww created new service")
}
