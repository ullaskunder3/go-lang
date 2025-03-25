package basicI

import (
	"fmt"
	"os"
)

func BasicFirst() {
	fmt.Println("This is basic1 function from a.go")
	name := os.Getenv("USER")
	fmt.Println("os env user:->", name)
}
