package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(reader, &n, &m)

	events := make([][3]int, 0)
	anslist := make([][3]int, m)

	for i := 0; i < n; i++ {
		var avalue int
		fmt.Fscan(reader, &avalue)
		events = append(events, [3]int{avalue, 0, -1})
	}

	for i := 0; i < m; i++ {
		var left, right int
		fmt.Scan(&left, &right)
		anslist[i] = [3]int{left, right, 0}
		events = append(events, [3]int{left, -1, i})
		events = append(events, [3]int{right, 1, i})
	}

	sort.SliceStable(events, func(i, j int) bool {
		if events[i][0] != events[j][0] {
			return events[i][0] < events[j][0]
		}
		return events[i][1] < events[j][1]
	})

	count := 0
	for _, event := range events {
		if event[1] == -1 {
			anslist[event[2]][2] = count
		} else if event[1] == 1 {
			anslist[event[2]][2] = count - anslist[event[2]][2]
		} else {
			count++
		}
	}

	for i := 0; i < len(anslist); i++ {
		fmt.Println(anslist[i][2])
	}
}
