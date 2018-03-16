package trees

import "fmt"

type BinaryTree struct {
	Left *BinaryTree
	Right *BinaryTree
	Value int
}

func (bt *BinaryTree) DFS () {
	if bt == nil {
		return
	}

	fmt.Print(bt.Value)
	fmt.Print(", ")

	bt.Left.DFS()
	bt.Right.DFS()

}

func (bt *BinaryTree) DFSStack () {
	stack := Stack{}

	stack.Add(bt)

	var val *BinaryTree

	for !stack.Empty() {
		val = stack.Pop()
		fmt.Print(val.Value, ", ")

		if val.Right != nil {
			stack.Add(val.Right)
		}

		if val.Left != nil {
			stack.Add(val.Left)
		}
	}
}

func (bt *BinaryTree) BFSQueue () {
	queue := Queue{}

	queue.Add(bt)

	var val *BinaryTree

	for !queue.Empty() {
		val = queue.Head()
		fmt.Print(val.Value, ", ")

		if val.Left != nil {
			queue.Add(val.Left)
		}

		if val.Right != nil {
			queue.Add(val.Right)
		}
	}
}