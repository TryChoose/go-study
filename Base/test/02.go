package main

import "fmt"

//托普利兹矩阵

func main() {
	matric := [...][4]int{
		{1, 2, 3, 4},
		{5, 1, 2, 3},
		{9, 1, 1, 2},
	}
	var flag bool
	flag = true
	for i := 0; i < len(matric)-1; i++ {
		for j := 0; j < len(matric[0])-1; j++ {
			if matric[i][j] != matric[i+1][j+1] {
				flag = false
				break
			}
		}
	}
	fmt.Println(flag)
}
