package classfile

//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.4.6
//TODO:需要了解一下这个类型存在的意义，对应java里哪些类型呢
type ConstantNameAndTypeInfo struct {
	cp              ConstantPool
	nameIndex       uint16
	descriptorIndex uint16
}

func (n *ConstantNameAndTypeInfo) readInfo(r *ClassReader) {
	n.nameIndex = r.readUint16()
	n.descriptorIndex = r.readUint16()
}
