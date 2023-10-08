package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)

	max := math.MinInt64
	currmax := 0
	for i := 1; i <= n; i++ {
		if currmax < 0 {
			currmax = 0
		}
		var value int
		fmt.Fscan(reader, &value)
		currmax += value
		if currmax > max {
			max = currmax
		}
	}
	fmt.Println(max)
}
