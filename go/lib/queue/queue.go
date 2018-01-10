// package queue contains interface for a queue and a simple implementaton

package queue

import (
	"errors"
)

type Queue interface {
	Push(interface{})
	Pop() error
	Front() (interface{}, error)
	Back() (interface{}, error)
	Empty() bool
	Size() int
}

// no generics kind of sucks...
type SimpleQueue []interface{}

func (q *SimpleQueue) Push(e interface{}) {
	*q = append(*q, e)
}

func (q *SimpleQueue) Pop() error {
	if len(*q) == 0 {
		return errors.New("queue is empty")
	}
	*q = (*q)[1:]
	return nil
}

func (q *SimpleQueue) Front() (interface{}, error) {
	if len(*q) == 0 {
		return nil, errors.New("queue is empty")
	}
	return (*q)[0], nil
}

func (q *SimpleQueue) Back() (interface{}, error) {
	if len(*q) == 0 {
		return nil, errors.New("queue is empty")
	}
	return (*q)[len(*q)], nil
}

func (q *SimpleQueue) Empty() bool {
	return len(*q) == 0
}

func (q *SimpleQueue) Size() int {
	return len(*q)
}
