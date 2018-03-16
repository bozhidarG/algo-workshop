package structs

type NodeQueue struct {
	next *NodeQueue
	Value int
}

type Queue struct {
	head *NodeQueue
	tail *NodeQueue
}

func (s *Queue) Add (value int) {
	node := &NodeQueue{Value:value}

	if s.head == nil && s.tail == nil {
		s.head = node
		s.tail = node
	} else {
		s.tail.next = node
		s.tail = s.tail.next
	}
}

func (s *Queue) Head () (res int) {

	res = s.head.Value
	s.head = s.head.next

	if s.head == nil {
		s.tail = nil
	}

	return
}

func (s *Queue) Empty () (bool) {
	return s.head == nil
}
