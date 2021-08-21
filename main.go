package main

import (
	"log"
	"os"
	"todo_api/app"
	"todo_api/app/database"
)

func main()  {
	a := app.New()
	a.DB = &database.DB{}
	err := a.DB.Open()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer func() {
		_ = a.DB.Close()
	}()


	server := NewServer(":9009", nil)
	server.NewApp("/", a.Router.ServeHTTP)
	server.Run()
}

