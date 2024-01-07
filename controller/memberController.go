package controller

import (
	"github.com/infinity-api/service/member"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MemberController interface {
	Get(ctx *gin.Context)
	Post(ctx *gin.Context)
}

type memberController struct {
	memberService member.MemberService
}

func NewMemberController(memberService member.MemberService) MemberController {
	return &memberController{
		memberService: memberService,
	}
}

func (controller *memberController) Get(ctx *gin.Context) {
	members, err := controller.memberService.GetAllMembers()
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
	//var requestBody model.Member
	//var err error
	//if err = ctx.BindJSON(&requestBody); err != nil {
	//	// DO SOMETHING WITH THE ERROR
	//	if err != nil {
	//		ctx.JSON(http.StatusInternalServerError, err)
	//	}
	//}
	//
	//err = controller.memberRepository.Insert(requestBody)
	//if err != nil {
	//	log.Panic(err)
	//}
	//fmt.Println(requestBody.Name)

	ctx.JSON(http.StatusOK, "Inserted")
}
