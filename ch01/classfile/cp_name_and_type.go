package classfile

//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.4.6
//used to represent a field or method, without indicating which class or interface type it belongs to
// CONSTANT_NameAndType_info {
//     u1 tag;
//     u2 name_index;
//     u2 descriptor_index;
// }
type ConstantNameAndTypeInfo struct {
	cp              ConstantPool
	nameIndex       uint16
	descriptorIndex uint16
}

func (n *ConstantNameAndTypeInfo) readInfo(r *ClassReader) {
	n.nameIndex = r.readUint16()
	n.descriptorIndex = r.readUint16()
}
