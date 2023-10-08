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

	map1 := make(map[int][]int)

	for i := 0; i < n-1; i++ {
		var value1, value2 int
		fmt.Fscan(reader, &value1, &value2)

		if _, ok := map1[value1]; !ok {
			map1[value1] = []int{value2}
		} else {
			map1[value1] = append(map1[value1], value2)
		}

		if _, ok := map1[value2]; !ok {
			map1[value2] = []int{value1}
		} else {
			map1[value2] = append(map1[value2], value1)
		}
	}

	fmt.Println(length(map1))
}

func length(map1 map[int][]int) int {
	var ans int = 0

	for len(map1) >= 2 {
		singles := singlebead(map1)
		for _, bead := range singles {
			for _, batya := range map1[bead] {
				map1[batya] = deletebead(map1[batya], bead)
			}
			delete(map1, bead)
		}

		if len(singles) < 2 {
			ans++
		} else {
			ans += 2
		}
	}

	ans += len(map1)
	return ans
}

func singlebead(map1 map[int][]int) []int {
	anslist := []int{}
	for key := range map1 {
		if len(map1[key]) == 1 {
			anslist = append(anslist, key)
		}
	}
	return anslist
}

func deletebead(beads []int, value int) []int {
	index := -1
	for ind, elem := range beads {
		if elem == value {
			index = ind
			break
		}
	}
	if index == -1 {
		return beads
	}
	return append(beads[:index], beads[index+1:]...)
}
