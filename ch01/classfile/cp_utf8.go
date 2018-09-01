package classfile

//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.4.7
//规范里，CONSTANT_Utf8_Info是一个字节数组加字节数组长度
//这里定义直接使用了string,长度加字节数组在readInfo方法中可以体现
type ConstantUtf8Info struct {
	str string
}

func (u *ConstantUtf8Info) readInfo(r *ClassReader) {
	length := uint32(r.readUint16())
	//按长度读取字节
	bytes := r.readBytes(length)
	//String content is encoded in modified UTF-8.
	//Modified UTF-8 strings are encoded so that
	//code point sequences that contain only non-null ASCII characters
	// can be represented using only 1 byte per code point,
	//but all code points in the Unicode codespace can be represented.
	//Java中的utf8用的是Modified UTF-8,所以这里需要处理一下，文档里有具体转换方式
	u.str = decodeMUTF8(bytes)
}

//这里暂时简化实现，后面再按照规范实现
func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}
