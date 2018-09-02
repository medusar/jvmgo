package classfile

//https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.7.10
type SourceFileAttribute struct {
	cp ConstantPool
	// nameIndex       uint16  //nameIndex和lengthIndex是已知的(AttributeInfo类文件里)，所以不再需要在里面定义
	// lengthIndex     uint32
	sourceFileIndex uint16
}

func (s *SourceFileAttribute) readInfo(r *ClassReader) {
	s.sourceFileIndex = r.readUint16()
}

//TODO:why会需要这样一个方法呢？JVM规范里好像没有要求
func (s *SourceFileAttribute) FileName() string {
	return s.cp.getUtf8(s.sourceFileIndex)
}

func (s *SourceFileAttribute) String() string {
	return "TODO"
}
