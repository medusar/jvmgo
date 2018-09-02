package classfile

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.7.3
// Code_attribute {
//     u2 attribute_name_index;
//     u4 attribute_length;
//     u2 max_stack;
//     u2 max_locals;
//     u4 code_length;
//     u1 code[code_length];
//     u2 exception_table_length;
//     {   u2 start_pc;
//         u2 end_pc;
//         u2 handler_pc;
//         u2 catch_type;
//     } exception_table[exception_table_length];
//     u2 attributes_count;
//     attribute_info attributes[attributes_count];
// }
type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func (c *CodeAttribute) readInfo(r *ClassReader) {
	log.Println("Code begin, readerIndex:" + strconv.Itoa(r.Index()))
	c.maxStack = r.readUint16()
	c.maxLocals = r.readUint16()

	codeLength := r.readUint32()
	c.code = r.readBytes(codeLength)

	exceptionTableCount := r.readUint16()
	c.exceptionTable = make([]*ExceptionTableEntry, exceptionTableCount)
	for i := range c.exceptionTable {
		c.exceptionTable[i] = &ExceptionTableEntry{
			startPc:   r.readUint16(),
			endPc:     r.readUint16(),
			handlerPc: r.readUint16(),
			catchType: r.readUint16(),
		}
	}
	c.attributes = readAttributes(r, c.cp)

	log.Println("Code end, readerIndex:" + strconv.Itoa(r.Index()))
}

func (c *CodeAttribute) String() string {
	s := &strings.Builder{}
	fmt.Fprintf(s, "Code:\n")
	fmt.Fprintf(s, "    stack=%d, locals=%d, args_size=%d\n", c.maxStack, c.maxLocals, len(c.code)+len(c.exceptionTable))
	for _, attr := range c.attributes {
		fmt.Fprintf(s, "  %s\n", attr.String())
	}
	return s.String()
}
