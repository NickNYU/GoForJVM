package comparisons

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// Jump if reference(address) comparison succeeds
// EQ == Equals
type IF_ACMPEQ struct{ base.BranchInstruction }

// NE == Not Equals
type IF_ACMPNE struct{ base.BranchInstruction }

// If equals, then jump
func (self *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	if _acmp(frame) {
		base.Jump(frame, self.Offset)
	}
}

// If not equals, then jump
func (self *IF_ACMPNE) Execute(frame *rtda.Frame) {
	if !_acmp(frame) {
		base.Jump(frame, self.Offset)
	}
}

func _acmp(frame *rtda.Frame) bool {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	return ref1 == ref2 // todo
}
