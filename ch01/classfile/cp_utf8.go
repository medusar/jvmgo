package classfile

import (
	"fmt"
	"unicode/utf16"
)

//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.4.7
//规范里，CONSTANT_Utf8_Info是一个字节数组加字节数组长度
//这里定义直接使用了string,长度加字节数组在readInfo方法中可以体现
// CONSTANT_Utf8_info {
//     u1 tag;
//     u2 length;
//     u1 bytes[length];
// }
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

	utflen := len(bytes)
	chararr := make([]uint16, utflen)

	var c, char2, char3 uint16
	count := 0
	chararr_count := 0

	for count < utflen {
		c = uint16(bytes[count])
		if c > 127 {
			break
		}
		count++
		chararr[chararr_count] = c
		chararr_count++
	}

	for count < utflen {
		c = uint16(bytes[count])
		switch c >> 4 {
		case 0, 1, 2, 3, 4, 5, 6, 7:
			/* 0xxxxxxx*/
			count++
			chararr[chararr_count] = c
			chararr_count++
		case 12, 13:
			/* 110x xxxx   10xx xxxx*/
			count += 2
			if count > utflen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytes[count-1])
			if char2&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", count))
			}
			chararr[chararr_count] = c&0x1F<<6 | char2&0x3F
			chararr_count++
		case 14:
			/* 1110 xxxx  10xx xxxx  10xx xxxx*/
			count += 3
			if count > utflen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytes[count-2])
			char3 = uint16(bytes[count-1])
			if char2&0xC0 != 0x80 || char3&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", (count - 1)))
			}
			chararr[chararr_count] = c&0x0F<<12 | char2&0x3F<<6 | char3&0x3F<<0
			chararr_count++
		default:
			/* 10xx xxxx,  1111 xxxx */
			panic(fmt.Errorf("malformed input around byte %v", count))
		}
	}
	// The number of chars produced may be less than utflen
	chararr = chararr[0:chararr_count]
	runes := utf16.Decode(chararr)
	return string(runes)

}
