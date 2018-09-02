package classfile

import (
	"encoding/binary"
)

//用于读取class文件字节
type ClassReader struct {
	data  []byte
	index int
}

func (c *ClassReader) readUint8() uint8 {
	val := c.data[0]
	c.data = c.data[1:]
	c.index++
	return val
}
func (c *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(c.data)
	//uint16需要两个字节，所以取[2:]
	c.data = c.data[2:]
	c.index += 2
	return val
}
func (c *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(c.data)
	c.data = c.data[4:]
	c.index += 4
	return val
}
func (c *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(c.data)
	c.data = c.data[8:]
	c.index += 8
	return val
}
func (c *ClassReader) readUint16s() []uint16 {
	n := c.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = c.readUint16()
	}

	return s
}
func (c *ClassReader) readBytes(n uint32) []byte {
	bytes := c.data[:n]
	c.data = c.data[n:]
	c.index += int(n)
	return bytes
}

// 返回读取index
func (c *ClassReader) Index() int {
	return c.index
}
