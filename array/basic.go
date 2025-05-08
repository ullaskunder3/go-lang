package array

func Map(nums []int, fn func(int) int) []int {
	result := make([]int, len(nums))
	for i, val := range nums {
		result[i] = fn(val)
	}
	return result
}

func Filter(nums []int, fn func(int) bool) []int {
	result := []int{}
	for _, val := range nums {
		if fn(val) {
			result = append(result, val)
		}
	}
	return result
}

func Reduce(nums []int, fn func(int, int) int, initial int) int {
	result := initial
	for _, val := range nums {
		result = fn(result, val)
	}
	return result
}
