package main

import "math"

//func maxSubArray(nums []int) int {
//	length := len(nums)
//	if length == 1 {
//		return nums[0]
//	}
//	var sumArr = make([]int, length)
//	var sum int
//	for i := 0; i < length; i++ {
//		sum += nums[i]
//		sumArr[i] = sum
//	}
//
//	return maxDif(sumArr)
//}
//func maxDif(nums []int) int {
//	minIndex := 0
//	max := nums[0]
//	for i := 1; i < len(nums); i++ {
//		temp := nums[i] - nums[minIndex]
//		if temp > max || nums[i] > max {
//			if temp > nums[i] {
//				max = temp
//			} else {
//				max = nums[i]
//			}
//		}
//		if nums[i] < nums[minIndex] {
//			minIndex = i
//		}
//
//	}
//	return max
//}

//分治法
//最大和子数组=>Smax(n)
//Smax(n-1)<0=>Smax(n)=max(Smax(n-1),An)
//Smax(n-1)>0=>Smax(n)=max(Smax(n-1),An+Smax(n-1))
func maxSubArray(nums []int) int {
	maxBefore := math.MinInt
	var sum int
	for i := 0; i < len(nums); i++ {
		if sum < 0 {
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		if maxBefore < sum {
			maxBefore = sum
		}
	}

	return maxBefore
}
