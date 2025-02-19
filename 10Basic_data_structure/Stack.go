package Basic_data_structure

import (
	"errors"
)

type Stack struct {
	List []int
}

func (s *Stack) StackEmpty() bool {
	return len(s.List) == 0
}
func (s *Stack) Push(v int) {
	s.List = append(s.List, v)
}
func (s *Stack) Pop() (int, error) {
	if len(s.List) == 0 {
		return 0, errors.New("stack is empty")
	} else {
		v := s.List[len(s.List)-1]
		s.List = s.List[:len(s.List)-1]
		return v, nil
	}
}
