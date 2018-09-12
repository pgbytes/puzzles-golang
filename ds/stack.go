package ds

import "errors"

type Stack struct {
	data []string
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(element string) {
	s.data = append(s.data, element)
}

func (s *Stack) Pop() (string, error) {
	topIndex := len(s.data) - 1
	if topIndex == -1 {
		return "", errors.New("stack is empty")
	}
	top := s.data[topIndex]
	s.data[topIndex-1] = "" // to avoid memory leak
	s.data = s.data[:topIndex]
	return top, nil
}

func (s *Stack) Peep() string {
	topIndex := len(s.data) - 1
	if topIndex == -1 {
		return ""
	}
	return s.data[topIndex]
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}
