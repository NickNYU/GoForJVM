package extended

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

// Branch if reference is null
type IFNULL struct{ base.BranchInstruction }
// Branch if reference not null
type IFNONNULL struct{ base.BranchInstruction }


func (self *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Jump(frame, self.Offset)
	}
}


func (self *IFNONNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Jump(frame, self.Offset)
	}
}
