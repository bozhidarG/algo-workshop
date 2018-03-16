package main

import (
	"algoWorkshop/trees"
	"fmt"
	"algoWorkshop/structs"
)

func main() {
	stack := structs.Stack{}
	stack.Add(1)
	stack.Add(2)
	stack.Add(3)
	fmt.Println(stack.Pop())
	stack.Add(4)
	stack.Add(5)
	for !stack.Empty() {
		fmt.Print(stack.Pop(), " ,")
	}
	fmt.Println()

	queue := structs.Queue{}
	queue.Add(1)
	queue.Add(2)
	queue.Add(3)
	fmt.Println(queue.Head())
	queue.Add(4)
	queue.Add(5)
	queue.Add(6)

	for !queue.Empty() {
		fmt.Print(queue.Head(), " ,")
	}
	fmt.Println()


	t := trees.BinaryTree{
		Value:3,
		Left:&trees.BinaryTree{
			Value:1,
			Left:&trees.BinaryTree{
				Value:4,
			},
			Right:&trees.BinaryTree{
				Value:7,
				Left:&trees.BinaryTree{Value:8},
				Right:&trees.BinaryTree{Value:9},
			},
		},
		Right:&trees.BinaryTree{
			Value:2,
			Left:&trees.BinaryTree{Value:5},
			Right:&trees.BinaryTree{Value:6},
		},
	}

	t.DFSStack()
	fmt.Println()
	t.DFS()

	fmt.Println()
	t.BFSQueue()
}
