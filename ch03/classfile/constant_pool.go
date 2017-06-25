package classfile

type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	length := int(reader.readUint16())
	pool := make(ConstantPool, length)
	for index := 1; index < length; index++ {
		pool[index] = readConstantInfo(reader, pool)
		switch pool[index].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			index++
		}
	}
	return pool
}

func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if poolInfo := self[index]; poolInfo != nil {
		return poolInfo
	}
	panic("Invalid constant pool index!")
}

func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	nameAndType := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(nameAndType.nameIndex)
	_type := self.getUtf8(nameAndType.descriptorIndex)
	return name, _type
}

func (self ConstantPool) getClassName(index uint16) string {
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}

func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
