package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// Add double
type DADD struct{ base.NoOperandsInstruction }

// Add float
type FADD struct{ base.NoOperandsInstruction }

// Add Int
type IADD struct{ base.NoOperandsInstruction }

// Add long
type LADD struct{ base.NoOperandsInstruction }

func (self *DADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopDouble()
	val2 := stack.PopDouble()
	stack.PushDouble(val1 + val2)
}

func (self *FADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopFloat()
	val2 := stack.PopFloat()
	stack.PushFloat(val1 + val2)
}

func (self *IADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopInt()
	val2 := stack.PopInt()
	stack.PushInt(val1 + val2)
}

func (self *LADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopLong()
	val2 := stack.PopLong()
	stack.PushLong(val1 + val2)
}
