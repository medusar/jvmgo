package rtdata

//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-2.html#jvms-2.6.2
type OperandStack struct {
	size  uint
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}
