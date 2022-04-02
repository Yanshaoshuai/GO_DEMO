package main

func longestCommonPrefix(strs []string) string {
	l := len(strs)
	minIndex := 0
	for i := 1; i < l; i++ {
		if len(strs[minIndex]) > len(strs[i]) {
			minIndex = i
		}
	}
	var contains bool
	for len(strs[minIndex]) != 0 {
		for i := 0; i < l; i++ {
			contains = strs[minIndex] == strs[i][:len(strs[minIndex])]
			if !contains {
				break
			}
		}
		if !contains {
			strs[minIndex] = strs[minIndex][:len(strs[minIndex])-1]
		} else {
			return strs[minIndex]
		}
		contains = false
	}
	return ""
}
