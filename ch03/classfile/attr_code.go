package classfile

/*
Code_attribute {
	u2 attribute_name_index
	u4 attribute_length
	u2 max_stack
	u2 max_locals
	u4 code_length
	u1 code[code_length]
	u2 exception_table_length
	{
		u2 start_pc
		u2 end_pc
		u2 handler_pc
		u2 catch_type
	} exception_table[exception_table_length]

	u2 attributes_count
	attribute_info attributes[attributes_count]
}
*/
type CodeAttribute struct {
	pool           ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

type ExceptionTableEntry struct {
	startPC   uint16
	endPC     uint16
	handlerPC uint16
	catchType uint16
}

func (self *CodeAttribute) readInfo(reader *ClassReader) {
	self.maxStack = reader.readUint16()
	self.maxLocals = reader.readUint16()
	codeLen := reader.readUint32()
	self.code = reader.readBytes(codeLen)
	self.exceptionTable = readExceptionTable(reader)
	self.attributes = readAttributes(reader, self.pool)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLen := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLen)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPC:   reader.readUint16(),
			endPC:     reader.readUint16(),
			handlerPC: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}
