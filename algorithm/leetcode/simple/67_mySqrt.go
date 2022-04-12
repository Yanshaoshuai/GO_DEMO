package main

//func mySqrt(x int) int {
//	i := 0
//	for ; i <= x/2+1; i++ {
//		muti := i * i
//		if muti == x {
//			return i
//		}
//		if i*i > x {
//			return i - 1
//		}
//	}
//	return 0
//}
//二分查找
func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	var min, max = 0, x
	for max-min > 1 {
		midle := (min + max) / 2
		if midle*midle > x {
			max = midle
		} else {
			min = midle
		}
	}
	return min
}
