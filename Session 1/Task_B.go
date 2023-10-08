package main

import "fmt"

func main() {
	var n, i, j int
	fmt.Scan(&n, &i, &j)
	anscase_1 := (max(i, j) - min(i, j) - 1)
	anscase_2 := n - anscase_1 - 2
	fmt.Println(min(anscase_1, anscase_2))
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
