package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	line, _ := reader.ReadString('\n')
	nums := make(map[int]bool)

	scanner := bufio.NewScanner(strings.NewReader(line))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		var num int
		fmt.Sscanf(scanner.Text(), "%d", &num)
		if nums[num] {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
			nums[num] = true
		}
	}
}
