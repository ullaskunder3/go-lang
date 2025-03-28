# **Go Language Notes**

A collection of Go language concepts and examples that explore Go's core features, including variables, types, pointers, structs, arrays, and slices.

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

## **Arrays & Slices in Go**

### **1. Array vs. Slice**

| Feature  | Array | Slice |
|----------|------------|------------|
| **Size**  | Fixed length | Dynamic size (can grow/shrink) |
| **Mutability** | Cannot resize | Can resize using `append` |
| **Underlying Data** | Contiguous memory | References an array |

#### Example:

```go
// Array (Fixed size)
arr := [3]int{1, 2, 3} 

// Slice (Can grow)
slice := []int{1, 2, 3}
slice = append(slice, 4) // Returns a new slice with 4 added
```

### **2. `append()` Does Not Modify the Original Array**

- `append()` **returns a new slice**, it does **not modify** the original one.

```go
nums := []int{1, 2, 3}
newNums := append(nums, 4)

fmt.Println(nums)    // [1 2 3]  (Original remains unchanged)
fmt.Println(newNums) // [1 2 3 4]  (New slice)
```

---

## **Pointers in Go**

Go allows the use of pointers, which are variables that store the memory address of another variable.

### **1. Declaring and Using Pointers**

- Pointers are declared using the `*` syntax, and you can access the value stored at that address using the `*` dereferencing operator.

```go
x := 58
ptr := &x // ptr now stores the memory address of x

fmt.Println(ptr)  // Prints the memory address
fmt.Println(*ptr) // Dereferencing: Prints the value stored at the memory address (58)
```

### **2. Passing by Value vs Passing by Reference**

- Go **passes variables by value**, meaning a copy is made when passing arguments to functions. However, when passing **pointers**, the function can modify the original value at the memory address.

#### Example:
```go
type person struct {
    firstName string
    lastName  string
}

func (pointerToPerson *person) updateName(newFName string) {
    (*pointerToPerson).firstName = newFName
}

func (p person) print() {
    fmt.Printf("%+v\n", p)
}

func SecondProject() {
    jack := person{
        firstName: "jack",
        lastName:  "jill",
    }

    jack.print()
    jack.updateName("jimmy")  // `updateName` modifies the original `jack` using a pointer
    jack.print()
}
```

### **3. Go's Automatic Pointer Dereferencing**

Go can **automatically convert a value to a pointer** when a method expects a pointer. This behavior is called **automatic pointer dereferencing**.

```go
func SecondProject() {
    jack := person{
        firstName: "jack",
        lastName:  "jill",
    }
    
    jack.print()
    jack.updateName("jimmy")  // Go automatically converts `jack` to a pointer
    jack.print()
}
```

---

## **Functions in Go**

### **1. Function Signatures**

In Go, functions are declared with the following syntax:

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

When defining a method for a type, you can specify whether it should receive a **pointer** to the type (to modify its state) or a **value** of the type.

```go
func (pointerToPerson *person) updateName(newFName string) {
    (*pointerToPerson).firstName = newFName
}
```

- If the receiver is a **pointer**, the method can modify the original object.
- If the receiver is a **value**, the method works with a copy and does not modify the original object.

---

## **Common Go Data Types**

### **1. Structs**

Structs are used to group different types of variables (fields) together. This is similar to objects in other languages.

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