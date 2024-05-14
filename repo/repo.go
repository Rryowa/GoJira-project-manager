package repo

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/rryowa/Gojira-project-manager/entity"
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
	CreateUser(u *entity.User) (*entity.User, error)
	GetUserByID(id string) (*entity.User, error)
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

func (r *Repository) CreateTask(t *entity.Task) (*entity.Task, error) {
	var id int64
	err := r.db.QueryRow("INSERT INTO tasks (name, status, project_id, assigned_to)"+
		" VALUES ($1, $2, $3, $4) RETURNING id",
		t.Name, t.Status, t.ProjectID, t.AssignedTo).Scan(&id)
	if err != nil {
		return nil, err
	}

	t.ID = id
	return t, nil
}

func (r *Repository) CreateUser(u *entity.User) (*entity.User, error) {
	var id int64
	err := r.db.QueryRow("INSERT INTO users (email, firstName, lastName, password)"+
		" VALUES ($1, $2, $3, $4) RETURNING id",
		u.Email, u.FirstName, u.LastName, u.Password).Scan(&id)
	if err != nil {
		return nil, err
	}

	u.ID = id
	return u, nil
}

func (r *Repository) GetTask(id string) (*entity.Task, error) {
	var t entity.Task
	err := r.db.QueryRow("SELECT id, name, status, project_id, "+
		" assigned_to, createdAt FROM tasks WHERE id = $1", id).
		Scan(&t.ID, &t.Name, &t.Status, &t.ProjectID, &t.AssignedTo, &t.CreatedAt)
	return &t, err
}

func (r *Repository) GetUserByID(id string) (*entity.User, error) {
	var u entity.User
	err := r.db.QueryRow("SELECT id, email, firstName, "+
		"lastName, createdAt FROM users WHERE id = $1", id).
		Scan(&u.ID, &u.Email, &u.FirstName, &u.LastName, &u.CreatedAt)
	return &u, err
}
