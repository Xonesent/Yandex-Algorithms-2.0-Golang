package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	line1, _ := reader.ReadString('\n')
	line2, _ := reader.ReadString('\n')
	str1 := makemap(strings.TrimSpace(line1))
	str2 := makemap(strings.TrimSpace(line2))

	result := count(str1, str2)
	fmt.Println(result)
}

func count(map1, map2 map[int]bool) int {
	ans := 0
	for key := range map1 {
		if _, ok := map2[key]; ok {
			ans++
		}
	}
	return ans
}

func makemap(line string) map[int]bool {
	nums := strings.Fields(line)
	return_map := make(map[int]bool)

	for _, value := range nums {
		var num int
		fmt.Sscanf(value, "%d", &num)
		return_map[num] = true
	}

	return return_map
}
