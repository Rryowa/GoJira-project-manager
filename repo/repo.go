package repo

import (
	"database/sql"

	"github.com/rryowa/go-jwt-auth/entity"
)

/*
The layer is responsible for interacting with storage
(database, file system, and so on). Other layers must
use abstraction (interfaces) for interacting with this
layer.
*/
// interface to communicate with db
// and we can inject into services
type Repo interface {
	CreateUser() error
	CreateTask(task *entity.Task) (*entity.Task, error)
	GetTask(id string) (*entity.Task, error)
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

func (r *Repository) CreateTask(t *entity.Task) (*entity.Task, error) {
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

func (r *Repository) GetTask(id string) (*entity.Task, error) {
	var t entity.Task
	err := r.db.QueryRow("SELECT id, name, status, project_id, "+
		" assigned_to, createdAt FROM tasks WHERE id = ?", id).
		Scan(&t.ID, &t.Name, &t.Status, &t.ProjectID, &t.AssignedTo, &t.CreatedAt)
	return &t, err
}
