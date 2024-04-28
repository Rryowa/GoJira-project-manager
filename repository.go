package main

import "database/sql"

// interface to communicate with db
// and we can inject into services
type Repo interface {
	CreateUser() error
	CreateTask(task *Task) (*Task, error)
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

func (r *Repository) CreateUser() error {
	return nil
}

func (r *Repository) CreateTask(t *Task) (*Task, error) {
	rows, err := r.db.Exec("INSERT INTO tasks (name, status, project_id, assigned_to)"+
		" VALUES (?, ?, ?, ?)", t.Name, t.Status, t.ProjectID, t.AssignedTo)

	if err != nil {
		return nil, err
	}

	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}

	t.ID = id
	return t, nil
}
