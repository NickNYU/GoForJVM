package classfile

import "encoding/binary"

type ClassReader struct {
	data []byte
}

func (self *ClassReader) readUint8() uint8 {
	val := self.data[0]
	self.data = self.data[1:]
	return val
}

func (self *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

// The size of the table will be provided by the table head
func (self *ClassReader) readUint16s() []uint16 {
	length := self.readUint16()
	arr := make([]uint16, length)
	for i := range arr {
		arr[i] = self.readUint16()
	}
	return arr
}

func (self *ClassReader) readBytes(length uint32) []byte {
	bytes := self.data[:length]
	self.data = self.data[length:]
	return bytes
}
