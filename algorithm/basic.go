package algorithm

import "fmt"

func Basic() {
	var arr = []int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	for _, value := range arr {
		fmt.Println("value in arr", value)
	}
	fmt.Println("Capacity", cap(arr))
	arr = append(arr, 6)
	fmt.Println(arr)
	fmt.Println("Capacity after append", cap(arr))
	fmt.Println("length", len(arr))

	// 2d
	var twoDArray [4][4]int
	twoDArray[3][2] = 18
	twoDArray[2][1] = 3
	fmt.Println(twoDArray)

	var rows int
	var cols int

	rows = 5
	cols = 4
	var twoDSlice = make([][]int, rows)
	var i int
	for i = range twoDSlice {
		twoDSlice[i] = make([]int, cols)
	}
	fmt.Println("2d slice", twoDSlice)
}
