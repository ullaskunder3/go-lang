package models

// Todo represents a task item.
type Todo struct {
	ID    string `json:"id"`    // Unique identifier
	Title string `json:"title"` // Title of the task
	Done  bool   `json:"done"`  // Completion status
}
