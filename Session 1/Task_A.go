package main

import "fmt"

func main() {
	var r, i, c int
	fmt.Scan(&r, &i, &c)
	switch i {
	case 0:
		switch {
		case r == 0:
			fmt.Println(c)
		case -128 <= r && r <= 127 && r != 0:
			fmt.Println(3)
		default:
			fmt.Println(i)
		}
	case 1:
		fmt.Println(c)
	case 4:
		switch {
		case r == 0:
			fmt.Println(4)
		case -128 <= r && r <= 127 && r != 0:
			fmt.Println(3)
		default:
			fmt.Println(i)
		}
	case 6:
		fmt.Println(0)
	case 7:
		fmt.Println(1)
	default:
		fmt.Println(i)
	}
}
