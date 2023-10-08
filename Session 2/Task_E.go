package main

import "fmt"

func main() {
	var quantity int
	fmt.Scan(&quantity)
	list := make([]int, 0)
	for i := 0; i < quantity; i++ {
		var value int
		fmt.Scan(&value)
		list = append(list, value)
	}
	qsort(list, 0, len(list)-1)
	var count int
	for i := 0; i < quantity-1; i++ {
		count += list[i]
	}
	fmt.Println(count)
}

func qsort(arr []int, start int, end int) {
	if start < end {
		var left, right, middle int
		left = start
		right = end
		middle = arr[(left+right)/2]
		for left < right {
			for arr[left] < middle {
				left++
			}
			for arr[right] > middle {
				right--
			}
			if left <= right {
				arr[left], arr[right] = arr[right], arr[left]
				left++
				right--
			}
		}
		qsort(arr, start, right)
		qsort(arr, left, end)
	}
}
