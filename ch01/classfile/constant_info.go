package classfile

import (
	"log"
)

//常量池
//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.4

//常量flag
//cp_info中flag表示常量类型，info值随着flag的不同而不同
//golang里不推荐在命名中使用下划线，这里为了与jvm规范一致，直接从文档中拷贝过来了
const (
	CONSTANT_Class              = 7
	CONSTANT_Fieldref           = 9
	CONSTANT_Methodref          = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_String             = 8
	CONSTANT_Integer            = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_NameAndType        = 12
	CONSTANT_Utf8               = 1
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_InvokeDynamic      = 18
)

//这里为什么没有向规范里一样，定义成结构体呢?
//type ConstantInfo struct{
// flag byte
// info byte
//}
//因为根据不同的flag，info的类型也是不同的，接口更方便一些
type ConstantInfo interface {
	readInfo(r *ClassReader)
}

func readConstantInfo(r *ClassReader, cp ConstantPool) ConstantInfo {
	//读取tag的值
	tag := r.readUint8()
	log.Printf("tag:%v\n", tag)
	//根据不同的tag值，创建具体的常量
	c := newConstantInfo(tag, cp)
	//TODO:fixme
	if c == nil {
		return nil
	}
	//读取完善常量信息
	c.readInfo(r)

	return c
}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Long:
		return &ConstantLongInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{}
	case CONSTANT_String:
		return &ConstantStringInfo{cp: cp}
	case CONSTANT_Class:
		return &ConstantClassInfo{cp: cp}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{cp: cp}
	case CONSTANT_Fieldref:
		return &ConstantFieldRefInfo{ConstantRefInfo{cp: cp}}
	case CONSTANT_Methodref:
		return &ConstantMethodRefInfo{ConstantRefInfo{cp: cp}}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceMethodRefInfo{ConstantRefInfo{cp: cp}}
	case CONSTANT_MethodHandle:
		return &ConstantMethodHandleInfo{cp: cp}
	case CONSTANT_MethodType:
		return &ConstantMethodTypeInfo{cp: cp}
	case CONSTANT_InvokeDynamic:
		return &ConstantInvokeDynamicInfo{cp: cp}
	default:
		panic("java.lang.ClassFormatError:constant pool tag!")
	}
}
