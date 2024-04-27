package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {
	//To create db connection we need config
	cfg := mysql.Config{
		User:      Envs.DBUser,
		Passwd:    Envs.DBPassword,
		Addr:      Envs.DBAddress,
		DBName:    Envs.DBName,
		Net:       "tcp",
		ParseTime: true,
	}

	sqlRepository := NewMySQLRepository(cfg)
	//extract db
	db, err := sqlRepository.Init()
	if err != nil {
		log.Fatal(err)
	}

	repo := NewRepository(db)
	api := NewAPIServer(":3000", repo)
	api.Serve()
}
