package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var m int
	fmt.Fscan(reader, &m)
	checkers1 := make([]string, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &checkers1[i])
	}

	var n int
	fmt.Fscan(reader, &n)
	checkers2 := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &checkers2[i])
	}

	max := 0
	anslist := []string{}

	for _, suspect := range checkers2 {
		count := 0
		for _, witness := range checkers1 {
			if match(witness, suspect) {
				count++
			}
		}

		if count > max {
			max = count
			anslist = []string{suspect}
		} else if count == max {
			anslist = append(anslist, suspect)
		}
	}

	for _, ans := range anslist {
		fmt.Println(ans)
	}
}

func match(n, m string) bool {
	for _, char := range n {
		if strings.IndexRune(m, char) == -1 {
			return false
		}
	}
	return true
}
