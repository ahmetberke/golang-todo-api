package models

type Todo struct {
	ID				int64	`json:"id" db:"id"`
	Title			string	`json:"title" db:"title"`
	Content 		string	`json:"content" db:"content"`
	IsDone			bool	`json:"is_done" db:"is_done"`
	CreatedDate		int64	`json:"created_date" db:"created_date"`
	CompletionDate 	int64	`json:"completion_date" db:"completion_date"`
}