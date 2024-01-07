package router

import (
	"fmt"
	"github.com/infinity-api/service/member"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	controller "github.com/infinity-api/controller"
	repo "github.com/infinity-api/repository"
	"github.com/jmoiron/sqlx"
)

func Init(db *sqlx.DB) *gin.Engine {
	router := gin.Default()
	RegisterRouter(router, db)
	return router
}

func InitDb() *sqlx.DB {
	host := os.Getenv("POSTGRES_HOST")
	port, _ := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	userName := os.Getenv("INFINITY_USERNAME")
	password := os.Getenv("INFINITY_PASSWORD")
	dbName := os.Getenv("INFINITY_DB")
	connectionString := fmt.Sprintf("host= %s port= %d user=%s password =%s dbname = %s sslmode=disable", host, port, userName, password, dbName)
	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func RegisterRouter(router *gin.Engine, db *sqlx.DB) {
	memberRepository := repo.NewMemberRepository(db)
	memberService := member.NewMemberService(memberRepository)
	memberController := controller.NewMemberController(memberService)
	router.GET("member/get", memberController.Get)
	router.POST("member/insert", memberController.Post)
}
