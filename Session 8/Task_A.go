package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type root struct {
	value  int
	length int
	left   *root
	right  *root
}

type tree struct {
	root *root
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	tree := new(tree)
	for {
		line, _ := reader.ReadString('\n')
		if line == "" || line == "\n" {
			break
		}
		params := strings.Fields(line)
		switch params[0] {
		case "ADD":
			value, _ := strconv.Atoi(params[1])
			tree.Insert(value, 0)
		case "SEARCH":
			value, _ := strconv.Atoi(params[1])
			tree.Search(value)
		case "PRINTTREE":
			tree.Print()
		}
	}
}

func initialization(value, length int) *root {
	return &root{value, length, nil, nil}
}

func (tree *tree) Insert(value, length int) {
	if tree.root == nil {
		tree.root = initialization(value, length)
		fmt.Println("DONE")
	} else {
		tree.root.insert(value, length+1)
	}
}

func (root *root) insert(value, length int) {
	if value > root.value {
		if root.right == nil {
			root.right = initialization(value, length)
			fmt.Println("DONE")
		} else {
			root.right.insert(value, length+1)
		}
	} else if value < root.value {
		if root.left == nil {
			root.left = initialization(value, length)
			fmt.Println("DONE")
		} else {
			root.left.insert(value, length+1)
		}
	} else {
		fmt.Println("ALREADY")
	}
}

func (tree *tree) Search(value int) {
	if tree.root == nil {
		fmt.Println("NO")
	} else {
		tree.root.search(value)
	}
}

func (root *root) search(value int) {
	if value > root.value {
		if root.right == nil {
			fmt.Println("NO")
		} else {
			root.right.search(value)
		}
	} else if value < root.value {
		if root.left == nil {
			fmt.Println("NO")
		} else {
			root.left.search(value)
		}
	} else {
		fmt.Println("YES")
	}
}

func (tree *tree) Print() {
	if tree.root != nil {
		tree.root.print()
	}
}

func (root *root) print() {
	if root != nil {
		root.left.print()
		for i := 0; i < root.length; i++ {
			fmt.Print(".")
		}
		fmt.Print(root.value)
		fmt.Println()
		root.right.print()
	}
}
