package base

import "jvmgo/ch05/rtda"

type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}
type NoOperandsInstruction struct {
}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do
}

func (self *NoOperandsInstruction) Execute(frame *rtda.Frame) {

}

type BranchInstruction struct {
	Offset int
}
