package classfile

// https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.7.13
// LocalVariableTable_attribute {
//     u2 attribute_name_index;
//     u4 attribute_length;
//     u2 local_variable_table_length;
//     {   u2 start_pc;
//         u2 length;
//         u2 name_index;
//         u2 descriptor_index;
//         u2 index;
//     } local_variable_table[local_variable_table_length];
// }
type LocalVariableTableAttribute struct {
	localVariableTable []*LocalVaritableTableEntry
}
type LocalVaritableTableEntry struct {
	startPc         uint16
	length          uint16
	nameIndex       uint16
	descriptorIndex uint16
	index           uint16
}

func (l *LocalVariableTableAttribute) readInfo(r *ClassReader) {
	count := r.readUint16()
	l.localVariableTable = make([]*LocalVaritableTableEntry, count)
	for i := range l.localVariableTable {
		l.localVariableTable[i] = &LocalVaritableTableEntry{
			startPc:         r.readUint16(),
			length:          r.readUint16(),
			nameIndex:       r.readUint16(),
			descriptorIndex: r.readUint16(),
			index:           r.readUint16(),
		}
	}
}

func (l *LocalVariableTableAttribute) String() string {
	return "TODO"
}
