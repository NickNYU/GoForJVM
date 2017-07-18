package loads

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// Load reference(address) from local variable
type FLOAD struct{ base.Index8Instruction }
type FLOAD_0 struct{ base.Index8Instruction }
type FLOAD_1 struct{ base.Index8Instruction }
type FLOAD_2 struct{ base.Index8Instruction }
type FLOAD_3 struct{ base.Index8Instruction }

func _fload(frame *rtda.Frame, index uint) {
	f := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(f)
}

func (self *FLOAD) Execute(frame *rtda.Frame) {
	index := self.Index
	_fload(frame, index)
}

func (self *FLOAD_0) Execute(frame *rtda.Frame) {
	_fload(frame, 0)
}

func (self *FLOAD_1) Execute(frame *rtda.Frame) {
	_fload(frame, 1)
}

func (self *FLOAD_2) Execute(frame *rtda.Frame) {
	_fload(frame, 2)
}

func (self *FLOAD_3) Execute(frame *rtda.Frame) {
	_fload(frame, 3)
}

