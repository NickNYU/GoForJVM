package rtda


type Frame struct {
	lower *Frame
	localVars LocalVars
	operandStack *OperandStack
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		lower: nil,
		localVars: newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}