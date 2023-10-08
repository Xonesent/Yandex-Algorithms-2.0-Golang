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

		fmt.Println(findparent(person1, person2, people))
	}
}

func findparent(person1 string, person2 string, people map[string]string) string {
	ansmap := make(map[string]bool)
	ansmap[person1] = true
	ansmap[person2] = true
	for people[person1] != "" && people[person2] != "" {
		if _, ok1 := ansmap[people[person1]]; !ok1 {
			ansmap[people[person1]] = true
		} else {
			return people[person1]
		}
		if _, ok2 := ansmap[people[person2]]; !ok2 {
			ansmap[people[person2]] = true
		} else {
			return people[person2]
		}
		person1 = people[person1]
		person2 = people[person2]
	}

	if people[person1] == "" {
		for _, ok := ansmap[people[person2]]; !ok && people[person2] != ""; _, ok = ansmap[people[person2]] {
			person2 = people[person2]
		}
		return people[person2]
	} else {
		for _, ok := ansmap[people[person1]]; !ok && people[person1] != ""; _, ok = ansmap[people[person1]] {
			person1 = people[person1]
		}
		return people[person1]
	}
}
