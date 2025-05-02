package basic

import "fmt"

type User struct {
	email string
}

// Value receiver (makes a copy)
func (u User) ChangeEmail(newEmail string) {
	u.email = newEmail
	fmt.Println("[Inside ChangeEmail] Updated email to:", u.email)
}

// Pointer receiver (modifies original)
func (u *User) ChangeEmailPointer(newEmail string) {
	u.email = newEmail
	fmt.Println("[Inside ChangeEmailPointer] Updated email to:", u.email)
}

func PointerBasic() {
	// Creating an instance of User with an email field
	user := User{
		email: "ullaskunder3@gmail.com",
	}

	// Working with pointers in Go
	var a int = 10

	// Declaring a pointer 'p' and assigning it the address of 'a'
	var p *int = &a

	// Logging the address of 'a' stored in the pointer 'p'
	fmt.Println("Pointer 'p' is pointing to address:", p) // e.g., 0xc00008c0a8
	// Dereferencing the pointer 'p' to print the value at that address
	fmt.Println("Value stored at address *p:", *p) // *p => 10

	// Changing the value of 'a' through the pointer 'p'
	*p = 20

	// Logging the updated value of 'a'
	fmt.Println("Updated value of 'a':", a) // a => 20
	// Checking the value at the pointer location after the update
	fmt.Println("Updated value through pointer *p:", *p) // *p => 20

	// Calling value receiver (copy)
	user.ChangeEmail("copy@example.com")
	fmt.Println("[After ChangeEmail] User's email:", user.email) // ❌ Not changed

	// Calling pointer receiver (modifies original)
	user.ChangeEmailPointer("pointer@example.com")
	fmt.Println("[After ChangeEmailPointer] User's email:", user.email) // ✅ Changed
}
