package classfile

import (
	"fmt"
	"strings"
)

// java类文件结构
// https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html
type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			//由于err已经在返回值中命名，所以这里无法使用:=的形式
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html
func (c *ClassFile) read(reader *ClassReader) {
	c.readAndCheckMagic(reader)
	c.readAndCheckVersion(reader)
	c.constantPool = readConstantPool(reader)
	c.accessFlags = reader.readUint16()
	c.thisClass = reader.readUint16()
	c.superClass = reader.readUint16()
	c.interfaces = reader.readUint16s()
	c.fields = readMembers(reader, c.constantPool)
	c.methods = readMembers(reader, c.constantPool)
	c.attributes = readAttributes(reader, c.constantPool)
}

// getters
func (c *ClassFile) MajorVersion() uint16 {
	return c.majorVersion
}

func (c *ClassFile) SuperClassName() string {
	if c.superClass > 0 {
		return c.constantPool.getClassName(c.superClass)
	}
	return "" //java.lang.Object没有父类
}

func (c *ClassFile) InterfaceNames() []string {
	names := make([]string, len(c.interfaces))
	for i, cpIndex := range c.interfaces {
		names[i] = c.constantPool.getClassName(cpIndex)
	}
	return names
}

//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-5.html#jvms-5.3
func (c *ClassFile) readAndCheckMagic(reader *ClassReader) {
	c.magic = reader.readUint32()
	if c.magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

func (c *ClassFile) readAndCheckVersion(reader *ClassReader) {
	c.minorVersion = reader.readUint16()
	c.majorVersion = reader.readUint16()
	switch c.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if c.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

func (c *ClassFile) String() string {
	s := &strings.Builder{}
	fmt.Fprintf(s, "version: %v.%v\n", c.majorVersion, c.minorVersion)
	fmt.Fprintf(s, "constants count: %v\n", len(c.constantPool))
	fmt.Fprintf(s, "access flags: 0x%x\n", c.accessFlags)
	fmt.Fprintf(s, "this class: %v\n", c.thisClass)
	fmt.Fprintf(s, "super class: %v\n", c.superClass)
	fmt.Fprintf(s, "interfaces: %v\n", c.interfaces)
	fmt.Fprintf(s, "fields count: %v\n", len(c.fields))
	for _, f := range c.fields {
		fmt.Fprintf(s, "  %s\n", f.Name())
	}
	fmt.Fprintf(s, "methods count: %v\n", len(c.methods))
	for _, m := range c.methods {
		fmt.Fprintf(s, "  %s\n", m.Name())
	}
	return s.String()
}
