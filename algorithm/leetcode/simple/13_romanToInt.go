package main

func romanToInt(s string) int {
	//specMap := map[string]int{"IV": 4, "IX": 9, "XL": 40, "XC": 90, "CD": 400, "CM": 900}
	//if _, ok := specMap[s]; ok {
	//	return specMap[s]
	//}
	//var result int
	//for key, value := range specMap {
	//	if strings.Contains(s, key) {
	//		s = strings.Replace(s, key, "", 1)
	//		result += value
	//	}
	//}
	//m := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	//for _, c := range s {
	//	result += m[c]
	//}
	//return result
	m := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	var result int
	i := 1
	for ; i < len(s); i++ {
		if m[s[i-1]] < m[s[i]] {
			result -= m[s[i-1]]
		} else {
			result += m[s[i-1]]
		}
	}
	result += m[s[i-1]]
	return result
}
