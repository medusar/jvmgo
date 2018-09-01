package classfile

import (
	"math"
)

//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.4.4

// CONSTANT_Integer_info {
//     u1 tag;
//     u4 bytes;
// }

// CONSTANT_Float_info {
//     u1 tag;
//     u4 bytes;
// }
type ConstantIntegerInfo struct {
	val int32
}

//实现ConstantInfo接口
func (i *ConstantIntegerInfo) readInfo(r *ClassReader) {
	//java中的int是有符号的，所以这里需要把无符号的转为有符号的
	//doc:The bytes item of the CONSTANT_Integer_info structure
	//represents the value of the int constant.
	//The bytes of the value are stored in big-endian (high byte first) order.
	i.val = int32(r.readUint32())
}

//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.4.4
type ConstantFloatInfo struct {
	//The bytes item of the CONSTANT_Float_info structure
	//represents the value of the float constant in IEEE 754 floating-point single format (§2.3.2).
	//The bytes of the single format representation are stored in big-endian (high byte first) order.
	val float32
}

func (f *ConstantFloatInfo) readInfo(r *ClassReader) {
	//java中float是32位
	//nt in IEEE 754 floating-point single format
	f.val = math.Float32frombits(r.readUint32())
}

//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.4.5
type ConstantLongInfo struct {
	val int64
}

func (l *ConstantLongInfo) readInfo(r *ClassReader) {
	l.val = int64(r.readUint64())
}

type ConstantDoubleInfo struct {
	val float64
}

func (d *ConstantDoubleInfo) readInfo(r *ClassReader) {
	d.val = math.Float64frombits(r.readUint64())
}
