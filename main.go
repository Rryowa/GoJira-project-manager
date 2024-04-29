package main

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/rryowa/go-jwt-auth/controller"
	repo "github.com/rryowa/go-jwt-auth/repo"
	postgres "github.com/rryowa/go-jwt-auth/repo/postgres"
	"github.com/rryowa/go-jwt-auth/util"
)

func main() {
	cfg := util.NewConfig()
	sqlRepository := postgres.NewSQLRepository(cfg)
	//extract db
	db, err := sqlRepository.Init()
	if err != nil {
		log.Fatal(err)
	}
	repo := repo.NewRepository(db)

	api := controller.NewAPIServer(":3000", repo)
	api.Serve()
}
