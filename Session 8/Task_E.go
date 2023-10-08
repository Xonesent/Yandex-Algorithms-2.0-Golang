package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
	left   *node
	right  *node
	predok *node
	check  int
	value  string
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)

	lines := []string{}

	for i := 0; i < n; i++ {
		var line string
		fmt.Fscan(reader, &line)
		lines = append(lines, line)
	}

	for _, value := range lines {
		var curr string = ""
		var flag bool = false

		tree := &node{nil, nil, nil, 0, ""}
		root := tree
		current := root

		for _, letter := range value {
			if letter == 'D' {
				curr = fmt.Sprint(curr + "0")
				elem := &node{nil, nil, current, -1, curr}
				current.left = elem
				current = elem
				flag = false
			} else if letter == 'U' {
				if flag {
					curr = curr[:len(curr)-2]
				} else {
					curr = curr[:len(curr)-1]
				}
				curr = fmt.Sprint(curr + "1")
				for current.check == 1 {
					current = current.predok
				}
				current = current.predok
				elem := &node{nil, nil, current, 1, curr}
				current.right = elem
				current = elem
				flag = true
			}
		}
		anslist := make([]string, 0)
		proxod(tree, &anslist)

		fmt.Println(len(anslist))
		for _, value := range anslist {
			fmt.Println(value)
		}
	}
}

func proxod(root *node, anslist *[]string) {
	if root != nil {
		proxod(root.left, anslist)
		if root.left == nil && root.right == nil {
			*anslist = append(*anslist, root.value)
		}
		proxod(root.right, anslist)
	}
}
