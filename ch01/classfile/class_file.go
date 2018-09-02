package classfile

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

/* java类文件结构
https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html
*/
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
				log.Fatal(err)
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{data: classData, index: 0}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html
func (c *ClassFile) read(reader *ClassReader) {
	c.readAndCheckMagic(reader)
	log.Println("magic number:" + strconv.Itoa(int(c.magic)))
	c.readAndCheckVersion(reader)
	log.Println("minor version:" + strconv.Itoa(int(c.minorVersion)))
	log.Println("major version:" + strconv.Itoa(int(c.majorVersion)))
	c.constantPool = readConstantPool(reader)
	log.Println("constant pool finished")
	log.Println(c.constantPool)
	c.accessFlags = reader.readUint16()
	log.Println("access flag finished:" + toStringAccessFlags(c.accessFlags))
	c.thisClass = reader.readUint16()
	log.Println("this class finished")
	c.superClass = reader.readUint16()
	log.Println("super class finished")
	c.interfaces = reader.readUint16s()
	log.Println("interfaces finished")
	c.fields = readMembers(reader, c.constantPool)
	log.Println("fields finished")
	c.methods = readMembers(reader, c.constantPool)
	log.Println("methods finished")
	c.attributes = readAttributes(reader, c.constantPool)
	log.Println("attributes finished")
}

/*
ACC_PUBLIC	0x0001	Declared public; may be accessed from outside its package.
ACC_FINAL	0x0010	Declared final; no subclasses allowed.
ACC_SUPER	0x0020	Treat superclass methods specially when invoked by the invokespecial instruction.
ACC_INTERFACE	0x0200	Is an interface, not a class.
ACC_ABSTRACT	0x0400	Declared abstract; must not be instantiated.
ACC_SYNTHETIC	0x1000	Declared synthetic; not present in the source code.
ACC_ANNOTATION	0x2000	Declared as an annotation type.
ACC_ENUM	0x4000
*/
//TODO:不同类型的accessFlags不同
func toStringAccessFlags(flags uint16) string {
	names := make([]string, 1)
	if 0x0001&flags == 0x0001 {
		names = append(names, "ACC_PUBLIC")
	}
	if 0x0010&flags == 0x0010 {
		names = append(names, "ACC_FINAL")
	}
	if 0x0020&flags == 0x0020 {
		names = append(names, "ACC_SUPER")
	}
	if 0x0200&flags == 0x0200 {
		names = append(names, "ACC_INTERFACE")
	}
	if 0x0400&flags == 0x0400 {
		names = append(names, "ACC_ABSTRACT")
	}
	if 0x1000&flags == 0x1000 {
		names = append(names, "ACC_SYNTHETIC")
	}
	if 0x2000&flags == 0x2000 {
		names = append(names, "ACC_ANNOTATION")
	}
	if 0x4000&flags == 0x4000 {
		names = append(names, "ACC_ENUM")
	}
	return strings.Join(names, ", ")
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
