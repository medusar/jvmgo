package classfile

//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.7.5
// Exceptions_attribute {
//     u2 attribute_name_index;
//     u4 attribute_length;
//     u2 number_of_exceptions;
//     u2 exception_index_table[number_of_exceptions];
// }
type ExceptionsAttribute struct {
	exceptionIndexTable []uint16 //u2类型
}

func (e *ExceptionsAttribute) readInfo(r *ClassReader) {
	// count := r.readUint16()
	// e.exceptionIndexTable = make([]uint16, count)
	// for i := range e.exceptionIndexTable {
	// 	e.exceptionIndexTable[i] = r.readUint16()
	// }
	e.exceptionIndexTable = r.readUint16s()
}

//getter
func (e *ExceptionsAttribute) ExceptionIndexTable() []uint16 {
	return e.exceptionIndexTable
}
