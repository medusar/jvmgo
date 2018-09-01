package classfile

//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.4.8
// CONSTANT_MethodHandle_info {
//     u1 tag;
//     u1 reference_kind;
//     u2 reference_index;
// }
type ConstantMethodHandleInfo struct {
	cp             ConstantPool
	referenceKind  uint8
	referenceIndex uint16
}

func (m *ConstantMethodHandleInfo) readInfo(r *ClassReader) {
	m.referenceKind = r.readUint8()
	m.referenceIndex = r.readUint16()
}
