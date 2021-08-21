package database

const createSchema = `
CREATE TABLE IF NOT EXISTS todos
(
	id SERIAL PRIMARY KEY,
	title TEXT,
	content TEXT,
	is_done BOOL,
	created_date BIGINT,
	completion_date BIGINT
)
`

var insertTodoSchema = `
INSERT INTO todos(title, content, is_done, created_date, completion_date) VALUES($1,$2,$3,$4,$5) RETURNING id
`

var GetTodoSchema = `
SELECT * FROM todos WHERE id = $1
`

var GetTodosSchema = `
SELECT * FROM todos
`

var DeleteTodoSchema = `
DELETE FROM todos WHERE id = $1
`

var UpdateTodoSchema = `
UPDATE todos SET is_done = NOT is_done, completion_date=$2 WHERE id=$1
`