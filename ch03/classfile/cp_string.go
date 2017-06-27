package classfile

type ConstantStringInfo struct {
	pool        ConstantPool
	stringIndex uint16
}

func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()
}

func (self *ConstantStringInfo) String() string {
	return self.pool.getUtf8(self.stringIndex)
}
