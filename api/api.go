package api

import (
	"encoding/json"
	"go-lang/api/models"
	"go-lang/api/storage"
	"net/http"
)

// Example in-memory storage (can later use a database)
var store = storage.NewMemoryStore()

// GetTodos handles the retrieval of all todos.
// Method: GET
func GetTodos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	todos := store.GetAll() // Fetch all todos from the store
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos) // Convert todos to JSON and send as response
}

// CreateTodo handles the creation of a new todo.
// Method: POST
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	store.Add(todo) // Add the todo to the store
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Todo created successfully"))
}

// DeleteTodo handles deleting a todo by ID.
// Method: DELETE
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id") // Get the ID from query parameters
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	store.Delete(id) // Delete the todo by ID
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Todo deleted successfully"))
}
