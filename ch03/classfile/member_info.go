package classfile

type MemberInfo struct {
	constantPool ConstantPool
	accessFlag uint16
	nameIndex uint16
	descriptorIndex uint16
	attributes []AttributeInfo
}

func readMembers(reader *ClassReader, pool ConstantPool) []*MemberInfo {
	length := reader.readUint16()
	memberInfos := make([]*MemberInfo, length)
	for i := range memberInfos {
		memberInfos[i] = readMember(reader, pool)
	}
	return memberInfos
}

func readMember(reader *ClassReader, pool ConstantPool) *MemberInfo {
	return &MemberInfo{
		constantPool: pool,
		accessFlag: reader.readUint16(),
		nameIndex: reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes: readAttributes(reader, pool)
	}
}

func (self *MemberInfo) AccessFlag() uint16 {
	return self.accessFlag
}

func (self *MemberInfo) Name() uint16 {
	return self.constantPool.getUtf8(self.nameIndex)
}

func (self *MemberInfo) Descriptor() uint16 {
	return self.constantPool.getUtf8(self.descriptorIndex)
}
