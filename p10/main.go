package main

import (
	"encoding/json"
	"log"
	"net/http"
	
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Model data
type Todo struct {
	ID        string `json:"id"`
	TitleID   string `json:"title_id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// Inisialisasi slice untuk menyimpan data
var todos []Todo

func main() {
	// Inisialisasi router
	router := mux.NewRouter()

	// Menambahkan handler untuk route "/todos"
	router.HandleFunc("/todos", GetTodos).Methods("GET")
	router.HandleFunc("/todos", CreateTodo).Methods("POST")

	// Menjalankan server pada port 8000
	log.Fatal(http.ListenAndServe(":8000", router))
}

// Handler untuk mendapatkan semua todos
func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// Handler untuk membuat todo baru
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newTodo Todo
	json.NewDecoder(r.Body).Decode(&newTodo)
	newTodo.ID = uuid.New().String()
	todos = append(todos, newTodo)
	json.NewEncoder(w).Encode(newTodo)
}
