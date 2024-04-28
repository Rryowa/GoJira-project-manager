package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type SQLRepository struct {
	db *sql.DB
}

// make connection
func NewSQLRepository(cfg *Config) *SQLRepository {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		cfg.User, cfg.Password, cfg.Name, cfg.Host, cfg.Port)

	newDb, err := sql.Open("postgres", connStr)

	// if db is down app is down also
	if err != nil {
		log.Fatal(err)
	}
	// verify db connection
	if err = newDb.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to databases")
	return &SQLRepository{db: newDb}
}

// Init() to initialize tables in db
func (r *SQLRepository) Init() (*sql.DB, error) {
	//TODO: init tables

	return r.db, nil
}
