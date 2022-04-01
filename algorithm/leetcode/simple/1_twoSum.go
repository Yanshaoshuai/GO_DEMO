package main

func twoSum(nums []int, target int) []int {
	result := make([]int, 2)
	m := map[int]int{}
	for index, num := range nums {
		m[num] = index
	}
	for index, num := range nums {
		anotherValue := target - num
		if anotherIndex, ok := m[anotherValue]; ok && anotherIndex != index {
			result[0] = index
			result[1] = anotherIndex
			return result
		}

	}
	return result
}
