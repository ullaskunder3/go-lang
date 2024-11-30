package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")
	var testVariable = "this is mr awesome ="
	testVariableNewWay := "this is mr awesome :="
	// list of string
	var taskList = []string{"this is awesome", "wow", "go"}

	fmt.Println(testVariable)
	fmt.Println(testVariableNewWay)
	fmt.Println(taskList[0])
	fmt.Println(taskList[1:3])

	for index, task := range taskList {
		// fmt.Println(task, index)
		fmt.Printf("%d, %s\n", index+1, task)
	}

	name := "ullas"
	greet(name)
}

func greet(name string) {
	fmt.Println("greetings...")
	fmt.Printf("Hello %s\n", name)
}
