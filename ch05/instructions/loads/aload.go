package loads

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// Load reference(address) from local variable
type ALOAD struct{ base.Index8Instruction }
type ALOAD_0 struct{ base.Index8Instruction }
type ALOAD_1 struct{ base.Index8Instruction }
type ALOAD_2 struct{ base.Index8Instruction }
type ALOAD_3 struct{ base.Index8Instruction }

func _aload(frame *rtda.Frame, index uint) {
	stack := frame.OperandStack()
	localVars := frame.LocalVars()
	ref := localVars.GetRef(index)
	stack.PushRef(ref)
}

func (self *ALOAD) Execute(frame *rtda.Frame) {
	index := self.Index
	_aload(frame, index)
}

func (self *ALOAD_0) Execute(frame *rtda.Frame) {
	_aload(frame, 0)
}

func (self *ALOAD_1) Execute(frame *rtda.Frame) {
	_aload(frame, 1)
}

func (self *ALOAD_2) Execute(frame *rtda.Frame) {
	_aload(frame, 2)
}

func (self *ALOAD_3) Execute(frame *rtda.Frame) {
	_aload(frame, 3)
}
