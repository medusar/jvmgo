package classfile

import (
	"fmt"
	"strings"
)

//用于Deprecated和Synthetic两种类型的attribute_length字段都为0，所以实际上没有内容

//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.7.15
//u2 attribute_name_index;
// u4 attribute_length;
//The value of the attribute_length item is zero.
type DeprecatedAttribute struct{ MarkerAttribute }

//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.7.8

//u2 attribute_name_index;
// u4 attribute_length;
//The value of the attribute_length item is zero.
type SyntheticAttribute struct {
	MarkerAttribute
}

type MarkerAttribute struct{}

func (m *MarkerAttribute) readInfo(r *ClassReader) {}

func (c *DeprecatedAttribute) String() string {
	s := &strings.Builder{}
	fmt.Fprintf(s, "Deprecated\n")
	return s.String()
}

func (c *SyntheticAttribute) String() string {
	s := &strings.Builder{}
	fmt.Fprintf(s, "Synthetic\n")
	return s.String()
}
