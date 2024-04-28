package main

import (
	"log"

	_ "github.com/lib/pq"
)

func main() {
	cfg := InitConfig()
	sqlRepository := NewSQLRepository(cfg)
	//extract db
	db, err := sqlRepository.Init()
	if err != nil {
		log.Fatal(err)
	}
	repo := NewRepository(db)

	api := NewAPIServer(":3000", repo)
	api.Serve()
}
