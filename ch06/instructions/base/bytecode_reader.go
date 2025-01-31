package base

type BytecodeReader struct {
	code []byte // 存放字节码
	pc int // 记录读取到了哪个字节
}

/**
为了避免每次解码指令都新创建一个BytecodeReader实例
 */
func (self *BytecodeReader) Reset(code []byte, pc int) {
	self.code = code
	self.pc = pc
}

func (self *BytecodeReader) PC() int {
	return self.pc
}

/**
实现一系列的Read（）方法
 */
func (self *BytecodeReader) ReadUint8() uint8 {
	i := self.code[self.pc]
	self.pc++
	return i
}
func (self *BytecodeReader) ReadInt8() int8 {
	// 调用ReadUint8（），然后把读取到的值转成int8返回
	return int8(self.ReadUint8())
}
func (self *BytecodeReader) ReadUint16() uint16 {
	// 连续读取两字节
	byte1 := uint16(self.ReadUint8())
	byte2 := uint16(self.ReadUint8())
	return (byte1 << 8) | byte2
}
func (self *BytecodeReader) ReadInt16() int16 {
	// 调用ReadUint16（），然后把读取到的值转成int16返回
	return int16(self.ReadUint16())
}
func (self *BytecodeReader) ReadInt32() int32 {
	// 连续读取4字节
	byte1 := int32(self.ReadUint8())
	byte2 := int32(self.ReadUint8())
	byte3 := int32(self.ReadUint8())
	byte4 := int32(self.ReadUint8())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}

/**
这两个方法只有tableswitch和lookupswitch指令使用
 */
func (self *BytecodeReader) ReadInt32s(n int32) []int32 {
	ints := make([]int32, n)
	for i := range ints {
		ints[i] = self.ReadInt32()
	}
	return ints
}
func (self *BytecodeReader) SkipPadding() {
	for self.pc%4 != 0 {
		self.ReadUint8()
	}
}







