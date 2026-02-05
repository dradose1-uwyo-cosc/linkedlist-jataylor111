package ds

type Queue struct {
	list LinkedList
}

// add tail node
func (q *Queue) Push(value string) {
	q.list.Insert(value)
}

// remove the head
func (q *Queue) Pop() (string, error) {
	data := q.list.Head.data
	q.list.RemoveAt(0)
	return data, nil
}
