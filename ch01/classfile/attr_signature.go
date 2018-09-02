package classfile

import "strconv"

//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.7.9
/*
Signature_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 signature_index;
}
*/
type SignatureAttribute struct {
	cp             ConstantPool
	signatureIndex uint16
}

func (s *SignatureAttribute) readInfo(r *ClassReader) {
	s.signatureIndex = r.readUint16()
}

//格式参考javap -verbose输出
func (s *SignatureAttribute) String() string {
	return "    Signature: #" + strconv.Itoa(int(s.signatureIndex))
}
