package constants

import "JVM-GO/ch09/instructions/base"
import "JVM-GO/ch09/rtda"
import "JVM-GO/ch09/rtda/heap"

// Push item from run-time constant pool
type LDC struct{ base.Index8Instruction }

func (self *LDC) Execute(frame *rtda.Frame) {
	_ldc(frame, self.Index)
}

// Push item from run-time constant pool (wide index)
type LDC_W struct{ base.Index16Instruction }

func (self *LDC_W) Execute(frame *rtda.Frame) {
	_ldc(frame, self.Index)
}

func _ldc(frame *rtda.Frame, index uint) {
	stack := frame.OperandStack()
	class := frame.Method().Class()
	c := class.ConstantPool().GetConstant(index)

	switch c.(type) {
	case int32:
		stack.PushInt(c.(int32))
	case float32:
		stack.PushFloat(c.(float32))
	case string:
		internedStr := heap.JString(class.Loader(), c.(string))
		stack.PushRef(internedStr)
	case *heap.ClassRef:
		// 如果运行时，常量池中的常量是类引用，则解析类引用，然后把类的类对象推入操作数栈顶。
		classRef := c.(*heap.ClassRef)
		classObj := classRef.ResolvedClass().JClass()
		stack.PushRef(classObj)
	// case MethodType, MethodHandle
	default:
		panic("todo: ldc!")
	}
}

// Push long or double from run-time constant pool (wide index)
type LDC2_W struct{ base.Index16Instruction }

func (self *LDC2_W) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(self.Index)

	switch c.(type) {
	case int64:
		stack.PushLong(c.(int64))
	case float64:
		stack.PushDouble(c.(float64))
	default:
		panic("java.lang.ClassFormatError")
	}
}
