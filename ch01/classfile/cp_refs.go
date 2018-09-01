package classfile

//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.4.2
//由于三种类型字段一样
type ConstantRefInfo struct {
	cp               ConstantPool
	classIndex       uint16 //u2
	nameAndTypeIndex uint16 //u2
}

func (c *ConstantRefInfo) readInfo(r *ClassReader) {
	c.classIndex = r.readUint16()
	c.nameAndTypeIndex = r.readUint16()
}

func (c *ConstantRefInfo) ClassName() string {
	return c.cp.getClassName(c.classIndex)
}

func (c *ConstantRefInfo) NameAndDescriptor() (string, string) {
	return c.cp.getNameAndType(c.nameAndTypeIndex)
}

//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.4.2
type ConstantFieldRefInfo struct {
	ConstantRefInfo
}

//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.4.2
type ConstantMethodRefInfo struct{ ConstantRefInfo }

//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.4.2
type ConstantInterfaceMethodRefInfo struct{ ConstantRefInfo }
