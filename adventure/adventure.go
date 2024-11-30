package adventure

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

// GoMaster represents a learning journey
type GoMaster struct {
	Name      string
	Level     int
	Knowledge map[string]bool
}

// NewGoMaster creates a new learning adventure
func NewGoMaster(name string) *GoMaster {
	return &GoMaster{
		Name:      name,
		Level:     1,
		Knowledge: make(map[string]bool),
	}
}

// CompleteGoMasterAdventure is the main learning journey
func CompleteGoMasterAdventure() {
	// Create a new Go learning adventure
	master := NewGoMaster("CodeNinja")

	// Interactive Learning Sections
	master.BasicDataTypes()
	master.ControlStructures()
	master.FunctionsAndErrorHandling()
	master.AdvancedDataStructures()
	master.PointersAndReflection()
	master.ConcurrencyMastery()
	master.FileIOAndNetworking()
	master.AdvancedTopics()
}

// BasicDataTypes covers fundamental data types
func (gm *GoMaster) BasicDataTypes() {
	fmt.Println("🌟 Chapter 1: Mastering Basic Data Types 🌟")

	// Primitive Types
	var intVar int = 42
	var floatVar float64 = 3.14
	var stringVar string = "GoNinja"
	var boolVar bool = true

	// Type Inference
	inferredInt := 100
	// inferredString := "Dynamic Typing"

	// Complex Types
	numbers := []int{1, 2, 3, 4, 5}
	mapping := map[string]int{
		"apple":  5,
		"banana": 7,
	}

	// Type Conversion
	floatToInt := int(floatVar)
	stringToInt, _ := strconv.Atoi("123")

	fmt.Printf("Data Types Showcase:\n")
	fmt.Printf("Integer: %d\n", intVar)
	fmt.Printf("Float: %f\n", floatVar)
	fmt.Printf("String: %s\n", stringVar)
	fmt.Printf("Boolean: %t\n", boolVar)
	fmt.Printf("Inferred Int: %d\n", inferredInt)
	fmt.Printf("Numbers: %v\n", numbers)
	fmt.Printf("Mapping: %v\n", mapping)
	fmt.Printf("Float to Int: %d\n", floatToInt)
	fmt.Printf("String to Int: %d\n", stringToInt)
}

// ControlStructures demonstrates flow control
func (gm *GoMaster) ControlStructures() {
	fmt.Println("\n🌈 Chapter 2: Mastering Control Structures 🌈")

	// If-Else Conditions
	score := 75
	if score >= 90 {
		fmt.Println("Excellent!")
	} else if score >= 70 {
		fmt.Println("Good Job!")
	} else {
		fmt.Println("Keep Practicing!")
	}

	// Switch Statement
	language := "Go"
	switch language {
	case "Go":
		fmt.Println("Go is awesome!")
	case "Python":
		fmt.Println("Python is cool!")
	default:
		fmt.Println("Many great languages!")
	}

	// For Loops
	for i := 0; i < 5; i++ {
		fmt.Printf("Iteration %d\n", i)
	}

	// Range Loop
	fruits := []string{"Apple", "Banana", "Cherry"}
	for index, fruit := range fruits {
		fmt.Printf("%d: %s\n", index, fruit)
	}
}

// FunctionsAndErrorHandling covers function concepts
func (gm *GoMaster) FunctionsAndErrorHandling() {
	fmt.Println("\n🚀 Chapter 3: Functions & Error Handling 🚀")

	// Multiple Return Values
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Division Result: %f\n", result)
	}

	// Anonymous Functions
	multiplier := func(a, b int) int {
		return a * b
	}
	fmt.Printf("Multiplication: %d\n", multiplier(5, 6))

	// Defer, Panic, Recover
	deferDemo()
}

// divide demonstrates error handling
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// deferDemo shows defer mechanism
func deferDemo() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Deferred function demonstration")
}

// AdvancedDataStructures explores complex data handling
func (gm *GoMaster) AdvancedDataStructures() {
	fmt.Println("\n🔬 Chapter 4: Advanced Data Structures 🔬")

	// Slices
	numbers := []int{1, 2, 3, 4, 5}
	numbers = append(numbers, 6, 7)
	fmt.Println("Slice:", numbers)

	// Maps
	userScores := make(map[string]int)
	userScores["Alice"] = 95
	userScores["Bob"] = 87
	fmt.Println("User Scores:", userScores)

	// Structs
	type Player struct {
		Name  string
		Score int
	}

	players := []Player{
		{Name: "Alice", Score: 95},
		{Name: "Bob", Score: 87},
	}
	fmt.Println("Players:", players)
}

// PointersAndReflection demonstrates pointer and reflection concepts
func (gm *GoMaster) PointersAndReflection() {
	fmt.Println("\n🔮 Chapter 5: Pointers & Reflection 🔮")

	// Pointers
	x := 10
	changeValue(&x)
	fmt.Println("Modified Value:", x)

	// Reflection
	typeOfStruct := reflect.TypeOf(struct{}{})
	fmt.Println("Type of empty struct:", typeOfStruct)
}

// changeValue modifies value through pointer
func changeValue(ptr *int) {
	*ptr = 20
}

// ConcurrencyMastery demonstrates concurrent programming
func (gm *GoMaster) ConcurrencyMastery() {
	fmt.Println("\n⚡ Chapter 6: Concurrency Mastery ⚡")

	// WaitGroup for synchronization
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d working\n", id)
			time.Sleep(time.Second)
		}(i)
	}

	wg.Wait()
	fmt.Println("All goroutines completed")
}

// FileIOAndNetworking covers file and network interactions
func (gm *GoMaster) FileIOAndNetworking() {
	fmt.Println("\n📡 Chapter 7: File I/O & Networking 📡")

	// User Input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Printf("Hello, %s!\n", name)
}

// AdvancedTopics covers additional complex concepts
func (gm *GoMaster) AdvancedTopics() {
	fmt.Println("\n🌠 Chapter 8: Advanced Go Concepts 🌠")

	// Interfaces
	var shape Shape = &Circle{Radius: 5}
	fmt.Printf("Circle Area: %f\n", shape.Area())
}

// Shape interface for polymorphism
type Shape interface {
	Area() float64
}

// Circle struct implementing Shape
type Circle struct {
	Radius float64
}

// Area calculates circle area
func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Adventure() {
	// Set random seed
	rand.Seed(time.Now().UnixNano())

	// Log setup
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Print system info
	fmt.Printf("OS: %s\n", runtime.GOOS)
	fmt.Printf("Arch: %s\n", runtime.GOARCH)

	// Start the Go Master Adventure
	CompleteGoMasterAdventure()
}
