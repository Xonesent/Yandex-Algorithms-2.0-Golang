package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	if (x >= 1 && x <= 12) && (y >= 1 && y <= 12) && (x != y) {
		fmt.Println(0)
	} else {
		fmt.Println(1)
	}
}
