package main

import "database/sql"

// interface to communicate with db
// and we can inject into services
type Repo interface {
	CreateUser() error
}

type Repository struct {
	db *sql.DB
}

// Constructor
func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (s *Repository) CreateUser() error {
	return nil
}
