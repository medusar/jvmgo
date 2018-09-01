package classfile

//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.4.9
// CONSTANT_MethodType_info {
//     u1 tag;
//     u2 descriptor_index;
// }
type ConstantMethodTypeInfo struct {
	cp              ConstantPool
	descriptorIndex uint16
}

func (m *ConstantMethodTypeInfo) readInfo(r *ClassReader) {
	m.descriptorIndex = r.readUint16()
}
