package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(reader, &n, &k)
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		var value int
		fmt.Fscan(reader, &value)
		slice[i] = value
	}

	sort.Ints(slice)

	ostatki := slice[0] - 1
	for i := 0; i < n; i++ {
		slice[i] -= ostatki
	}

	left := 0
	right := slice[len(slice)-1]

	for left < right {
		mid := (left + right) / 2
		if check(slice, mid) <= k {
			right = mid
		} else {
			left = mid + 1
		}
	}
	fmt.Println(left)
}

func check(arr []int, length int) int {
	var count int = 1
	curr := arr[0]
	for _, value := range arr {
		if curr+length < value {
			curr = value
			count++
		}
	}
	return count
}
