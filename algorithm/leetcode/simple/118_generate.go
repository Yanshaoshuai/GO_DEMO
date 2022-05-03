package main

import "fmt"

//a1=1
//a2=1,1
//a3=1,2,1
//an(1)=1 an(2)=an-1(1)+an-1(2)
func generate(numRows int) [][]int {
	var result [][]int
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1, i+1)
		result = append(result, row)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				row[j] = 1
				continue
			}
			row[j] = result[i-1][j-1] + result[i-1][j]
		}
	}
	return result
}

func main() {
	result := generate(100)
	for index, row := range result {
		fmt.Printf("row %d--%v \n", index, row)
	}
}
