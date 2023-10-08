package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	ansmap := make(map[int]int)
	anslist := make([]int, 0)

	for i := 0; i < n; i++ {
		var color, value int
		fmt.Scan(&color, &value)
		if _, ok := ansmap[color]; !ok {
			anslist = append(anslist, color)
		}
		ansmap[color] += value
	}

	sort.Ints(anslist)

	for _, value := range anslist {
		fmt.Printf("%d %d\n", value, ansmap[value])
	}
}
