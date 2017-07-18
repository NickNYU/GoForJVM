package stores

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

// Store reference(address) into local variable
type ASTORE struct{ base.Index8Instruction }
type ASTORE_0 struct{ base.Index8Instruction }
type ASTORE_1 struct{ base.Index8Instruction }
type ASTORE_2 struct{ base.Index8Instruction }
type ASTORE_3 struct{ base.Index8Instruction }

func _astore(frame *rtda.Frame, index uint) {
	ref := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, ref)
}

func (self *ASTORE) Execute(frame *rtda.Frame) {
	_astore(frame, uint(self.Index))
}

func (self *ASTORE_0) Execute(frame *rtda.Frame) {
	_astore(frame, 0)
}

func (self *ASTORE_1) Execute(frame *rtda.Frame) {
	_astore(frame, 1)
}

func (self *ASTORE_2) Execute(frame *rtda.Frame) {
	_astore(frame, 2)
}

func (self *ASTORE_3) Execute(frame *rtda.Frame) {
	_astore(frame, 3)
}
