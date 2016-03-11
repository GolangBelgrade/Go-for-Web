package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var db *sql.DB

func GetDB() *sql.DB {
	if db != nil {
		return db
	} else {
		db, err := sql.Open("postgres", "database=startit user=postgres sslmode=disable")
		if err != nil {
			panic(err)
		}
		return db
	}
}

type Todo struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	IsDone      bool   `json:"is_done"`
}

var TodoList []Todo

func GetTodo() Todo {
	return Todo{1, "Gopeher meetup", false}
}

func AddTodo(newItem Todo) {

	db := GetDB()

	_, err := db.Exec("INSERT INTO todo(description,is_done) VALUES($1,$2)", newItem.Description, newItem.IsDone)
	if err != nil {
		panic(err)
	}
}

func GetAll() []Todo {

	db := GetDB()

	var todoList []Todo

	rows, err := db.Query("SELECT * FROM todo")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {

		var todo Todo

		err := rows.Scan(&todo.Id, &todo.Description, &todo.IsDone)
		if err != nil {
			panic(err)
		}

		todoList = append(todoList, todo)
	}

	return todoList
}

func DeleteTodo(id int) {
	db := GetDB()

	if _, err := db.Exec("DELETE FROM todo WHERE id = $1", id); err != nil {
		panic(err)
	}

}
