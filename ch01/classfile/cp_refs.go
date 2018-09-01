package classfile

//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.4.2
//由于三种类型字段一样
type constantRefInfo struct {
	cp               ConstantPool
	classIndex       uint16 //u2
	nameAndTypeIndex uint16 //u2
}

func (c *constantRefInfo) readInfo(r *ClassReader) {
	c.classIndex = r.readUint16()
	c.nameAndTypeIndex = r.readUint16()
}

func (c *constantRefInfo) ClassName() string {
	return c.cp.getClassName(c.classIndex)
}

func (c *constantRefInfo) NameAndDescriptor() (string, string) {
	return c.cp.getNameAndType(c.nameAndTypeIndex)
}

//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.4.2
type ConstantFieldRefInfo struct {
	constantRefInfo
}

//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.4.2
type ConstantMethodRefInfo struct{ constantRefInfo }

//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.4.2
type ConstantInterfaceMethodRefInfo struct{ constantRefInfo }
