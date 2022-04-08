package main

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
