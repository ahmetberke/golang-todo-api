package app

func (app *App) initRoutes()  {

	// ROUTES
	app.Router.HandleFunc("/", app.All()).Methods("GET")
	app.Router.HandleFunc("/{id}", app.Single()).Methods("GET")
	app.Router.HandleFunc("/", app.New()).Methods("POST")
	app.Router.HandleFunc("/{id}", app.Delete()).Methods("DELETE")
	app.Router.HandleFunc("/{id}", app.Complete()).Methods("PUT")

}