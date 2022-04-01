package main

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	////双指针
	//b := []byte(strconv.Itoa(x))
	//var l, r int
	//len := len(b)
	//for ; l < len; l++ {
	//	r = len - 1 - l
	//	if l >= r {
	//		return true
	//	}
	//	if b[l] != b[r] {
	//		return false
	//	}
	//}
	//return false

	//左乘右除
	var l, r int
	r = x
	for r != 0 {
		l = l*10 + r%10
		r = r / 10
	}
	return x == l

}
