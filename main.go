package main

type queue struct {
	first *queueNode
	last  *queueNode
	size  int
}
type queueNode struct {
	value interface{}
	next  *queueNode
	prev  *queueNode
}

func (receiver *queue) Len() int {
	return receiver.size
}
func (receiver *queue) First() *queueNode {
	return receiver.first
}
func (receiver *queue) Last() *queueNode {
	return receiver.last
}
func (receiver *queue) equeue(addElement interface{}) {
	if receiver.Len() == 0 {
		receiver.first = &queueNode{
			value: addElement,
			next:  nil,
			prev:  nil,
		}
		receiver.last = receiver.First()
		receiver.size++
		return
	}
	receiver.size++
	current := receiver.First()
	for {
		if current.next == nil {
			current.next = &queueNode{
				addElement,
				nil,
				current,
			}
			return
		}
		current = current.next
	}
}
func (receiver *queue) dequeue() {
	if receiver.Len() == 0 {
		return
	}

	if receiver.Len() == 1 {
		receiver.first = nil
		receiver.last = nil
	}
	receiver.size--
}
func main() {}
