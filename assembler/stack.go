package assembler

import "errors"

type itemStack struct {
	items []interface{}
}

type stack interface {
	push(interface{})
	pop() (interface{}, error)
	peek() (interface{}, error)
	isEmpty() bool
}

func newStack() stack {
	// create stack with default capacity of 20
	return &itemStack{make([]interface{}, 0, 20)}
}

func (s *itemStack) push(i interface{}) {
	s.items = append(s.items, i)
}

func (s *itemStack) pop() (interface{}, error) {
	top := len(s.items) - 1

	if top < 0 {
		return nil, errors.New("Stack is empty")
	}

	i := s.items[top]
	s.items = s.items[:top]
	return i, nil
}

func (s *itemStack) peek() (interface{}, error) {
	top := len(s.items) - 1

	if top < 0 {
		return nil, errors.New("Stack is empty")
	}

	i := s.items[top]
	return i, nil
}

func (s *itemStack) isEmpty() bool {
	return len(s.items) == 0
}
