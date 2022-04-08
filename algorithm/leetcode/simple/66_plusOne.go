package main

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+1 < 10 {
			digits[i] = digits[i] + 1
			return digits
		} else {
			digits[i] = 0
		}
	}
	digits[0] = 1
	digits = append(digits, 0)
	return digits
}
