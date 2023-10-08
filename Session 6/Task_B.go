package main

import (
	"bufio"
	"fmt"
	"os"
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

	var k int
	fmt.Fscan(reader, &k)

	for i := 0; i < k; i++ {
		var num int
		fmt.Fscan(reader, &num)

		var indexleft, indexright int = 0, len(slice) - 1

		var l, r int = 0, len(slice) - 1
		for l < r {
			mid := (l + r) / 2
			if slice[mid] < num {
				l = mid + 1
			} else {
				r = mid
			}
		}
		indexleft = l

		l, r = 0, len(slice)-1
		for l < r {
			mid := (l + r + 1) / 2
			if slice[mid] > num {
				r = mid - 1
			} else {
				l = mid
			}
		}
		indexright = l

		if num < slice[0] {
			fmt.Println(0, 0)
		} else if num > slice[len(slice)-1] {
			fmt.Println(0, 0)
		} else if indexright < indexleft {
			fmt.Println(0, 0)
		} else {
			fmt.Println(indexleft+1, indexright+1)
		}
	}
}
