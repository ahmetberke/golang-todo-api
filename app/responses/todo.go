package responses

import "todo_api/app/models"

type Todo struct {
	Success		bool			`json:"success"`
	Error		string			`json:"error"`
	Data 		*models.Todo 	`json:"data"`
}

type Todos struct {
	Success		bool			`json:"success"`
	Error		string			`json:"error"`
	Data 		[]*models.Todo 	`json:"data"`
}


