package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func main() {

}

func connect(root *Node) *Node {
	breadth_first := make([][]*Node, 1000)
	getHeight(root, breadth_first)
	for _, list := range breadth_first {
		for counter, item := range list {
			fmt.Printf("len(list): %v\n", len(list))
			fmt.Printf("counter: %v\n", counter-2)
			if len(list) > counter-2 {
				item.Next = list[counter+1]
			} else {
				item.Next = nil
			}
		}

	}
	return root
}

func getHeight(root *Node, breadth_first [][]*Node) int {

	if root == nil {
		return -1
	}
	height := max(getHeight(root.Left, breadth_first), getHeight(root.Right, breadth_first)) + 1

	breadth_first[height] = append(breadth_first[height], root)
	return height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
