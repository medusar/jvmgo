package classfile

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.5
//MemmberInfo可以用来表示Field和Method，他们在jvm规范里的结构都一样
//不同之处就在于Attribute
/*
field_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}*/
/*method_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}*/
type MemberInfo struct {
	cp          ConstantPool
	accessFlags uint16
	//Field:The value of the name_index item must be a valid index into the constant_pool table.
	//The constant_pool entry at that index must be a CONSTANT_Utf8_info (§4.4.7) structure
	//which must represent a valid unqualified name (§4.2.2) denoting a field.
	//Method:The value of the name_index item must be a valid index into the constant_pool table.
	//The constant_pool entry at that index must be a CONSTANT_Utf8_info (§4.4.7) structure
	//representing either one of the special method names (§2.9) <init> or <clinit>,
	// or a valid unqualified name (§4.2.2) denoting a method.
	nameIndex uint16

	//Field:The value of the descriptor_index item must be a valid index into the constant_pool table.
	//The constant_pool entry at that index must be a CONSTANT_Utf8_info (§4.4.7) structure
	// that must represent a valid field descriptor (§4.3.2).
	//Method:The value of the descriptor_index item must be a valid index into the constant_pool table.
	//The constant_pool entry at that index must be a CONSTANT_Utf8_info (§4.4.7) structure
	//representing a valid method descriptor (§4.3.3).
	descriptorIndex uint16

	attributes []AttributeInfo
}

/*
 	u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
*/
func readMembers(r *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := int(r.readUint16())
	log.Println("member count:" + strconv.Itoa(memberCount))
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		log.Println("member index:" + strconv.Itoa(i))
		members[i] = readMember(r, cp)
		log.Println("member info:" + members[i].String())
	}
	return members
}

func readMember(r *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlags:     r.readUint16(),
		nameIndex:       r.readUint16(),
		descriptorIndex: r.readUint16(),
		attributes:      readAttributes(r, cp),
	}
}

//返回类型名，从常量池查找字段或方法名
func (m *MemberInfo) Name() string {
	return m.cp.getUtf8(m.nameIndex)
}

func (m *MemberInfo) Descriptor() string {
	return m.cp.getUtf8(m.descriptorIndex)
}

func (m *MemberInfo) String() string {
	s := &strings.Builder{}
	fmt.Fprintf(s, " name:%s\n", m.Name())
	fmt.Fprintf(s, " flags:%s\n", toStringAccessFlags(m.accessFlags))
	fmt.Fprintf(s, " descriptor:%s\n", m.Descriptor())
	if len(m.attributes) > 0 {
		for _, attr := range m.attributes {
			fmt.Fprintf(s, "  %s\n", attr)
		}
	}
	return s.String()
}
