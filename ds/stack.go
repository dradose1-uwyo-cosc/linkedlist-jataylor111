package ds

type Stack struct {
	list LinkedList
}

// add new head node
func (s *Stack) Push(value string) {
	s.list.InsertAt(0, value)
}

// remove the head
func (s *Stack) Pop() (string, bool) {
	data := s.list.Head.data
	s.list.RemoveAt(0)
	return data, true
}

