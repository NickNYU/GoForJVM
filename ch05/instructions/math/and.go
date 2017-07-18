package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type IAND struct{ base.NoOperandsInstruction }
type LAND struct{ base.NoOperandsInstruction }

func (self *IAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	result := val1 & val2
	stack.PushInt(result)
}

func (self *LAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopLong()
	val1 := stack.PopLong()
	result := val1 & val2
	stack.PushLong(result)
}
