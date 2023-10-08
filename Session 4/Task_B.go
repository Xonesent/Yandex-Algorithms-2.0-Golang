package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	ansmap := make(map[string]int)
	anslist := make([]string, 0)

	for {
		var person string
		var votes int
		rep, _ := fmt.Fscan(reader, &person, &votes)
		if rep == 2 {
			if _, ok := ansmap[person]; !ok {
				anslist = append(anslist, person)
			}
			ansmap[person] += votes
		} else {
			break
		}
	}

	sort.Strings(anslist)
	for _, value := range anslist {
		fmt.Printf("%s %d\n", value, ansmap[value])
	}
}
