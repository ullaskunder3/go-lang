package leetcode

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		diff := target - num
		if j, ok := m[diff]; ok {
			return []int{j, i}
		}
		m[num] = i
	}
	return nil
}
func AllTwoSums(nums []int, target int) [][]int {
	m := make(map[int][]int)
	var result [][]int

	for i, num := range nums {
		diff := target - num
		if indices, ok := m[diff]; ok {
			for _, j := range indices {
				result = append(result, []int{j, i})
			}
		}
		m[num] = append(m[num], i)
	}
	return result
}
