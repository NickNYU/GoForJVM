package classfile

const (
	CONSTANT_CLASS                = 7
	CONSTANT_FIELD_REF            = 9
	CONSTANT_METHOD_REF           = 10
	CONSTANT_INTERFACE_METHOD_REF = 11
	CONSTANT_STRING               = 8
	CONSTANT_INTEGER              = 3
	CONSTANT_FLOAT                = 4
	CONSTANT_LONG                 = 5
	CONSTANT_DOUBLE               = 6
	CONSTANT_NAME_AND_TYPE        = 12
	CONSTANT_UTF8                 = 1
	CONSTANT_METHOD_HANDLE        = 15
	CONSTANT_METHOD_TYPE          = 16
	CONSTANT_INVOKE_DYNAMIC       = 18
)

type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

func readConstantInfo(reader *ClassReader, pool ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	constant := newConstantInfo(tag, pool)
	constant.readInfo(reader)
	return constant
}

func newConstantInfo(tag uint8, pool ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_INTEGER:
		return &ConstantIntegerInfo{}

	}
}
