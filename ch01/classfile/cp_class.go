package classfile

// https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.4.1
// The value of the name_index item must be a valid index into the constant_pool table.
// The constant_pool entry at that index must be a CONSTANT_Utf8_info (§4.4.7) structure
// representing a valid binary class or interface name encoded in internal form (§4.2.1).
type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (c *ConstantClassInfo) readInfo(r *ClassReader) {
	c.nameIndex = r.readUint16()
}

func (c *ConstantClassInfo) Name() string {
	return c.cp.getUtf8(c.nameIndex)
}
