package main

import (
	"fmt"
	"math"
)

func main() {
	var d, x, y int
	fmt.Scan(&d, &x, &y)
	var check bool = false
	for z := 0; z <= d && check == false; z++ {
		if x >= 0 && x <= d-z && y >= 0 && y <= z {
			check = true
		}
	}
	if check {
		fmt.Println(0)
	} else {
		distance1 := math.Sqrt(math.Pow(float64(x), 2) + (math.Pow(float64(y), 2)))
		distance2 := math.Sqrt(math.Pow(float64(x-d), 2) + (math.Pow(float64(y), 2)))
		distance3 := math.Sqrt(math.Pow(float64(x), 2) + (math.Pow(float64(y-d), 2)))
		fmt.Println(findmin(distance1, distance2, distance3))
	}
}

func findmin(a float64, b float64, c float64) int {
	if c < a && c < b {
		return 3
	} else if b < a && b <= c {
		return 2
	} else {
		return 1
	}
}
