package classfile

//通用的Attribute，对于一些比较复杂或者需要调试还没有来得及实现的Attribute
//用该Attribute可以跳过对应字节，从而不影响测试
//attribute_info {
// u2 attribute_name_index;
// u4 attribute_length;
// u1 info[attribute_length];
// }
type UnparsedAttribute struct {
	name   string
	length uint32
	info   []byte
}

func (a *UnparsedAttribute) readInfo(r *ClassReader) {
	a.info = r.readBytes(a.length)
}

func (a *UnparsedAttribute) String() string {
	return "attribute type: " + a.name + " unparsed yet"
}
