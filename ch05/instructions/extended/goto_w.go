package extended

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

// Branch always (wide index)
type GOTO_W struct {
	offset int
}

func (self *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	self.offset = int(reader.ReadInt32())
}


func (self *GOTO_W) Execute(frame *rtda.Frame) {
	base.Jump(frame, self.offset)
}
