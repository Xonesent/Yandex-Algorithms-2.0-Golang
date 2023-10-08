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

	line, _ := reader.ReadString('\n')
	linelist := strings.Fields(line)

	num_to_cnt := make(map[int]int)
	var nums []int

	for i := 0; i < len(linelist); i++ {
		num, _ := strconv.Atoi(linelist[i])
		num_to_cnt[num]++
		nums = append(nums, num)
	}

	for _, num := range nums {
		if num_to_cnt[num] == 1 {
			fmt.Printf("%d ", num)
		}
	}
}
