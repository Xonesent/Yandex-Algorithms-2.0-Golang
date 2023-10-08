package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type words struct {
	word  string
	count int
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	ansmap := make(map[string]int)
	anslist := make([]words, 0)

	for {
		var word string
		rep, _ := fmt.Fscan(reader, &word)
		if rep != 1 {
			break
		}
		word = strings.TrimRight(word, ".,!?")
		word = strings.ToLower(word)
		ansmap[word]++
	}

	for word, freq := range ansmap {
		anslist = append(anslist, words{word: word, count: freq})
	}

	sort.Slice(anslist, func(i, j int) bool {
		if anslist[i].count == anslist[j].count {
			return anslist[i].word < anslist[j].word
		}
		return anslist[i].count > anslist[j].count
	})

	for _, value := range anslist {
		fmt.Println(value.word)
	}
}
