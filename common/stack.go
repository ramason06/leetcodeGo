package common

import "errors"

type Stack struct {
	elements []string
}

func (s *Stack) Push(value string) {
	s.elements = append(s.elements, value)
}

func (s *Stack) Pop() (string, error) {
	if len(s.elements) == 0 {
		return "", errors.New("stack is empty")
	}
	value := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return value, nil
}

func (s *Stack) Peek() (string, error) {
	if len(s.elements) == 0 {
		return "", errors.New("stack is empty")
	}
	return s.elements[len(s.elements)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack) Size() int {
	return len(s.elements)
}
