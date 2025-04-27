package basic

import "fmt"

func PointerBasic() {
	b := 255
	var a *int = &b

	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)

}
