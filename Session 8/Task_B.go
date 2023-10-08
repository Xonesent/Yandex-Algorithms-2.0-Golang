package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)

	people := make(map[string]string, n-1)
	for i := 0; i < n; i++ {
		var child, predok string
		fmt.Fscanf(reader, "%s %s\n", &child, &predok)
		people[child] = predok
	}

	for {
		line, _ := reader.ReadString('\n')
		line = strings.TrimRight(line, "\n")

		if line == "" || line == "\n" {
			break
		}

		params := strings.Fields(line)
		person1 := params[0]
		person2 := params[1]

		if findparent(person1, person2, people) {
			fmt.Print(2, " ")
		} else if findparent(person2, person1, people) {
			fmt.Print(1, " ")
		} else {
			fmt.Print(0, " ")
		}
	}
}

func findparent(person string, compare string, people map[string]string) bool {
	var check bool = false
	for people[person] != "" && !check {
		if people[person] == compare {
			check = true
		}
		person = people[person]
	}
	return check
}
