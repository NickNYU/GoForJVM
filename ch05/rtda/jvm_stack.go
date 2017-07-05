package rtda

type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame
}

func (self *Stack) push(frame *Frame) {
	if self.size == self.maxSize {
		panic("java.lang.StackOverflowError")
	}
	frame.lower = self._top
	self._top = frame
	self.size++
}

func (self *Stack) pop() *Frame {
	if self.size == 0 {
		panic("jvm stack is empty!")
	}
	frame := self._top
	self._top = frame.lower
	frame.lower = nil
	self.size--
	return frame
}

func (self *Stack) top() *Frame {
	if self.size == 0 {
		panic("jvm stack is empty!")
	}
	return self._top
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
		size:    0,
		_top:    nil,
	}
}
