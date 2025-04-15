package algorithm

import "fmt"

func BasicArray() {
	myArr := [8]int{1, 2, 3, 4}
	newArray := []int{1, 2, 3, 7}

	// Convert the fixed-size array to a slice using [:]
	// This is needed because append() only works with slices, not arrays.
	// myArr[:] creates a slice from the full array (from index 0 to end)

	// The '...' operator unpacks the slice (newArray)
	// This is required because append() needs individual elements,
	// not a single slice as an argument.
	combined := append(myArr[:], newArray...)

	fmt.Println("Combined:", combined)
}
