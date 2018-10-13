# jvmgo
simple jvm written in go, based on a tutorial, for learning go only!

1. The tutorial book is called 《自己动手写Java虚拟机》(Build your own JVM) can be found [here](https://item.jd.com/11935272.html)
2. All the codes in this project is coded by myself, based on the instructions and demos, with little change anyway, mainly for practice.

3. JVM Specification: [Java7](https://docs.oracle.com/javase/specs/jls/se7/html/index.html)

# How to run?
1. `go install jvmgo/ch01`, you should make sure codes could be found in $GOPATH or $GOROOT.
2. goto directory $GOPATH/bin
3. `./ch01 {parameters}`

# TODO list
- [ ] read class file and display detailed information.
- [ ] refactor code, more encapsulated.