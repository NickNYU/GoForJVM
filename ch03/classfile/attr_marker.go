package classfile

type DeprecateAttribute struct {
	MarkerAttribute
}

type SyntheticAttribute struct {
	MarkerAttribute
}

type MarkerAttribute struct{}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}
