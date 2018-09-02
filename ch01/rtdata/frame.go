package rtdata

//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-2.html#jvms-2.6

type Frame struct {
	lower        *Frame        //栈中的底下一个帧
	localVars    LocalVars     //本地变量表,因为是个数组，所以无需*
	operandStack *OperandStack //操作数栈
}
