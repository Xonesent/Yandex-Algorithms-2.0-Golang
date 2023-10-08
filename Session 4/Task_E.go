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

	topics := make(map[string]int)
	stroki := make(map[int]string)

	for i := 1; i <= n; i++ {
		line1, _ := reader.ReadString('\n')
		line1 = strings.TrimRight(line1, "\n")
		num, _ := strconv.Atoi(line1)
		if num == 0 {
			line2, _ := reader.ReadString('\n')
			line2 = strings.TrimRight(line2, "\n")
			topics[line2] += 1
			stroki[i] = line2
		} else {
			topics[stroki[num]] += 1
			stroki[i] = stroki[num]
		}
		line3, _ := reader.ReadString('\n')
		line3 = strings.TrimRight(line3, "\n")
	}

	maxTopic := ""
	maxCount := 0

	for topic, count := range topics {
		if count > maxCount {
			maxCount = count
			maxTopic = topic
		}
	}

	fmt.Println(maxTopic)
}
