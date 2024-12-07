package api

import (
	"log"
	"net/http"
)

func Server() {
	// Define routes and associate them with handler functions
	http.HandleFunc("/todos", GetTodos)         // GET: Fetch all todos
	http.HandleFunc("/todo", CreateTodo)        // POST: Create a new todo
	http.HandleFunc("/todo/delete", DeleteTodo) // DELETE: Delete a todo

	// Start the server
	port := ":8080"
	log.Printf("Server started at http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil)) // Log if server fails
}
