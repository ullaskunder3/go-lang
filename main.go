package main

import (
	// "go-lang/project"
	// "go-lang/basic"
	// "go-lang/functions"
	// "go-lang/api"
	// "go-lang/basicI"
	// "go-lang/book"
	// "go-lang/project"
	// "go-lang/task"
	// "go-lang/error"
	// "go-lang/algorithm"
	// "go-lang/system"
	"fmt"
	"go-lang/array"
	"go-lang/leetcode"
)

func main() {
	// basics
	// basic.Example()

	// basicI
	// basicI.BasicFirst()

	// error handling
	// error.BasicErrorHandelling()

	// first project
	// project.FirstProject()
	// project.SecondProject()
	// project.ThirdProject()
	// project.FourthProject()

	// from book
	// book.CalebDox()

	// api
	// api.Server()

	// algorithm
	// result1 := algorithm.Fibo(6)
	// fmt.Println(result1)

	// task prettyjson
	// task.PrettyJSON()

	// task cli password generator
	// task.CliPasswordGen()
	// task.AdvanceCliPass()

	// algorithm
	// algorithm.Basic()
	// algorithm list
	// algorithm.GoList()
	// algorithm.BasicArray()
	// algorithm.BasicSets()

	// basic of interface
	// basic.InterfaceUHWork()

	// build system book
	// system.PassingArgument()

	// project
	// project.FifthProjectHttp()

	// basic pointer
	// basic.PointerBasic()

	//functions
	// functions.Basic()

	// task
	// task.PointerToPointer()

	// input := "Echo"
	// modifier := func(s string) string {
	// 	return s + "!"
	// }

	// log := task.Transform(input, modifier)
	// fmt.Println(log)

	// leetcode
	arr := []int{1, 2, 3, 4, 5}
	target := 9
	log := leetcode.TwoSum(arr, target)
	fmt.Println("two sum:", log)
	arr2 := []int{1, 5, 3, 4, 5}
	log2 := leetcode.AllTwoSums(arr2, target)
	fmt.Println("all two sum", log2)

	double := array.Map(arr, func(i int) int {
		return i * 2
	})
	fmt.Println("Array Map function:", double)

	evens := array.Filter(arr, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("Evens in array", evens)

	sum := array.Reduce(arr, func(acc, val int) int {
		return acc + val
	}, 0)
	fmt.Println("sum in array", sum)
}
