package extended

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/instructions/loads"
)

// Extend local variable index by additional bytes
type WIDE struct {
	modifiedInstruction base.Instruction
}

func (self *WIDE) FetchOperands(reader *base.BytecodeReader) {
	opcode := reader.ReadUint8()
	switch opcode {
	case 0x15:
		inst := &loads.ILOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0xa9: //ret
		panic("Unsupported opcode: 0xa9!")
	}
}
