package my_idea

import "math"

//乱序数组 找出符合i<j 且nums[j]-nums[i]或nums[j]最大的数 返回差值 其中如果
//-1,-2,20,-20,4 => 24
func maxDif(nums []int) int {
	min, max := math.MaxInt, math.MinInt
	for i := 0; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
		if nums[i]-min > max {
			max = nums[i] - min
		}
	}
	return max
}
func main() {
}
