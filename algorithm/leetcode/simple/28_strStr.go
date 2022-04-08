package main

//todo KMP
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) == 0 {
		return -1
	}
	var i, j, k int
	for i < len(haystack) && j < len(needle) {
		if haystack[i] == needle[j] {
			if j == 0 {
				k = i
			}
			i++
			j++
		} else {
			i = k
			i++
			k++
			j = 0
		}
	}
	if j == len(needle) {
		return i - len(needle)
	}
	return -1
}
