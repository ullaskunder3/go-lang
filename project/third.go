package project

import (
	"fmt"
	"sync"
)

func ThirdProject() {
	// Declaring a map where all keys are strings and values are strings
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
	}

	// Another way to declare an empty map
	// var colors map[string]string
	// colors := make(map[string]string)

	colors["white"] = "#ffffff"

	fmt.Println("Initial Colors:", colors)

	// Iterating over the map
	printMap(colors)

	// Checking if a key exists
	value, exists := colors["blue"]
	if exists {
		fmt.Println("Found blue:", value)
	} else {
		fmt.Println("Key 'blue' does not exist")
	}

	// Deleting a key
	delete(colors, "green")
	fmt.Println("After deleting green:", colors)

	// Demonstrating maps with different data
	m := map[string]string{
		"dog": "bark",
		"cat": "purr",
	}

	for _, value := range m {
		fmt.Println(value)
	}

	// Maps with structs
	people := map[string]Person{
		"user1": {"Alice", 30},
		"user2": {"Bob", 25},
	}
	fmt.Println("People:", people)

	// Modifying map inside a function (shows reference behavior)
	modifyMap(colors)
	fmt.Println("After modifying inside function:", colors)

	// Using sync.Map for concurrency safety
	var safeMap sync.Map

	safeMap.Store("red", "#ff0000")
	safeMap.Store("blue", "#0000ff")

	val, ok := safeMap.Load("red")
	if ok {
		fmt.Println("From sync.Map - red:", val)
	}
}

type Person struct {
	Name string
	Age  int
}

func printMap(c map[string]string) {
	for i, col := range c {
		fmt.Println(i, col)
	}
}

func modifyMap(m map[string]string) {
	m["purple"] = "#800080"
}
