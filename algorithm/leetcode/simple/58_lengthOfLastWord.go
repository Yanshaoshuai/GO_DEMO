package main

func lengthOfLastWord(s string) int {
	var sum int
	var flag bool
	for i := len(s) - 1; i >= 0; i-- {
		if (65 <= s[i] && 90 >= s[i]) || (97 <= s[i] && 122 >= s[i]) {
			flag = true
			sum++
		} else {
			if flag {
				return sum
			}
		}
	}
	return sum
}
