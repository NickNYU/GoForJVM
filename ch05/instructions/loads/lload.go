package loads

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// Load reference(address) from local variable
type LLOAD struct{ base.Index8Instruction }
type LLOAD_0 struct{ base.Index8Instruction }
type LLOAD_1 struct{ base.Index8Instruction }
type LLOAD_2 struct{ base.Index8Instruction }
type LLOAD_3 struct{ base.Index8Instruction }

func _lload(frame *rtda.Frame, index uint) {
	stack := frame.OperandStack()
	localVars := frame.LocalVars()
	l := localVars.GetLong(index)
	stack.PushLong(l)
}

func (self *LLOAD) Execute(frame *rtda.Frame) {
	index := self.Index
	_lload(frame, index)
}

func (self *LLOAD_0) Execute(frame *rtda.Frame) {
	_lload(frame, 0)
}

func (self *LLOAD_1) Execute(frame *rtda.Frame) {
	_lload(frame, 1)
}

func (self *LLOAD_2) Execute(frame *rtda.Frame) {
	_lload(frame, 2)
}

func (self *LLOAD_3) Execute(frame *rtda.Frame) {
	_lload(frame, 3)
}
