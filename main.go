package main

import (
	"encoding/json"
	"golang.org/x/net/websocket"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./static/"))

	mux.Handle("/", http.HandlerFunc(rootHandler))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.Handle("/todo", http.HandlerFunc(todoHandler))

	mux.Handle("/liveadd", websocket.Handler(liveAddHandler))

	http.ListenAndServe(":8080", mux)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Web server"))
}

//gorilla
// negroni
// martini
// bone
// httprouter

func todoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":

		todoList := GetAll()

		result, err := json.Marshal(todoList)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(result)

		break
	case "POST":

		decoder := json.NewDecoder(r.Body)
		var todo Todo

		err := decoder.Decode(&todo)
		if err != nil {
			panic(err)
		}
		AddTodo(todo)
		w.WriteHeader(http.StatusOK)
		break
	case "DELETE":
		var id int

		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&id); err != nil {
			panic(err)
		}

		DeleteTodo(id)

		break
	}

}

func liveAddHandler(conn *websocket.Conn) {

	var todo Todo
	for {
		websocket.JSON.Receive(conn, &todo)
		AddTodo(todo)
		websocket.JSON.Send(conn, GetAll())

	}
}
