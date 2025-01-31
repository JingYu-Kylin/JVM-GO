package heap

type Object struct {
	class *Class
	data  interface{} // Slots for Object, []int32 for int[] ...
	// 把fields字段改为data，类型也从Slots变成了interface{}。
	// Go语言的interface{}类型很像C语言中的void*，该类型的变量可以容纳任何类型的值。
	// 对于普通对象来说，data字段中存放的仍然还是Slots变量。
	// 但是对于数组，可以在其中放各种类型的数组
}

// create normal (non-array) object
func newObject(class *Class) *Object {
	return &Object{
		class: class,
		data:  newSlots(class.instanceSlotCount),
	}
}

// getters
func (self *Object) Class() *Class {
	return self.class
}
func (self *Object) Fields() Slots {
	return self.data.(Slots)
}

func (self *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(self.class)
}

// reflection
func (self *Object) GetRefVar(name, descriptor string) *Object {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	return slots.GetRef(field.slotId)
}

// 直接给对象的引用类型实例变量赋值
func (self *Object) SetRefVar(name, descriptor string, ref *Object) {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	slots.SetRef(field.slotId, ref)
}
