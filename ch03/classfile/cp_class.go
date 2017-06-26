package classfile

type ConstantClassInfo struct {
	pool      ConstantPool
	nameIndex uint16
}

func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
}

func (self *ConstantClassInfo) Name() string {
	return self.pool.getUtf8(self.nameIndex)
}
