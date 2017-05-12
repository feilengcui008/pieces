package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := new(SimpleQueue)
	q.Push(12)
	q.Push("string")
	e1, err := q.Front()
	if err != nil {
		t.Fail()
	}
	if i, ok := e1.(int); !ok || i != 12 {
		t.Fail()
	}
	q.Pop()
	e2, err := q.Front()
	if err != nil {
		t.Fail()
	}
	if s, ok := e2.(string); !ok || s != "string" {
		t.Fail()
	}
	q.Pop()
	if q.Size() != 0 {
		t.Fail()
	}
}
