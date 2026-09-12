package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var (
	tasks  = []Task{}
	nextID = 1
	mu     sync.Mutex
)

func main() {
	http.HandleFunc("/tasks", handleTasks)
	http.HandleFunc("/tasks/", handleTaskByID)

	// Servir os arquivos estáticos (frontend)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	fmt.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handleTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		mu.Lock()
		defer mu.Unlock()
		json.NewEncoder(w).Encode(tasks)

	case http.MethodPost:
		var task Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}

		mu.Lock()
		task.ID = nextID
		nextID++
		tasks = append(tasks, task)
		mu.Unlock()

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(task)

	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

func handleTaskByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	index := -1
	for i, t := range tasks {
		if t.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(tasks[index])

	case http.MethodPut:
		var updated Task
		if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}
		updated.ID = id
		tasks[index] = updated
		json.NewEncoder(w).Encode(updated)

	case http.MethodDelete:
		tasks = append(tasks[:index], tasks[index+1:]...)
		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}
