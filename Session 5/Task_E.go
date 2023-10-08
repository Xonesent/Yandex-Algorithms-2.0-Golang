package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	matrix := make([][]int, 3)
	ans := make([]int, 3)
	var s int
	fmt.Fscan(reader, &s)

	for i := 0; i < 3; i++ {
		var q int
		fmt.Fscan(reader, &q)
		matrix[i] = make([]int, q)
		for ind := 0; ind < q; ind++ {
			var value int
			fmt.Fscan(reader, &value)
			matrix[i][ind] = value
		}
	}

	map3 := make(map[int]int)
	for index := 0; index < len(matrix[2]); index++ {
		if _, ok := map3[matrix[2][index]]; !ok {
			map3[matrix[2][index]] = index
		}
	}

	var flag bool = true
	for i := 0; i < len(matrix[0]) && flag; i++ {
		for ind := 0; ind < len(matrix[1]) && flag; ind++ {
			if _, ok := map3[s-matrix[0][i]-matrix[1][ind]]; ok {
				ans = []int{i, ind, map3[s-matrix[0][i]-matrix[1][ind]]}
				flag = false
			}
		}
	}

	if !flag {
		for i := 0; i < len(ans); i++ {
			fmt.Print(ans[i], " ")
		}
	} else {
		fmt.Println(-1)
	}
}
