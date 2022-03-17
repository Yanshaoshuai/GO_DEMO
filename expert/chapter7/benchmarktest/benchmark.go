package benchmarktest

func MakeSliceWithoutAlloc() []int {
	var newSlice []int
	for i := 0; i < 100000; i++ {
		newSlice = append(newSlice, 1)
	}
	return newSlice
}

func MakeSliceWithPreAlloc() []int {
	var newSlice = make([]int, 100000)
	for i := 0; i < 100000; i++ {
		newSlice = append(newSlice, 1)
	}
	return newSlice
}
