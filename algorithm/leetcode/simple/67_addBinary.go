package main

func addBinary(a string, b string) string {
	l := len(a) + 1
	if len(b)+1 > l {
		l = len(b) + 1
	}
	bytes := make([]byte, l)
	var enter byte = 0
	i, j, k := len(a)-1, len(b)-1, len(bytes)-1
	var temp byte
	for ; k >= 0; k-- {
		temp = 0
		if i >= 0 {
			temp += a[i] - 48
			i--
		}
		if j >= 0 {
			temp += b[j] - 48
			j--
		}
		temp += enter
		enter = 0
		if temp >= 2 {
			bytes[k] = 48 + temp%2
			enter = 1
		} else {
			bytes[k] = 48 + temp
		}
	}
	if bytes[0] == 48 {
		bytes = bytes[1:]
	}
	return string(bytes)
}

func main() {
	s := addBinary("0", "0")
	println(s)
	bytes := make([]byte, 3)
	println(0 == bytes[0])
}
