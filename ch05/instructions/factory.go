package instructions

import (
	"fmt"
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/instructions/constants"
)

func NewInstruction(opcode byte) base.Instruction {
	switch opcode {
	case 0x00:
		return &constants.NOP{}
	case 0x01:
		return &constants.ACONST_NULL{}
	default:
		panic(fmt.Errorf("Unsupported opcode: 0x%x!", opcode))
	}
}
