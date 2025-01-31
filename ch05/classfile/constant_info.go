package classfile

/**
常量
 */

const (
	// Java虚拟机规范一共定义了14种常量
	CONSTANT_Class              = 7
	CONSTANT_Fieldref           = 9
	CONSTANT_Methodref          = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_String             = 8
	CONSTANT_Integer            = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_NameAndType        = 12
	CONSTANT_Utf8               = 1
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_InvokeDynamic      = 18
)

/*
Java虚拟机规范给出的常量结构
cp_info {
    u1 tag;
    u1 info[];
}
*/
type ConstantInfo interface {
	// 读取常量信息
	readInfo(reader *ClassReader)
}

/**
 *
 */
func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	// 先读出tag值
	tag := reader.ReadUint8()
	// 然后调用newConstantInfo（）函数创建具体的常量
	c := newConstantInfo(tag, cp)
	// 最后调用常量的readInfo（）方法读取常量信息
	c.readInfo(reader)
	return c
}

/**
 * 根据tag值创建具体的常量
 */
func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Integer: return &ConstantIntegerInfo{}
	case CONSTANT_Float: return &ConstantFloatInfo{}
	case CONSTANT_Long: return &ConstantLongInfo{}
	case CONSTANT_Double: return &ConstantDoubleInfo{}
	case CONSTANT_Utf8: return &ConstantUtf8Info{}
	case CONSTANT_String: return &ConstantStringInfo{cp: cp}
	case CONSTANT_Class: return &ConstantClassInfo{cp: cp}
	case CONSTANT_Fieldref:
		return &ConstantFieldrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_Methodref:
		return &ConstantMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_NameAndType: return &ConstantNameAndTypeInfo{}
	case CONSTANT_MethodType: return &ConstantMethodTypeInfo{}
	case CONSTANT_MethodHandle: return &ConstantMethodHandleInfo{}
	case CONSTANT_InvokeDynamic: return &ConstantInvokeDynamicInfo{}
	default: panic("java.lang.ClassFormatError: constant pool tag!")
	}
}
