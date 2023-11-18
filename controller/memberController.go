package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/infinity-api/model"
	"github.com/infinity-api/repository"
)

type MemberController interface {
	Get(ctx *gin.Context)
	Post(ctx *gin.Context)
}

type memberController struct {
	memberRepository repository.MemberRepository
}

func NewMemberController(memberRepository repository.MemberRepository) *memberController {
	return &memberController{
		memberRepository: memberRepository,
	}
}

func (controller *memberController) Get(ctx *gin.Context) {
	members, err := controller.memberRepository.Get()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, members)
}

func (controller *memberController) Post(ctx *gin.Context) {
	// members, err := controller.memberRepository.Get()
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, err)
	// }
	var requestBody model.Member
	var err error
	if err = ctx.BindJSON(&requestBody); err != nil {
		// DO SOMETHING WITH THE ERROR
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		}
	}

	err = controller.memberRepository.Insert(requestBody)
	if err != nil {
		log.Panic(err)
		ctx.JSON(http.StatusInternalServerError, err)
	}
	fmt.Println(requestBody.Name)

	ctx.JSON(http.StatusOK, "Inserted")
}
