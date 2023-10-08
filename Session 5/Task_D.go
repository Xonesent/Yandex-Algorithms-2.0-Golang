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
	line = strings.TrimRight(line, "\n")
	var list []rune
	var flag bool = true

	for _, value := range line {
		if value == '(' {
			list = append(list, value)
		} else {
			if len(list) != 0 {
				list = append(list[:len(list)-1], list[len(list):]...)
			} else {
				flag = false
				break
			}
		}
	}

	if flag && len(list) == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
