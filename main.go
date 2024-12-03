package main

import (
	"go-lang/api"
	"log"
	"net/http"
)

// basics
// import "go-lang/basic"

func main() {
	// basics
	// basic.Example()
	// Define routes and associate them with handler functions
	http.HandleFunc("/todos", api.GetTodos)         // GET: Fetch all todos
	http.HandleFunc("/todo", api.CreateTodo)        // POST: Create a new todo
	http.HandleFunc("/todo/delete", api.DeleteTodo) // DELETE: Delete a todo

	// Start the server
	port := ":8080"
	log.Printf("Server started at http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil)) // Log if server fails
}
