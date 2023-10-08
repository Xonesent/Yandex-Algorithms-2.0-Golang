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

	events := make([][2]int, 2*n)
	for i := 0; i < 2*n; i++ {
		var left, right int
		fmt.Fscan(reader, &left, &right)
		events[i] = [2]int{left, 1}
		i++
		events[i] = [2]int{left + right, 0}
	}

	sort.SliceStable(events, func(i, j int) bool {
		if events[i][0] != events[j][0] {
			return events[i][0] < events[j][0]
		}
		return events[i][1] < events[j][1]
	})

	var ans, max int = 0, 0
	for i := 0; i < 2*n; i++ {
		if events[i][1] == 1 {
			ans++
			if ans > max {
				max = ans
			}
		}
		if events[i][1] == 0 {
			ans--
		}
	}

	fmt.Println(max)
}
