package main

import "testing"

func Test_queue_Len(t *testing.T) {
	q := queue{}
	q.equeue(0)
	if q.Len() != 1 {
		t.Error("После добавление элемента длина не должна быть нулевым got:", q.Len())
	}
}
func Test_queue_firstElement_lastElement(t *testing.T) {
	q := queue{}
	q.equeue(1)
	q.equeue(2)
	q.equeue(3)
	if q.First() != q.Last() {
		t.Error("После добавление одного элемента, первый элемент должен быть равен последнему got:", q.First())
	}
}
func Test_queue_deleting_first_element(t *testing.T) {
	q := queue{}
	q.equeue(1)
	q.dequeue()
	if q.First() != nil {
		t.Error("После удаление одного элемента, длина должна быть пустым got:", q.First())
	}
}
