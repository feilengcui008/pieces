package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	q := new(SimpleStack)
	q.Push(12)
	q.Push("string")
	e1, err := q.Top()
	if err != nil {
		t.Fail()
	}
	if s, ok := e1.(string); !ok || s != "string" {
		t.Fail()
	}
	q.Pop()
	e2, err := q.Top()
	if err != nil {
		t.Fail()
	}
	if i, ok := e2.(int); !ok || i != 12 {
		t.Fail()
	}
	q.Pop()
	if q.Size() != 0 {
		t.Fail()
	}
}
