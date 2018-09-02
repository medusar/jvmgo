package classfile

import (
	"fmt"
	"strings"
)

//	https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.7.4
/*
StackMapTable_attribute {
    u2              attribute_name_index;
    u4              attribute_length;
    u2              number_of_entries;
    stack_map_frame entries[number_of_entries];
}
*/
type StackMapTableAttribute struct {
	entries []*StackMapFrameEntry
}

/*
Frame类型
union stack_map_frame {
    same_frame;
    same_locals_1_stack_item_frame;
    same_locals_1_stack_item_frame_extended;
    chop_frame;
    same_frame_extended;
    append_frame;
    full_frame;
}
*/
type StackMapFrameEntry struct {
}

//TODO:比较多，有时间再一个一个实现
func readStackMapFrameEntry(r *ClassReader) *StackMapFrameEntry {
	return &StackMapFrameEntry{}
}

func (s *StackMapTableAttribute) readInfo(r *ClassReader) {
	size := r.readUint16()
	s.entries = make([]*StackMapFrameEntry, size)
	for i := range s.entries {
		s.entries[i] = readStackMapFrameEntry(r)
	}
}

// StackMapTable: number_of_entries = 4
// frame_type = 16 /* same */
// frame_type = 9 /* same */
// frame_type = 9 /* same */
// frame_type = 7 /* same */
func (s *StackMapTableAttribute) String() string {
	str := &strings.Builder{}
	fmt.Fprintf(str, "    StackMapTable: number_of_entries=%d\n", len(s.entries))
	for _, frame := range s.entries {
		//TODO
		fmt.Fprintf(str, "        frame_type = %d\n", frame)
	}
	return str.String()
}
