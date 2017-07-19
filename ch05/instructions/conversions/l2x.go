package conversions

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type L2F struct{ base.NoOperandsInstruction }
type L2I struct{ base.NoOperandsInstruction }
type L2D struct{ base.NoOperandsInstruction }

// convert long to float
func (self *L2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	f := float32(l)
	stack.PushFloat(f)
}

// convert long to int
func (self *L2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	i := int32(l)
	stack.PushInt(i)
}

// convert long to double
func (self *L2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	d := float64(l)
	stack.PushDouble(d)
}

