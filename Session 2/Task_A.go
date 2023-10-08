package main

import "fmt"

func main() {
	var num, max, count int = 0, 0, 0
	fmt.Scan(&num)
	if num != 0 {
		max = num
		count = 1
		for {
			if num == 0 {
				break
			}
			fmt.Scan(&num)
			if num > max {
				max = num
				count = 1
			} else if num == max {
				count++
			}
		}
	}
	fmt.Println(count)
}
