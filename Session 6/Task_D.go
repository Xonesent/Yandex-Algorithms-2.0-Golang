package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var a, k, b, m, x int
	fmt.Fscan(reader, &a, &k, &b, &m, &x)
	var left int = 0
	var right int = x
	for left < right {
		mid := (left + right) / 2
		if check(mid, a, k, b, m, x) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	fmt.Println(left)
}

func check(time, a, k, b, m, x int) bool {
	res1 := (time - time/k) * a
	res2 := (time - time/m) * b
	return res1+res2 >= x
}
