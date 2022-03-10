package main

import "fmt"

func main() {
	var arr [10]int
	var slice = arr[5:6]
	fmt.Printf("len(slice)=%d,cap(slice)=%d\n", len(slice), cap(slice))

	//share memory test
	s1 := []int{1, 2}
	s2 := s1           //s2 cap 2 len  2
	s2 = append(s2, 3) //s1 cap 2 len 2 s2 cap 4 len 3
	SliceRise(s1)
	SliceRise(s2)
	fmt.Println(s1, s2)

	SliceExtend()
}
func SliceRise(s []int) {
	s = append(s, 0) //s1 cap 2 len 2 s2 cap 4 len 3
	for i := range s {
		s[i]++
	}
}
func SliceExtend() {
	s := make([]int, 0, 10)
	s1 := append(s, 1, 2, 3)
	s2 := append(s1, 4)
	println(&s1[0] == &s2[0])
	//println(s[0])// index out of range [0] with length 0
}
func SliceExpress() {
	orderLen := 5
	order := make([]uint16, orderLen*2)               //cap 10 len 10
	pollorder := order[:orderLen:orderLen]            //low 0 high 5 max 5 cap=max-low=5 len=5 share [order[0],order[4]]
	lockorder := order[orderLen:][:orderLen:orderLen] //1.low 5 high 10
	//2.low 0 high 5 max 5=>cap=5 len=5 share [order[5],order[9]]
	fmt.Printf("len(pollorder)=%d\n", len(pollorder)) //5
	fmt.Printf("cap(pollorder)=%d\n", cap(pollorder)) //5
	fmt.Printf("len(lockorder)=%d\n", len(lockorder)) //5
	fmt.Printf("cap(lockorder)=%d\n", cap(lockorder)) //5
}
