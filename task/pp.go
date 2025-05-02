package task

import "fmt"

type User struct {
	name string
}

/*
Before: Ullas
After:  GoHero
*/

// ex 1
// 🎯 Your function here
func renameUser(pp **User) {
	// 🔥 Modify the name field of the actual User using **pp
	(**pp).name = "GoHero"
}

// ex2

// 🎯 Your function here
func doubleSliceValues(p *[]int) {
	// multiply each value by 2
	fmt.Println("p=>", *p)
	// for num := range *p {
	// 	(*p)[num] *= 2
	// }
	for index, num := range *p {
		(*p)[index] = num * 2
	}
}

// 🎯 Your function here
// Slice of Struct Pointers via Pointer to Slice
func updateUserNames(users *[]*User) {
	// loop and update each user's name to "GoMaster"
	for i, _ := range *users {
		(*users)[i].name = "GoMaster"
	}
}

func PointerToPointer() {
	user := User{name: "Ullas"}
	p := &user
	pp := &p

	fmt.Println("Before:", user.name)

	renameUser(pp)

	fmt.Println("After:", user.name)
	fmt.Println("------------------")

	nums := []int{1, 2, 3, 4}
	fmt.Println("Before:", nums)

	doubleSliceValues(&nums)

	fmt.Println("After:", nums)
	fmt.Println("------------------")

	u1 := &User{name: "Ullas"}
	u2 := &User{name: "Ken"}
	u3 := &User{name: "Taro"}

	userList := []*User{u1, u2, u3}
	fmt.Print("Before: [")
	for _, u := range userList {
		fmt.Print((*u).name, " ") // explictly
	}
	fmt.Println("]")

	updateUserNames(&userList)

	fmt.Print("After:  [")
	for _, u := range userList {
		fmt.Print(u.name, " ") //implecit automatically dereferences u to access name
	}
	fmt.Println("]")

}
