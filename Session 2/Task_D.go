package main

import "fmt"

func main() {
	var length, quantity, value, tmp_value int
	var flag bool = true
	fmt.Scan(&length, &quantity)
	if length%2 != 0 {
		flag = false
	}
	if flag {
		for i := 0; i < quantity && value <= length/2-1; i++ {
			tmp_value = value
			fmt.Scan(&value)
		}
		fmt.Println(tmp_value, value)
	} else {
		for i := 0; i < quantity && value < length/2; i++ {
			tmp_value = value
			fmt.Scan(&value)
		}
		if value == length/2 {
			fmt.Println(value)
		} else {
			fmt.Println(tmp_value, value)
		}
	}
}
