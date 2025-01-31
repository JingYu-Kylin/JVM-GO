package references

import "reflect"
import "JVM-GO/ch10/instructions/base"
import "JVM-GO/ch10/rtda"
import "JVM-GO/ch10/rtda/heap"

// Throw exception or error
// athrow指令的操作数是一个异常对象引用，从操作数栈弹出。
type ATHROW struct{ base.NoOperandsInstruction }

func (self *ATHROW) Execute(frame *rtda.Frame) {
	// 先从操作数栈中弹出异常对象引用，
	// 如果该引用是null，则抛出NullPointerException异常，
	// 否则看是否可以找到并跳转到异常处理代码。
	ex := frame.OperandStack().PopRef()
	if ex == nil {
		panic("java.lang.NullPointerException")
	}

	thread := frame.Thread()
	if !findAndGotoExceptionHandler(thread, ex) {
		handleUncaughtException(thread, ex)
	}
}

func findAndGotoExceptionHandler(thread *rtda.Thread, ex *heap.Object) bool {
	// 从当前帧开始，遍历Java虚拟机栈，查找方法的异常处理表。
	// 假设遍历到帧F，如果在F对应的方法中找不到异常处理项，则把F弹出，继续遍历。
	// 反之如果找到了异常处理项，在跳转到异常处理代码之前，要先把F的操作数栈清空，然后把异常对象引用推入栈顶。
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC() - 1

		handlerPC := frame.Method().FindExceptionHandler(ex.Class(), pc)
		if handlerPC > 0 {
			stack := frame.OperandStack()
			stack.Clear()
			stack.PushRef(ex)
			frame.SetNextPC(handlerPC)
			return true
		}

		thread.PopFrame()
		if thread.IsStackEmpty() {
			break
		}
	}
	return false
}

// todo 如果遍历完Java虚拟机栈还是找不到异常处理代码，则handleUncaughtException（）函数打印出Java虚拟机栈信息
func handleUncaughtException(thread *rtda.Thread, ex *heap.Object) {
	thread.ClearStack()

	jMsg := ex.GetRefVar("detailMessage", "Ljava/lang/String;")
	goMsg := heap.GoString(jMsg)
	println(ex.Class().JavaName() + ": " + goMsg)

	stes := reflect.ValueOf(ex.Extra())
	for i := 0; i < stes.Len(); i++ {
		ste := stes.Index(i).Interface().(interface {
			String() string
		})
		println("\tat " + ste.String())
	}
	// 把Java虚拟机栈清空，然后打印出异常信息。
	// 由于Java虚拟机栈已经空了，所以解释器也就终止执行了。
}
