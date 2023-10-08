package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, q int
	fmt.Fscan(reader, &n, &q)

	ansmap := make(map[int]int)
	for i := 1; i <= n; i++ {
		var value int
		fmt.Fscan(reader, &value)
		if i == 0 {
			ansmap[i] = value
		} else {
			ansmap[i] = value + ansmap[i-1]
		}
	}
	for i := 0; i < q; i++ {
		var val1, val2 int
		fmt.Fscan(reader, &val1, &val2)
		fmt.Println(ansmap[val2] - ansmap[val1-1])
	}
}
