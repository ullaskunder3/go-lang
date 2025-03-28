package project

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo //can also be just contactInfo will work same
}

func SecondProject() {
	// luffy := person{"luffy", "monkey"}
	// luffy := person{firstName: "luffy", lastName: "monkey"}

	// var luffy person
	// luffy.firstName = "luffy"
	// luffy.lastName = "monkey"

	// fmt.Printf("%+v", luffy)
	// fmt.Println("")

	// ----------
	jack := person{
		firstName: "jack",
		lastName:  "jill",
		contact: contactInfo{
			email:   "jak@gmail.com",
			zipcode: 1999,
		},
	}

	// jack.print()
	// jackPointer := &jack
	// jackPointer.updateName("jimmy")
	// jack.print()

	jack.print()
	// Go can automatically convert the value to a pointer when necessary.
	// This is called "automatic pointer dereferencing
	jack.updateName("jimmy")
	jack.print()

	// -------
	name := "mr crabs"

	namePointer := &name

	fmt.Println(&namePointer)
	printPointer(namePointer)

}
func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
}
func (pointerToperson *person) updateName(newFName string) {
	(*pointerToperson).firstName = newFName
}

func (p person) print() {
	fmt.Printf("%+v", p)
	fmt.Println("")
}
