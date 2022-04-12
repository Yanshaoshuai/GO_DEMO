package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	firstArr := make([]int, m)
	for k := 0; k < m; k++ {
		firstArr[k] = nums1[k]
	}
	var i, j, l = 0, 0, 0
	for ; i < m && j < n; l++ {
		if firstArr[i] < nums2[j] {
			nums1[l] = firstArr[i]
			i++
		} else {
			nums1[l] = nums2[j]
			j++
		}
	}
	for ; i < m; l, i = l+1, i+1 {
		nums1[l] = firstArr[i]
	}
	for ; j < n; l, j = l+1, j+1 {
		nums1[l] = nums2[j]
	}
}

func main() {
	var nums1 = []int{1, 2, 3, 0, 0, 0}
	var nums2 = []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
}
