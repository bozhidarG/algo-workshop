package structs

type NodeStack struct {
	next *NodeStack
	Value int
}

type Stack struct {
	head *NodeStack
}

func (s *Stack) Add (value int) {
	node := &NodeStack{Value:value}

	node.next = s.head
	s.head = node
}

func (s *Stack) Pop () (res int) {
	res = s.head.Value
	s.head = s.head.next

	return
}

func (s *Stack) Empty () (bool) {
	return s.head == nil
}