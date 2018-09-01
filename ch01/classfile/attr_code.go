package classfile

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
	c.maxStack = r.readUint16()
	c.maxLocals = r.readUint16()

	codeLength := r.readUint32()
	c.code = r.readBytes(codeLength)

	exceptionTableCount := r.readUint32()
	c.exceptionTable = make([]*ExceptionTableEntry, exceptionTableCount)
	for i := range c.exceptionTable {
		c.exceptionTable[i] = &ExceptionTableEntry{
			startPc:   r.readUint16(),
			endPc:     r.readUint16(),
			handlerPc: r.readUint16(),
			catchType: r.readUint16(),
		}
	}

	attributeCount := r.readUint16()
	c.attributes = make([]AttributeInfo, attributeCount)
	for i := range c.attributes {
		c.attributes[i] = readAttribute(r, c.cp)
	}
}
