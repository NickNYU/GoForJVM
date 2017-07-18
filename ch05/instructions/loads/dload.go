package loads

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// Load reference(address) from local variable
type DLOAD struct{ base.Index8Instruction }
type DLOAD_0 struct{ base.Index8Instruction }
type DLOAD_1 struct{ base.Index8Instruction }
type DLOAD_2 struct{ base.Index8Instruction }
type DLOAD_3 struct{ base.Index8Instruction }

func _dload(frame *rtda.Frame, index uint) {
	d := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(d)
}

func (self *DLOAD) Execute(frame *rtda.Frame) {
	index := self.Index
	_dload(frame, index)
}

func (self *DLOAD_0) Execute(frame *rtda.Frame) {
	_dload(frame, 0)
}

func (self *DLOAD_1) Execute(frame *rtda.Frame) {
	_dload(frame, 1)
}

func (self *DLOAD_2) Execute(frame *rtda.Frame) {
	_dload(frame, 2)
}

func (self *DLOAD_3) Execute(frame *rtda.Frame) {
	_dload(frame, 3)
}

