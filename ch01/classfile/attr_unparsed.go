package classfile

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
