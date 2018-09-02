package classfile

import (
	"fmt"
	"strings"
)

//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.7.2
type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (c *ConstantValueAttribute) readInfo(r *ClassReader) {
	c.constantValueIndex = r.readUint16()
}

//getter
func (c *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return c.constantValueIndex
}
func (c *ConstantValueAttribute) String() string {
	s := &strings.Builder{}
	fmt.Fprintf(s, "ConstantValue:\n")
	fmt.Fprintf(s, "  constantValueIndex=%d\n", c.constantValueIndex)
	return s.String()
}
