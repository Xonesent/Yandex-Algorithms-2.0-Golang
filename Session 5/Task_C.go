package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(reader, &n, &m)

	people := make([][2]int, n)
	rooms := make([][2]int, m)
	for i := 0; i < n; i++ {
		var value int
		fmt.Fscan(reader, &value)
		people[i] = [2]int{value + 1, i}
	}
	for i := 0; i < m; i++ {
		var value int
		fmt.Fscan(reader, &value)
		rooms[i] = [2]int{value, i}
	}

	sort.Slice(people, func(i, j int) bool {
		return people[i][0] > people[j][0]
	})
	sort.Slice(rooms, func(i, j int) bool {
		return rooms[i][0] > rooms[j][0]
	})

	peoplepointer := 0
	roompointer := 0
	anslist := make([]string, n)
	ans := 0

	for peoplepointer < n && roompointer < m {
		if people[peoplepointer][0] <= rooms[roompointer][0] {
			anslist[people[peoplepointer][1]] = strconv.Itoa(rooms[roompointer][1] + 1)
			peoplepointer++
			roompointer++
			ans++
		} else {
			anslist[people[peoplepointer][1]] = "0"
			peoplepointer++
		}
	}

	fmt.Println(ans)
	for _, value := range anslist {
		fmt.Print(value, " ")
	}
}
