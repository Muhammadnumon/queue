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
func (receiver *queue) First() interface{} {
	if receiver.first == nil {
		return nil
	}
	return receiver.first
}
func (receiver *queue) Last() interface{} {
	if receiver.last == nil {
		return nil
	}
	return receiver.last
}
func (receiver *queue) equeue(addElement interface{}) {
	if receiver.Len() == 0 {
		receiver.first = &queueNode{
			value: addElement,
			next:  nil,
			prev:  nil,
		}
		receiver.last = receiver.first
		receiver.size++
		return
	}
	receiver.size++
	current := receiver.first
	for {
		if current.next == nil {
			current.next = &queueNode{
				addElement,
				nil,
				current,
			}
		}
		current = current.next
	}
}
func (receiver *queue) dequeue() interface{} {
	element := receiver.first.value
	if receiver.Len() == 0 {
		return nil
	}

	if receiver.Len() == 1 {
		receiver.first = nil
		receiver.last = nil
		receiver.size--
		return element
	}
	receiver.first = receiver.first.next
	receiver.size--
	return element
}
func main() {}
