package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)

	slice := make([]int, n)
	for i := 0; i < n; i++ {
		var value int
		fmt.Fscan(reader, &value)
		slice[i] = value
	}

	sort.Ints(slice)

	var k int
	fmt.Fscan(reader, &k)

	for i := 0; i < k; i++ {
		var left, right int
		fmt.Fscan(reader, &left, &right)

		var indexleft, indexright int = 0, len(slice) - 1

		var l, r int = 0, len(slice) - 1
		for l < r {
			mid := (l + r) / 2
			if slice[mid] < left {
				l = mid + 1
			} else {
				r = mid
			}
		}
		indexleft = l

		l, r = 0, len(slice)-1
		for l < r {
			mid := (l + r + 1) / 2
			if slice[mid] > right {
				r = mid - 1
			} else {
				l = mid
			}
		}
		indexright = l

		if right < slice[0] {
			fmt.Print(0, " ")
		} else if left > slice[len(slice)-1] {
			fmt.Print(0, " ")
		} else if indexright < indexleft {
			fmt.Print(0, " ")
		} else {
			fmt.Print(indexright-indexleft+1, " ")
		}
	}
}
