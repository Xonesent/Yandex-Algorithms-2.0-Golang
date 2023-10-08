package main

import "fmt"

func main() {
	var shopplace, maxdistance int = 0, 0
	var value int
	distances := make(map[int][2]int)
	homes := make(map[int]bool)
	for i := 1; i <= 10; i++ {
		fmt.Scan(&value)
		if value == 2 {
			shopplace = i
			for key := range homes {
				distances[key] = [2]int{distances[key][0], shopplace - key}
				delete(homes, key)
			}
		} else if value == 1 {
			homes[i] = true
			distances[i] = [2]int{0, 0}
			if shopplace != 0 {
				distances[i] = [2]int{i - shopplace, distances[i][1]}
			}
		}
	}
	for key := range distances {
		if distances[key][0] != 0 && distances[key][1] != 0 {
			if min(distances[key][0], distances[key][1]) > maxdistance {
				maxdistance = min(distances[key][0], distances[key][1])
			}
		} else {
			if max(distances[key][0], distances[key][1]) > maxdistance {
				maxdistance = max(distances[key][0], distances[key][1])
			}
		}
	}
	fmt.Println(maxdistance)
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
