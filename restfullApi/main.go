package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

var (
	tasks = []string{}
	mu    sync.Mutex
)

func main() {
	http.HandleFunc("/tasks", tasksHandler)
	http.ListenAndServe(":8080", nil)
}

// Handler para operações CRUD na lista de tarefas
func tasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getTasks(w)
	case "POST":
		addTask(w, r)
	case "PUT":
		updateTask(w, r)
	case "DELETE":
		deleteTask(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getTasks(w http.ResponseWriter) {
	mu.Lock()
	defer mu.Unlock()
	json.NewEncoder(w).Encode(tasks)
}

func addTask(w http.ResponseWriter, r *http.Request) {
	var task string
	json.NewDecoder(r.Body).Decode(&task)
	mu.Lock()
	tasks = append(tasks, task)
	mu.Unlock()
	w.WriteHeader(http.StatusCreated)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	// Implementar a lógica de atualização
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	// Implementar a lógica de deleção
}
