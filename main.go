package main

import (
	"fmt"
	"net/http"
)

func main() {
	// API
	http.HandleFunc("/task", handleTasks)
	http.HandleFunc("/task/", handleTaskByID)
}

func handleTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "tasks")
}

func handleTaskByID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "task by id: %s", r.URL.Path)

	// Frontend
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	fmt.Println("Servidor rodando em http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
