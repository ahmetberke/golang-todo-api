package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
	"todo_api/app/models"
	"todo_api/app/responses"
)

func (app *App) All() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var Res responses.Todos
		writer.Header().Set("Content-Type", "application/json")

		todos, err := app.DB.GetTodos()

		if err != nil {
			Res.Error = err.Error()
			Res.Success = false
		}else {
			Res.Success = true
			Res.Data = todos
		}

		_ = json.NewEncoder(writer).Encode(Res)
	}
}


func (app *App) Single() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var Res responses.Todo
		writer.Header().Set("Content-Type", "application/json")
		params := mux.Vars(request)
		id,err := strconv.ParseInt(params["id"], 10, 64)
		if err != nil {
			Res.Success = false
			Res.Error = err.Error()
			_ = json.NewEncoder(writer).Encode(Res)
			return
		}

		todo, err := app.DB.GetTodo(id)
		if err != nil {
			Res.Success = false
			Res.Error = err.Error()
			_ = json.NewEncoder(writer).Encode(Res)
			return
		}

		Res.Success = true
		Res.Data = todo

		_ = json.NewEncoder(writer).Encode(Res)

	}
}

func (app *App) New() http.HandlerFunc{
	return func(writer http.ResponseWriter, request *http.Request) {
		var Res responses.Todo
		writer.Header().Set("Content-Type", "application/json")

		var todo models.Todo
		_ = json.NewDecoder(request.Body).Decode(&todo)
		todo.IsDone = false
		todo.CreatedDate = time.Now().UnixNano() / int64(time.Millisecond)

		err := app.DB.CreateTodo(&todo)
		if err != nil {
			Res.Success = false
			Res.Error = err.Error()
			_ = json.NewEncoder(writer).Encode(Res)
			return
		}

		Res.Success = true
		Res.Data = &todo

		_ = json.NewEncoder(writer).Encode(Res)
	}
}

func (app *App) Delete() http.HandlerFunc{
	return func(writer http.ResponseWriter, request *http.Request) {
		var Res responses.Todo
		writer.Header().Set("Content-Type", "application/json")

		params := mux.Vars(request)
		id, err := strconv.ParseInt(params["id"], 10, 64)
		if err != nil {
			Res.Success = false
			Res.Error = err.Error()
			_ = json.NewEncoder(writer).Encode(Res)
			return
		}

		err = app.DB.DeleteTodo(id)
		if err != nil {
			Res.Success = false
			Res.Error = err.Error()
			_ = json.NewEncoder(writer).Encode(Res)
			return
		}

		Res.Success = true

		_ = json.NewEncoder(writer).Encode(Res)
		return
	}
}


func (app *App) Complete() http.HandlerFunc{
	return func(writer http.ResponseWriter, request *http.Request) {
		var Res responses.Todo
		writer.Header().Set("Content-Type", "application/json")
		params := mux.Vars(request)
		id, err := strconv.ParseInt(params["id"], 10, 64)
		if err != nil {
			Res.Success = false
			Res.Error = err.Error()
			_ = json.NewEncoder(writer).Encode(Res)
			return
		}

		err = app.DB.CompleteTodo(id)
		if err != nil {
			Res.Success = false
			Res.Error = err.Error()
			_ = json.NewEncoder(writer).Encode(Res)
			return
		}

		Res.Success = true

		_ = json.NewEncoder(writer).Encode(Res)
		return
	}
}
