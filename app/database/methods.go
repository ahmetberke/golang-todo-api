package database

import (
	"log"
	"time"
	"todo_api/app/models"
)


func (d *DB) CreateTodo(t *models.Todo) error {
	res, err := d.db.Exec(insertTodoSchema, t.Title, t.Content, t.IsDone, t.CreatedDate, t.CompletionDate)
	if err != nil {
		return err
	}
	_,_ = res.LastInsertId()
	return nil
}

func (d *DB) GetTodos() ([]*models.Todo, error)  {
	var Todos []*models.Todo
	err := d.db.Select(&Todos, GetTodosSchema)
	if err != nil {
		return Todos, err
	}

	return Todos, nil
}

func (d *DB) GetTodo(id int64) (*models.Todo, error) {
	var Todo models.Todo
	err := d.db.QueryRow(GetTodoSchema, id).Scan(&Todo.ID, &Todo.Title, &Todo.Content, &Todo.IsDone, &Todo.CreatedDate, &Todo.CompletionDate)
	if err != nil {
		return nil, err
	}
	return &Todo, nil
}

func (d *DB) DeleteTodo(id int64) (error) {
	_, err := d.db.Exec(DeleteTodoSchema, id)
	if err != nil {
		return err
	}
	return nil
}

func (d *DB) CompleteTodo(id int64) (error) {
	var compDate int64 = time.Now().UnixNano() / int64(time.Millisecond)

	_, err := d.db.Exec(UpdateTodoSchema, id, compDate)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}