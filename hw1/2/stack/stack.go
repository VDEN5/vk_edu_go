package stack

type Stack struct {
	data []byte
}

func (s *Stack) Push(v byte) {
	s.data = append(s.data, v)
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}
func (s *Stack) Top() byte {
	return (*&s.data)[len(*&s.data)-1]
}

func (s *Stack) Pop() byte {
	v := s.Top()
	s.data = s.data[:len(s.data)-1]
	return v
}

func (s *Stack) isEmpty() bool {
	return len(s.data) == 0
}
