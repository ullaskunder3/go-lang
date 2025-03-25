package error

import (
	"errors"
	"fmt"
	"log"
)

// ----------------------------
// 1. Basic Error Handling
// ----------------------------

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

// ----------------------------
// 2. Custom Error Types
// ----------------------------

type MyError struct {
	Function string
	Message  string
}

// Implementing the error interface
func (e *MyError) Error() string {
	return fmt.Sprintf("Error in %s: %s", e.Function, e.Message)
}

func someFunc() error {
	return &MyError{"someFunc", "Something went wrong"}
}

// ----------------------------
// 3. Wrapping and Unwrapping Errors
// ----------------------------

func wrappedError() error {
	baseErr := errors.New("original error")
	return fmt.Errorf("wrapping error: %w", baseErr) // Wrap error
}

func checkWrappedError() {
	err := wrappedError()
	if errors.Is(err, errors.New("original error")) { // Incorrect way (Fix below)
		fmt.Println("Matched original error")
	} else {
		fmt.Println("Error does not match")
	}

	// Correct way: Define the original error as a package-level variable
	originalErr := errors.New("original error")
	if errors.Is(err, originalErr) {
		fmt.Println("Matched original error (correct check)")
	}

	unwrapped := errors.Unwrap(err)
	fmt.Println("Unwrapped error:", unwrapped)
}

// ----------------------------
// 4. Error Propagation
// ----------------------------

func readFile(filename string) error {
	if filename == "" {
		return fmt.Errorf("readFile failed: %w", errors.New("filename is empty"))
	}
	return nil
}

func processFile(filename string) error {
	err := readFile(filename)
	if err != nil {
		return fmt.Errorf("processFile failed: %w", err)
	}
	return nil
}

// ----------------------------
// 5. Type Assertion for Custom Errors
// ----------------------------

func checkCustomError() {
	err := someFunc()

	var myErr *MyError
	if errors.As(err, &myErr) { // Check if error is of type MyError
		fmt.Println("Custom Error Matched:", myErr.Function, "-", myErr.Message)
	} else {
		fmt.Println("Error does not match MyError type")
	}
}

// ----------------------------
// 6. Using Panic and Recover for Fatal Errors
// ----------------------------

func riskyOperation() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	fmt.Println("Starting risky operation...")
	panic("Something went terribly wrong!") // This will be recovered
	fmt.Println("This will not be printed")
}

// ----------------------------
// 7. Logging Errors
// ----------------------------

func logErrorExample() error {
	err := errors.New("something went wrong")
	log.Printf("ERROR: %v", err) // Proper logging
	return err
}

// ----------------------------
// 8. Main Function
// ----------------------------

func BasicErrorHandelling() {
	// 1. Basic Error Handling
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// 2. Custom Error Type
	err = someFunc()
	if err != nil {
		fmt.Println("Custom Error:", err)
	}

	// 3. Wrapping and Unwrapping Errors
	checkWrappedError()

	// 4. Error Propagation
	err = processFile("")
	if err != nil {
		fmt.Println("File Processing Error:", err)
	}

	// 5. Type Assertion for Custom Errors
	checkCustomError()

	// 6. Panic and Recover
	riskyOperation()

	// 7. Logging Errors
	logErrorExample()
}
