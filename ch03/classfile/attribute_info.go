package classfile

type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, pool ConstantPool) []AttributeInfo {

}

func readAttribute(reader *ClassReader, pool ConstantPool) AttributeInfo {

}

func newAttributeInfo(attrName string, attrLen uint32, pool ConstantPool) AttributeInfo {

}
