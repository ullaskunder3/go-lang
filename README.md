# **Go Language Notes**

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