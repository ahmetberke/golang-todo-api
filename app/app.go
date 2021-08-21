package app

import (
	"github.com/gorilla/mux"
	"todo_api/app/database"
)

type App struct {
	Router *mux.Router
	DB 		database.TodoDB
}

func New() *App {
	// Creating App
	app:= &App{
		Router: mux.NewRouter(),
	}

	// Routes are included
	app.initRoutes()

	return app
}