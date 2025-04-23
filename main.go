package main

import (
	"log"
	"net/http"

	"github.com/nabilulilallbab/TodoList/handle"
	"github.com/nabilulilallbab/TodoList/library"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", library.CreateHandler(handle.Index))

	mux.HandleFunc("/create", library.CreateHandler(handle.CreateForm))

	mux.HandleFunc("/store", library.CreateHandler(handle.CreateTodo))
	mux.HandleFunc("/toggle", library.CreateHandler(handle.ToggleTodo))
	mux.HandleFunc("/delete", library.CreateHandler(handle.DeleteTodo))
	mux.HandleFunc("/edit", library.CreateHandler(handle.EditForm))
	mux.HandleFunc("/update", library.CreateHandler(handle.UpdateTodo))
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))


	log.Println("Listen port :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
