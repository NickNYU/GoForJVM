package comparisons

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// Compare Double
type DCMPG struct{ base.NoOperandsInstruction }

func (self *DCMPG) Execute(frame *rtda.Frame) {
	_dcmp(frame, false)
}

func _dcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	d2 := stack.PopDouble()
	d1 := stack.PopDouble()
	if d1 > d2 {
		stack.PushInt(1)
	} else if d1 < d2 {
		stack.PushInt(-1)
	} else {
		if gFlag {
			stack.PushInt(1)
		} else {
			stack.PushInt(-1)
		}
	}
}
