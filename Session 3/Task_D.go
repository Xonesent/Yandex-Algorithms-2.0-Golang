package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanf(reader, "%d\n", &n)

	notans := make(map[int]bool)
	ans := make(map[int]bool)

	for {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "HELP" {
			break
		}
		nums := parse(line)
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)

		if answer == "NO" {
			for _, num := range nums {
				notans[num] = true
			}
		} else if answer == "YES" {
			if len(ans) == 0 {
				for _, num := range nums {
					ans[num] = true
				}
			} else {
				temp := make(map[int]bool)
				for _, num := range nums {
					if ans[num] {
						temp[num] = true
					}
				}
				ans = temp
			}
		}
	}

	for i := 1; i <= n; i++ {
		if !notans[i] && (len(ans) == 0 || ans[i]) {
			fmt.Printf("%d ", i)
		}
	}

}

func parse(line string) []int {
	numStrs := strings.Fields(line)
	nums := make([]int, len(numStrs))

	for i, numStr := range numStrs {
		num, _ := strconv.Atoi(numStr)
		nums[i] = num
	}

	return nums
}
