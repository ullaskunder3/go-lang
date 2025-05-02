package functions

import (
	"fmt"
	"time"
)

// Function factory: Returns a function that multiplies by `factor`
func multiplier(factor int) func(int) int {
	return func(num int) int {
		return num * factor // captures `factor` from outer scope
	}
}

// Closure with captured state: maintains a private `count` variable
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// Higher-order function: Accepts a function as a parameter
func repeat(n int, fn func(int)) {
	for i := 0; i < n; i++ {
		fn(i)
	}
}

// Returns a function: Used in functional-style middleware or logic building
func withPrefix(prefix string) func(string) {
	return func(msg string) {
		fmt.Println(prefix + msg)
	}
}

// Named return values
func profile(inputName string, intputAge int) (name string, age int) {
	name = inputName
	age = intputAge
	return
}

// variadic function
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func Basic() {
	// --- 1. IIFE: Immediately Invoked Function Expression ---
	msg := func() string {
		return "Hello from IIFE"
	}()
	fmt.Println(msg)

	// --- 2. Anonymous function as Goroutine Callback ---
	go func(name string) {
		fmt.Printf("Hello %s from a Goroutine\n", name)
	}("Ullas")

	time.Sleep(100 * time.Millisecond)

	// --- 3. Function Factories with Closures ---
	double := multiplier(2)
	fmt.Println("multiplier(2) -> double(2)", double(2)) // 4

	// --- 4. Closure for Capturing State ---
	c := counter()
	fmt.Println(c()) // 1
	fmt.Println(c()) // 2
	fmt.Println(c()) // 3

	anotherCounter := counter()
	fmt.Println(anotherCounter()) // 1

	// --- 5. Higher-order function (function as argument) ---
	repeat(3, func(i int) {
		fmt.Println("Repeat #", i+1)
	})

	// --- 6. Returning a function ---
	greeter := withPrefix("Hi, ")
	greeter("Ullas") // Hi, Ullas

	// 7 named return value
	fmt.Println(profile("ullas", 20)) //ukkas 20

	// 8 variadic function
	fmt.Println(sum(1, 2, 3, 4)) //10

	var a int = 42
	var p *int = &a
	var pp **int = &p
	fmt.Println("&a", &a)
	fmt.Println("pp", pp)
	fmt.Println("*pp", *pp)
	fmt.Println("**pp", **pp)

}
