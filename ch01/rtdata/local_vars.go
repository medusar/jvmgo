package rtdata

type Slot struct {
	num int32   //用于存放数值类型
	ref *Object //用于存放引用类型
}

//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-2.html#jvms-2.6.1
type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}
