package storage

import (
	"go-lang/api/models"
	"sync"
)

// MemoryStore is an in-memory storage for todos.
type MemoryStore struct {
	mu    sync.Mutex
	todos map[string]models.Todo
}

// NewMemoryStore initializes the in-memory store.
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		todos: make(map[string]models.Todo),
	}
}

// Add adds a new todo to the store.
func (s *MemoryStore) Add(todo models.Todo) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.todos[todo.ID] = todo
}

// GetAll retrieves all todos from the store.
func (s *MemoryStore) GetAll() []models.Todo {
	s.mu.Lock()
	defer s.mu.Unlock()

	var todos []models.Todo
	for _, todo := range s.todos {
		todos = append(todos, todo)
	}
	return todos
}

// Delete removes a todo by ID.
func (s *MemoryStore) Delete(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.todos, id)
}
