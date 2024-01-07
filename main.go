package main

import (
	"fmt"
	"net/http"

	"github.com/infinity-api/configuration"
	constants "github.com/infinity-api/constant"
	"github.com/infinity-api/model"
	"github.com/infinity-api/router"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	configData := createConfiguration()
	db := router.InitDb()
	router := router.Init(db)

	defer db.Close()

	err := http.ListenAndServe(fmt.Sprint(":", configData.Port), router)
	if err != nil {
		fmt.Println("**** error while starting ****")
		fmt.Println(err)
	}
	fmt.Println("started and listening in the port : ", configData.Port)
}

func createConfiguration() model.ConfigData {
	configLoader := configuration.NewConfigurationLoader()
	configData, _ := configLoader.Load(constants.AppConfig, constants.ConfigFilePath)
	return configData
}

func indexHandler(db *sqlx.DB) ([]model.Member, error) {
	var people []model.Member
	err := db.Select(&people, "SELECT * FROM company")
	if err != nil {
		return nil, err
	}
	jason, john := people[0], people[1]

	fmt.Printf("%#v\n%#v", jason, john)
	return people, nil
}
