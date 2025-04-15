package algorithm

import "fmt"

func BasicSets() {
	// Using map[string]bool to simulate a set
	// Keys are the elements, and 'true' means the element exists
	set := make(map[string]bool)

	// Add elements to the set
	set["apple"] = true
	set["banana"] = true
	set["orange"] = true

	// Try to add a duplicate (sets ignore duplicates)
	set["apple"] = true

	// Check if an element exists
	if set["banana"] {
		fmt.Println("banana is in the set")
	}

	// Remove an element
	delete(set, "orange")

	// Print all elements in the set
	fmt.Println("Set contents:")
	for item := range set {
		fmt.Println("-", item)
	}
	//- apple
	//- banana
}
