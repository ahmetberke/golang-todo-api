package database

import (
	"github.com/jmoiron/sqlx"
	"log"
	"todo_api/app/models"
	_ "github.com/lib/pq"
)

type TodoDB interface {
	Open() error
	Close() error
	CreateTodo(p *models.Todo) error
	GetTodos() ([]*models.Todo, error)
	GetTodo(id int64) (*models.Todo, error)
	DeleteTodo(id int64) (error)
	CompleteTodo(id int64) (error)
}

type DB struct {
	db *sqlx.DB
}

func (d *DB) Open() error {
	pg, err := sqlx.Open("postgres", pgConnStr)
	if err != nil {
		return err
	}

	log.Println("Connected to DB")

	pg.MustExec(createSchema)

	d.db = pg

	return  nil
}

func (d *DB) Close() error {
	return  d.db.Close()
}
