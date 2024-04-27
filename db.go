package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

type MySQLRepository struct {
	db *sql.DB
}

// NewMySQLRepository - initialize db and make the connection
func NewMySQLRepository(cfg mysql.Config) *MySQLRepository {
	newDb, err := sql.Open("mysql", cfg.FormatDSN())

	// if db is down app is down also
	if err != nil {
		log.Fatal(err)
	}
	// verify db connection
	if err = newDb.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MySQL")
	return &MySQLRepository{db: newDb}
}

// Init() to initialize tables in db
func (r *MySQLRepository) Init() (*sql.DB, error) {
	//TODO: init tables

	return r.db, nil
}
