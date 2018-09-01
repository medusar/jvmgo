package classfile

//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.5
//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.7
//Attribute也有很多种不同的实现，每种实现是通过attribute_name_index来区分
// ClassFile中Attribute的定义
// u2             attributes_count;
// attribute_info attributes[attributes_count];
type AttributeInfo interface {
	readInfo(r *ClassReader)
}

//读取ClassFile中的所有Attribute
func readAttributes(r *ClassReader, cp ConstantPool) []AttributeInfo {
	count := r.readUint16()
	attrs := make([]AttributeInfo, count)
	for i := range attrs {
		attrs[i] = readAttribute(r, cp)
	}
	return attrs
}

//attribute_info {
// u2 attribute_name_index;
// u4 attribute_length;
// u1 info[attribute_length];
// }
func readAttribute(r *ClassReader, cp ConstantPool) AttributeInfo {
	nameIndex := r.readUint16()
	attrName := cp.getUtf8(nameIndex)
	attrLengh := r.readUint32()
	return newAttributeInfo(attrName, attrLengh, cp)
}

//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.7
//attribute_info {
// u2 attribute_name_index;
// u4 attribute_length;
// u1 info[attribute_length];
// }

//attribute_name
// Table 4.6. Predefined class file attributes
// Attribute	Section	Java SE	class file
// ConstantValue	§4.7.2	1.0.2	45.3
// Code	§4.7.3	1.0.2	45.3
// StackMapTable	§4.7.4	6	50.0
// Exceptions	§4.7.5	1.0.2	45.3
// InnerClasses	§4.7.6	1.1	45.3
// EnclosingMethod	§4.7.7	5.0	49.0
// Synthetic	§4.7.8	1.1	45.3
// Signature	§4.7.9	5.0	49.0
// SourceFile	§4.7.10	1.0.2	45.3
// SourceDebugExtension	§4.7.11	5.0	49.0
// LineNumberTable	§4.7.12	1.0.2	45.3
// LocalVariableTable	§4.7.13	1.0.2	45.3
// LocalVariableTypeTable	§4.7.14	5.0	49.0
// Deprecated	§4.7.15	1.1	45.3
// RuntimeVisibleAnnotations	§4.7.16	5.0	49.0
// RuntimeInvisibleAnnotations	§4.7.17	5.0	49.0
// RuntimeVisibleParameterAnnotations	§4.7.18	5.0	49.0
// RuntimeInvisibleParameterAnnotations	§4.7.19	5.0	49.0
// AnnotationDefault	§4.7.20	5.0	49.0
// BootstrapMethods	§4.7.21	7	51.0
func newAttributeInfo(name string, length uint32, cp ConstantPool) AttributeInfo {
	switch name {
	case "Code":
		return nil
	case "ConstantValue":
		return nil
	default:
		return nil
	}
}
