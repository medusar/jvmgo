package classfile

import (
	"fmt"
	"strings"
)

//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.7.12
// LineNumberTable_attribute {
//     u2 attribute_name_index;
//     u4 attribute_length;
//     u2 line_number_table_length;
//     {   u2 start_pc;
//         u2 line_number;
//     } line_number_table[line_number_table_length];
// }
type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (l *LineNumberTableAttribute) readInfo(r *ClassReader) {
	count := r.readUint16()
	l.lineNumberTable = make([]*LineNumberTableEntry, count)
	for i := range l.lineNumberTable {
		l.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    r.readUint16(),
			lineNumber: r.readUint16(),
		}
	}
}

func (l *LineNumberTableAttribute) String() string {
	s := &strings.Builder{}
	fmt.Fprintf(s, "LineNumberTable:\n")
	for _, entry := range l.lineNumberTable {
		fmt.Fprintf(s, "  line %d:%d", entry.lineNumber, entry.startPc)
	}
	return s.String()
}
