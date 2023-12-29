package router

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	controller "github.com/infinity-api/controller"
	repo "github.com/infinity-api/repository"
	"github.com/jmoiron/sqlx"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "test"
)

func Init(db *sqlx.DB) *gin.Engine {
	router := gin.Default()
	RegisterRouter(router, db)
	return router
}

func InitDb() *sqlx.DB {
	connectionString := fmt.Sprintf("host= %s port= %d user=%s password =%s dbname = %s sslmode=disable", host, port, user, password, dbname)
	db, error := sqlx.Connect("postgres", connectionString)
	if error != nil {
		log.Fatal(error)
	}
	return db
}

func RegisterRouter(router *gin.Engine, db *sqlx.DB) {
	memberRepository := repo.NewMemberController(db)
	memberController := controller.NewMemberController(memberRepository)
	router.GET("member/get", memberController.Get)
	router.POST("member/insert", memberController.Post)
}
