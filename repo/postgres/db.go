package repo

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/rryowa/Gojira-project-manager/entity"
)

type SQLRepository struct {
	db *sql.DB
}

// make connection
func NewSQLRepository(cfg *entity.Config) *SQLRepository {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", cfg.User, cfg.Password, cfg.Name, cfg.Host, cfg.Port)

	newDb, err := sql.Open("postgres", connStr)

	// if db is down app is down also
	if err != nil {
		log.Fatal(err)
	}
	// verify db connection
	if err = newDb.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to database!")
	return &SQLRepository{db: newDb}
}

// Init() to initialize tables in db
func (r *SQLRepository) Init() (*sql.DB, error) {
	if err := r.createUsersTable(); err != nil {
		return nil, err
	}
	if err := r.createProjectsTable(); err != nil {
		return nil, err
	}
	if err := r.createTasksTable(); err != nil {
		return nil, err
	}
	return r.db, nil
}

func (r *SQLRepository) createUsersTable() error {
	_, err := r.db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL NOT NULL,
			email VARCHAR(255) NOT NULL,
			firstName VARCHAR(255) NOT NULL,
			lastName VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

			PRIMARY KEY(id),
			UNIQUE (email)
		);
		`)
	return err
}

func (r *SQLRepository) createTasksTable() error {
	_, err := r.db.Exec(`
		DROP TYPE IF EXISTS status_enum CASCADE;
		CREATE TYPE status_enum AS ENUM ('TODO', 'IN_PROGRESS', 'IN_TESTING', 'DONE');
		CREATE TABLE IF NOT EXISTS tasks (
			id SERIAL NOT NULL,
			name VARCHAR(255) NOT NULL,
			status status_enum NOT NULL DEFAULT 'TODO',
			projectId INT NOT NULL,
			assignedTo INT NOT NULL,
			createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

			PRIMARY KEY(id),
			FOREIGN KEY(assignedTo) REFERENCES users(id),
			FOREIGN KEY(projectId) REFERENCES projects(id)
		);
	`)
	return err
}

func (r *SQLRepository) createProjectsTable() error {
	_, err := r.db.Exec(`
		CREATE TABLE IF NOT EXISTS projects (
			id SERIAL NOT NULL,
			name VARCHAR(255) NOT NULL,
			createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

			PRIMARY KEY(id)
		);
		`)
	return err
}
