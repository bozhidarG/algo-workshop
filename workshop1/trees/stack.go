package trees

type NodeStack struct {
	next *NodeStack
	Value *BinaryTree
}

type Stack struct {
	head *NodeStack
}

func (s *Stack) Add (value *BinaryTree) {
	node := &NodeStack{Value:value}

	node.next = s.head
	s.head = node
}

func (s *Stack) Pop () (res *BinaryTree) {
	res = s.head.Value
	s.head = s.head.next

	return
}

func (s *Stack) Empty () (bool) {
	return s.head == nil
}