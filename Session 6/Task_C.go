package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b, c, d float64
	fmt.Fscan(reader, &a, &b, &c, &d)

	var left, right float64 = -1000000, 1000000
	var point float64 = 1
	if a < 0 {
		point = -1
	}

	for left+0.000001 < right {
		mid := (left + right) / 2
		if point*check(a, b, c, d, mid) < 0 {
			left = mid + 0.000001
		} else {
			right = mid
		}
		fmt.Println(left, right)
	}

	fmt.Printf("%.8f", left)
}

func check(a, b, c, d, num float64) float64 {
	return a*math.Pow(num, 3) + b*math.Pow(num, 2) + c*num + d
}
