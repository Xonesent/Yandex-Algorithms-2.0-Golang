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
		events[i] = [2]int{right, 0}
	}

	sort.SliceStable(events, func(i, j int) bool { return events[i][0] < events[j][0] })

	var ans int = 0
	var flag int = 0
	var start, end int
	for i := 0; i < 2*n; i++ {
		if events[i][1] == 1 && flag == 0 {
			start = events[i][0]
			flag = 1
		} else if events[i][1] == 1 {
			flag++
		}
		if events[i][1] == 0 && flag == 1 {
			end = events[i][0]
			flag = 0
			ans += end - start
		} else if events[i][1] == 0 {
			flag--
		}
	}

	fmt.Println(ans)
}
