package main

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/rryowa/Gojira-project-manager/controllers"
	"github.com/rryowa/Gojira-project-manager/repo"
	postgres "github.com/rryowa/Gojira-project-manager/repo/postgres"
	"github.com/rryowa/Gojira-project-manager/utils"
)

func main() {
	cfg := utils.NewConfig()
	sqlRepository := postgres.NewSQLRepository(cfg)
	//extract db
	db, err := sqlRepository.Init()
	if err != nil {
		log.Fatal(err)
	}
	repo := repo.NewRepository(db)

	api := controllers.NewAPIServer(":3000", repo)
	api.Serve()
}
