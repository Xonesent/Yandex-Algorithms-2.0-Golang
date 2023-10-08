package main

import "fmt"

func main() {
	var count int
	var word string
	fmt.Scan(&word)
	for i := 0; i <= len(word)/2-1; i++ {
		if word[i] != word[len(word)-i-1] {
			count++
		}
	}
	fmt.Println(count)
}
