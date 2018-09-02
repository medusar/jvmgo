package classfile

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
)

//常量池，其实就是常量信息表，这里用数组表示
//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.4
type ConstantPool []ConstantInfo

//ClassFile中定义
//u2             constant_pool_count;
// cp_info        constant_pool[constant_pool_count-1];
func readConstantPool(r *ClassReader) ConstantPool {
	//constant_pool_count
	// The value of the constant_pool_count item is equal to the number of entries
	// in the constant_pool table plus one.
	//A constant_pool index is considered valid if it is greater than zero
	//and less than constant_pool_count,
	//with the exception for constants of type long and double noted in §4.4.5.
	count := int(r.readUint16())
	log.Println("constant info size:" + strconv.Itoa(count))
	cp := make([]ConstantInfo, count)
	// The constant_pool table is indexed from 1 to constant_pool_count - 1.
	for i := 1; i < count; i++ {
		log.Println("constant info index:" + strconv.Itoa(i))
		cp[i] = readConstantInfo(r, cp)
		//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.4.5
		//The CONSTANT_Long_info and CONSTANT_Double_info represent 8-byte numeric (long and double) constants
		//All 8-byte constants take up two entries in the constant_pool table of the class file.
		//If a CONSTANT_Long_info or CONSTANT_Double_info structure is the item in the constant_pool table at index n,
		//then the next usable item in the pool is located at index n+2.
		//The constant_pool index n+1 must be valid but is considered unusable.
		switch cp[i].(type) { //类似反射的语法，获取对象类型
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return cp
}

//TODO:看下jvm规范里哪里提到的
func (c ConstantPool) getUtf8(index uint16) string {
	info := c.getConstantInfo(index).(*ConstantUtf8Info)
	return info.str
}

//TODO:看下哪里用到以及JVM规范如何定义的
func (c ConstantPool) getClassName(index uint16) string {
	info := c.getConstantInfo(index).(*ConstantClassInfo)
	return c.getUtf8(info.nameIndex)
}

//按索引查找常量
func (c ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	log.Println("cp pool size:" + strconv.Itoa(len(c)) + ", index:" + strconv.Itoa(int(index)))
	if cpInfo := c[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

//按索引获取常量名字和类型
func (c ConstantPool) getNameAndType(index uint16) (string, string) {
	nameAndTypeInfo := c.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	//上一步如果没查到会panic，所以这里不需要判断是否查到
	name := c.getUtf8(nameAndTypeInfo.nameIndex)
	_type := c.getUtf8(nameAndTypeInfo.descriptorIndex)
	return name, _type
}

func (c ConstantPool) String() string {
	s := &strings.Builder{}
	for i, cpInfo := range c {
		fmt.Fprintf(s, "ConstantPool[%d]: %v\n", i, reflect.TypeOf(cpInfo))
	}
	return s.String()
}
