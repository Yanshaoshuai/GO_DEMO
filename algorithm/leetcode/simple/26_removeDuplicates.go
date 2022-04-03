package main

import "fmt"

func removeDuplicates(nums []int) int {
	//i := 0
	//for i < len(nums)-1 {
	//	if nums[i] != nums[i+1] {
	//		if nums[i] > nums[i+1] {
	//			return i + 1
	//		}
	//		i++
	//	} else {
	//		j := i
	//		for ; j < len(nums)-1; j++ {
	//			nums[j] = nums[j+1]
	//			nums[j+1] = nums[0] - 1
	//		}
	//	}
	//}
	//if i != 0 {
	//	return i + 1
	//}
	//return len(nums)

	//快慢指针
	if len(nums) < 2 {
		return len(nums)
	}
	//慢指针之前都是不重复的
	j := 0
	for i := 1; i < len(nums); i++ {
		//如果遇到不重复的则添加不重复范围
		if nums[i] != nums[j] {
			j += 1
			nums[j] = nums[i]
		}
	}
	return j + 1
}

func aFun(a []int) {
	a = []int{1, 2, 3}
}
func main() {
	a := []int{1, 23}
	aFun(a)
	fmt.Printf("%v", a)
	removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
}
