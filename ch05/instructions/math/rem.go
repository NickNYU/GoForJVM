package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
	"math"
)

type DREM struct{ base.NoOperandsInstruction }
type FREM struct{ base.NoOperandsInstruction }
type IREM struct{ base.NoOperandsInstruction }
type LREM struct{ base.NoOperandsInstruction }

func (self *IREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val2 == 0 {
		panic("java.lang.ArithmeticException: /   by zero")
	}
	result := val1 % val2
	stack.PushInt(result)
}

func (self *DREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopDouble()
	val1 := stack.PopDouble()
	result := math.Mod(val1, val2)
	stack.PushDouble(result)
}

func (self *LREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := v1 % v2
	stack.PushLong(result)
}

func (self *FREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := float32(math.Mod(float64(v1), float64(v2)))
	stack.PushFloat(result)
}
