# **Go Language Notes**

A collection of Go language concepts and examples that explore Go's core features, including variables, types, pointers, structs, arrays, slices, and maps.

---

## **Running Go Code with Nodemon**

You can use `nodemon` to automatically restart your Go program on file changes:

```sh
nodemon --exec go run main.go --ext go
```

To run tests automatically on file changes:

```sh
nodemon --exec "go test ./project" --ext go
```

To test

```sh
nodemon --exec "go test -v ./test" --ext go
```

---

## **Go Basics**

### **1. Statically Typed Language**

- Go is **statically typed**, meaning that variable types are determined at **compile time** and cannot change at runtime.
- Even though Go supports **type inference** (`:=`), once a type is assigned, it **cannot change**.

```go
num := 10  // Type inferred as int
num = "hello" // ❌ Error: Cannot assign a string to an int variable
```

### **2. Why Doesn't `:=` Work Outside Functions?**

- `:=` is **only for function-scoped variables**.
- At the package level, you **must use** the `var` keyword.

```go
package main

var globalVar = "This works" // ✅ Allowed

func main() {
    localVar := "This also works" // ✅ Allowed inside a function
}
```

---

## **Maps in Go**

### **1. Declaring and Using Maps**

```go
colors := map[string]string{
    "red":   "#ff0000",
    "green": "#00ff00",
}
colors["white"] = "#ffffff"
fmt.Println(colors)
```

### **2. Checking if a Key Exists**

```go
value, exists := colors["blue"]
if exists {
    fmt.Println("Found blue:", value)
} else {
    fmt.Println("Key 'blue' does not exist")
}
```

### **3. Iterating Over a Map**

```go
for key, value := range colors {
    fmt.Println(key, value)
}
```

### **4. Passing Maps to Functions (Reference Behavior)**

```go
func modifyMap(m map[string]string) {
    m["purple"] = "#800080"
}

modifyMap(colors)
fmt.Println(colors) // purple is added
```

### **5. Maps with Structs**

```go
type Person struct {
    Name string
    Age  int
}

people := map[string]Person{
    "user1": {"Alice", 30},
    "user2": {"Bob", 25},
}
fmt.Println(people)
```

### **6. Thread-Safe Maps (`sync.Map`)**

```go
import "sync"

var safeMap sync.Map
safeMap.Store("red", "#ff0000")

val, ok := safeMap.Load("red")
if ok {
    fmt.Println(val)
}
```

---

## **Arrays & Slices in Go**

### **1. Array vs. Slice**

| Feature             | Array             | Slice                          |
| ------------------- | ----------------- | ------------------------------ |
| **Size**            | Fixed length      | Dynamic size (can grow/shrink) |
| **Mutability**      | Cannot resize     | Can resize using `append`      |
| **Underlying Data** | Contiguous memory | References an array            |

#### Example:

```go
// Array (Fixed size)
arr := [3]int{1, 2, 3}

// Slice (Can grow)
slice := []int{1, 2, 3}
slice = append(slice, 4) // Returns a new slice with 4 added
```

---

## **Pointers in Go**

### **1. Declaring and Using Pointers**

```go
x := 58
ptr := &x // ptr now stores the memory address of x

fmt.Println(ptr)  // Prints the memory address
fmt.Println(*ptr) // Dereferencing: Prints the value stored at the memory address (58)
```

### **2. Passing by Value vs Passing by Reference**

```go
type person struct {
    firstName string
    lastName  string
}

func (p *person) updateName(newFName string) {
    p.firstName = newFName
}
```

---

## **Functions in Go**

### **1. Function Signatures**

```go
func functionName(parameter1 type, parameter2 type) returnType {
    // Function body
}
```

Example:

```go
func add(a int, b int) int {
    return a + b
}
```

### **2. Function with Pointer Receiver**

```go
func (p *person) updateName(newFName string) {
    p.firstName = newFName
}
```

---

## **Common Go Data Types**

### **1. Structs**

```go
type contactInfo struct {
    email   string
    zipcode int
}

type person struct {
    firstName string
    lastName  string
    contact   contactInfo
}

```
