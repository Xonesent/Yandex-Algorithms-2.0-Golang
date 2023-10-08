package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var m int
	fmt.Fscan(reader, &m)

	events := make([][3]int, 0)
	for {
		var left, right int
		fmt.Fscan(reader, &left, &right)
		if left == 0 && right == 0 {
			break
		}
		if left >= 0 && left < m || right <= m && right > 0 {
			events = append(events, [3]int{left, 0, right})
		}
	}
	events = append(events, [3]int{0, 1, 0})

	supersort(events)

	check := true
	map1 := make(map[[3]int]bool)
	anslist := make([][2]int, 0)
	for i := 0; i < len(events); i++ {
		if events[i][1] == 0 {
			map1[events[i]] = true
		} else if events[i][1] == 1 {
			if len(map1) == 0 {
				check = false
				break
			} else {
				var max, min int
				for key := range map1 {
					if key[2] > max {
						max = key[2]
						min = key[0]
					}
				}
				anslist = append(anslist, [2]int{min, max})
				if max >= m {
					break
				}
				events = append(events, [3]int{max, 1, 0})
				supersort(events)
				map1 = make(map[[3]int]bool)
			}
		}
	}

	if check {
		fmt.Println(len(anslist))
		for _, value := range anslist {
			fmt.Println(value[0], value[1])
		}
	} else {
		fmt.Println("No solution")
	}
}

func supersort(events [][3]int) {
	sort.SliceStable(events, func(i, j int) bool {
		if events[i][0] != events[j][0] {
			return events[i][0] < events[j][0]
		}
		return events[i][1] < events[j][1]
	})
}
