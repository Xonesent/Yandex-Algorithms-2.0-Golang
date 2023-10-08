package main

import "fmt"

func main() {
	var n, value int
	fmt.Scan(&n)
	for i := 0; i <= n/2; i++ {
		fmt.Scan(&value)
	}
	fmt.Println(value)
}
