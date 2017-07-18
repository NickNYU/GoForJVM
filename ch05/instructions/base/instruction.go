package base

import "jvmgo/ch05/rtda"

type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct {}

// Jump Instruction
type BranchInstruction struct { Offset int }
// Load & Store Instruction
type Index8Instruction struct { Index uint }
// Runtime Constant Pool Instruction
type Index16Instruction struct { Index uint }



func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}


func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}
