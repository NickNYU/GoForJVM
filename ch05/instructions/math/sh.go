package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type ISHL struct{ base.NoOperandsInstruction }
type ISHR struct{ base.NoOperandsInstruction }
// unsigned right shift
type IUSHR struct{ base.NoOperandsInstruction }
type LSHL struct{ base.NoOperandsInstruction }
type LSHR struct{ base.NoOperandsInstruction }
type LUSHR struct{ base.NoOperandsInstruction }

func (self *ISHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	// int 总共32位，所以只取5位
	s := uint32(val2) & 0x1f
	result := val1 << s
	stack.PushInt(result)
}

func (self *ISHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 >> s
	stack.PushInt(result)
}

func (self *IUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	s := uint32(val2) & 0x1f
	result := int32(uint32(val1) >> s)
	stack.PushInt(result)
}

func (self *LSHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopLong()
	// Long 型有64位，2的6次方，比32为Int多一位
	s := uint32(val2) & 0x3f
	result := val1 << s
	stack.PushLong(result)
}

func (self *LSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopLong()
	s := uint32(val2) & 0x3f
	result := val1 >> s
	stack.PushLong(result)
}

func (self *LUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopLong()
	s := uint32(val2) & 0x3f
	result := int64(uint64(val1) >> s)
	stack.PushLong(result)
}