package classfile

// https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.4.3
// The value of the string_index item must be a valid index into the constant_pool table.
// The constant_pool entry at that index must be a CONSTANT_Utf8_info (§4.4.7) structure
// representing the sequence of Unicode code points to which the String object is to be initialized.
type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16 //string_index
}

func (s *ConstantStringInfo) readInfo(r *ClassReader) {
	s.stringIndex = r.readUint16()
}

//why会有个String()方法?
//string方法用于从常量池里查找名称并显示
func (s *ConstantStringInfo) String() string {
	return s.cp.getUtf8(s.stringIndex)
}
