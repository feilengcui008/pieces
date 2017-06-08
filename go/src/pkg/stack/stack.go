// package stack contains interface for a stack and a simple implementaton

package stack

import (
	"errors"
)

type Stack interface {
	Push(interface{})
	Pop() error
	Top() (interface{}, error)
	Empty() bool
	Size() int
}

type SimpleStack []interface{}

func (q *SimpleStack) Push(e interface{}) {
	*q = append(*q, e)
}

func (q *SimpleStack) Pop() error {
	if len(*q) == 0 {
		return errors.New("queue is empty")
	}
	*q = (*q)[:len(*q)-1]
	return nil
}

func (q *SimpleStack) Top() (interface{}, error) {
	if len(*q) == 0 {
		return nil, errors.New("queue is empty")
	}
	return (*q)[len(*q)-1], nil
}

func (q *SimpleStack) Empty() bool {
	return len(*q) == 0
}

func (q *SimpleStack) Size() int {
	return len(*q)
}
