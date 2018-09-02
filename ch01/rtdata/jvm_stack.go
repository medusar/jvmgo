package rtdata

//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-2.html#jvms-2.5.2
type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame
}

func (s *Stack) push(frame *Frame) {
	if s.size >= s.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if s._top != nil {
		frame.lower = s._top
	}
	s._top = frame
	s.size++
}

func (s *Stack) pop() *Frame {
	if s._top == nil {
		panic("jvm stack is empty")
	}
	top := s._top
	s._top = top.lower
	top.lower = nil //important
	s.size--

	return top
}

func (s *Stack) top() *Frame {
	if s._top == nil {
		panic("jvm stack is empty")
	}
	return s._top
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
		size:    0,
		_top:    nil,
	}
}
