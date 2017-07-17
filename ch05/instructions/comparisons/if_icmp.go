package comparisons

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// Jump if int comparison succeeds
// equals
type IF_ICMPEQ struct{ base.BranchInstruction }

// not equals
type IF_ICMPNE struct{ base.BranchInstruction }

// less than
type IF_ICMPLT struct{ base.BranchInstruction }

// less and equal
type IF_ICMPLE struct{ base.BranchInstruction }

// greater than
type IF_ICMPGT struct{ base.BranchInstruction }

// greater equal
type IF_ICMPGE struct{ base.BranchInstruction }

func (self *IF_ICMPEQ) Execute(frame *rtda.Frame) {
	val1, val2 := _getInts(frame)
	if val1 == val2 {
		base.Jump(frame, self.Offset)
	}
}

func (self *IF_ICMPNE) Execute(frame *rtda.Frame) {
	val1, val2 := _getInts(frame)
	if val1 != val2 {
		base.Jump(frame, self.Offset)
	}
}

func (self *IF_ICMPLT) Execute(frame *rtda.Frame) {
	val1, val2 := _getInts(frame)
	if val1 < val2 {
		base.Jump(frame, self.Offset)
	}
}

func (self *IF_ICMPLE) Execute(frame *rtda.Frame) {
	val1, val2 := _getInts(frame)
	if val1 <= val2 {
		base.Jump(frame, self.Offset)
	}
}

func (self *IF_ICMPGT) Execute(frame *rtda.Frame) {
	val1, val2 := _getInts(frame)
	if val1 > val2 {
		base.Jump(frame, self.Offset)
	}
}

func (self *IF_ICMPGE) Execute(frame *rtda.Frame) {
	val1, val2 := _getInts(frame)
	if val1 >= val2 {
		base.Jump(frame, self.Offset)
	}
}

func _getInts(frame *rtda.Frame) (val1, val2 int32) {
	stack := frame.OperandStack()
	val2 = stack.PopInt()
	val1 = stack.PopInt()
	return val1, val2
}
