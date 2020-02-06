package main

import "testing"

func Test_empty_queue_Len(t *testing.T) {
	q := queue{}
	if q.Len() != 0 {
		t.Error("Длина должна быть нулевым got:", q.Len())
	}
	if q.First() != nil {
		t.Error("Первый элемент должен быть нулевым got:", q.First())
	}
	if q.Last() != nil {
		t.Error("Последный элемент должен быть нулевым got:", q.Last())
	}

}
func Test_queue_with_one_element(t *testing.T) {
	q := queue{}
	q.equeue(1)
	if q.Len() != 1 {
		t.Error("После добавление одного элемента длина не должна быть нулевым, got:", q.Len())
	}
	if q.First() != 1 {
		t.Error("После добавление одного элемента Первый элемент должен быть 1, got:", q.First())
	}
	if q.Last() != 1 {
		t.Error("После добавление одного элемента Последный элемент должен быть 1, got:", q.Last())
	}
	if q.Last() != q.First() {
		t.Error("После добавление одного элемента Первый элемент и Последный элемент должны иметь одинаковые значение, got:", q.Last(), q.First())
	}
}
func Test_queue_with_multiple_elements(t *testing.T) {
	q := queue{}
	q.equeue(1)
	q.equeue(2)
	q.equeue(3)
	if q.Len() != 3 {
		t.Error("После добавление три элемента длина должна быть 3, got:", q.Len())
	}
	if q.First() != 1 {
		t.Error("Первый элемент должен быть 1, got:", q.First())
	}
	if q.Last() != 3 {
		t.Error("Последный элемент должен быть 3, got:", q.Last())
	}
}
func Test_queue_firstElement_lastElement(t *testing.T) {
	q := queue{}
	q.equeue(3)
	q.equeue(4)
	q.equeue(5)
	if q.First() != 3 {
		t.Error("Первый элемент должен быть 3, got:", q.First())
	}
	if q.Last() != 5 {
		t.Error("Последный элемент должен быть 5, got:", q.Last())
	}
	if q.First() == q.Last() {
		t.Error("Когда добавили 3 элемента, первый элемент не может быть равен последнему, got:", q.First(), q.Last())
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
