// stack.go
package stack

import (
    "errors"
)

type Stack interface {
    Push(item interface{}) error
    Pop() (interface{}, error)
    Peek() (interface{}, error)
}

type StackImpl struct {
    items []interface{}
}

func NewStack() Stack {
    return &StackImpl{
        items: []interface{}{},
    }
}

func (s *StackImpl) Push(item interface{}) error {
    s.items = append(s.items, item)
    return nil
}

func (s *StackImpl) Pop() (interface{}, error) {
    if len(s.items) == 0 {
        return nil, errors.New("Empty Stack")
    }
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item, nil
}

func (s *StackImpl) Peek() (interface{}, error) {
    if len(s.items) == 0 {
        return nil, errors.New("Empty Stack")
    }
    return s.items[len(s.items)-1], nil
}
