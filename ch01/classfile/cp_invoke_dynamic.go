package classfile

//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.4.10
// CONSTANT_InvokeDynamic_info {
//     u1 tag;
//     u2 bootstrap_method_attr_index;
//     u2 name_and_type_index;
// }
type ConstantInvokeDynamicInfo struct {
	cp                       ConstantPool
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

func (i *ConstantInvokeDynamicInfo) readInfo(r *ClassReader) {
	i.bootstrapMethodAttrIndex = r.readUint16()
	i.nameAndTypeIndex = r.readUint16()
}
