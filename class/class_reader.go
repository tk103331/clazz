package class

import (
	"github.com/tk103331/clazz/data"
	"io"
)

type ClassReader struct {
	r  *data.DataReader
	cv ClassVisitor
	d  *Data
}

func NewReader(reader io.Reader) *ClassReader {
	return &ClassReader{r: data.NewReader(reader), d: &Data{}}
}

func (cr *ClassReader) SetVisitor(visitor ClassVisitor) {
	cr.cv = visitor
}

func (cr *ClassReader) Read() error {
	cr.d.MagicNumber = cr.readU4()
	cr.d.MinorVersion = cr.readU2()
	cr.d.MajorVersion = cr.readU2()

	cr.d.ConstantCount = cr.readU2()
	cr.d.ConstantPool = cr.readConstantPool(cr.d.ConstantCount)

	cr.d.AccessFlags = cr.readU2()
	cr.d.ThisClass = cr.readU2()
	cr.d.SuperClass = cr.readU2()

	cr.d.FieldsCount = cr.readU2()
	cr.d.Fields = cr.readFields(cr.d.FieldsCount)

	cr.d.MethodsCount = cr.readU2()
	cr.d.Methods = cr.readMethods(cr.d.MethodsCount)

	cr.d.AttributesCount = cr.readU2()
	cr.d.Attributes = cr.readAttributes(cr.d.AttributesCount)
	return nil
}

func (cr *ClassReader) readConstantPool(count uint16) []Constant {
	pool := make([]Constant, count)
	pool[0] = nil
	for i := uint16(1); i < count; i++ {
		tag := cr.readU1()
		switch tag {
		case TAG_CONSTANT_UTF8:
			length := cr.readU2()
			str := cr.readUTF8(int(length))
			pool[i] = ConstantUTF8{Length: length, Value: str}
		case TAG_CONSTANT_INTEGER:
			integer := cr.readInt32()
			pool[i] = ConstantInteger{Value: integer}
		case TAG_CONSTANT_FLOAT:
			float := cr.readFloat32()
			pool[i] = ConstantFloat{Value: float}
		case TAG_CONSTANT_LONG:
			long := cr.readInt64()
			pool[i] = ConstantLong{Value: long}
		case TAG_CONSTANT_DOUBLE:
			double := cr.readFloat64()
			pool[i] = ConstantDouble{Value: double}
		case TAG_CONSTANT_CLASS:
			index := cr.readU2()
			pool[i] = ConstantClass{Index: index}
		case TAG_CONSTANT_STRING:
			index := cr.readU2()
			pool[i] = ConstantString{Index: index}
		case TAG_CONSTANT_FIELDREF:
			classIndex := cr.readU2()
			fieldIndex := cr.readU2()
			pool[i] = ConstantFieldRef{ClassIndex: classIndex, NameAndTypeIndex: fieldIndex}
		case TAG_CONSTANT_METHODREF:
			classIndex := cr.readU2()
			methodIndex := cr.readU2()
			pool[i] = ConstantMethodRef{ClassIndex: classIndex, NameAndTypeIndex: methodIndex}
		case TAG_CONSTANT_IFMETHODREF:
			interfaceIndex := cr.readU2()
			methodIndex := cr.readU2()
			pool[i] = ConstantIfMethodRef{InterfaceIndex: interfaceIndex, NameAndTypeIndex: methodIndex}
		case TAG_CONSTANT_NAMEANDTYPE:
			nameIndex := cr.readU2()
			typeIndex := cr.readU2()
			pool[i] = ConstantNameAndType{NameIndex: nameIndex, TypeIndex: typeIndex}
		}
	}
	return pool
}

func (cr *ClassReader) readU1() uint8 {
	v, _ := cr.r.ReadUint8()
	return v
}
func (cr *ClassReader) readU2() uint16 {
	v, _ := cr.r.ReadUint16()
	return v
}
func (cr *ClassReader) readU4() uint32 {
	v, _ := cr.r.ReadUint32()
	return v
}
func (cr *ClassReader) readInt32() int32 {
	v, _ := cr.r.ReadInt32()
	return v
}
func (cr *ClassReader) readInt64() int64 {
	v, _ := cr.r.ReadInt64()
	return v
}
func (cr *ClassReader) readFloat32() float32 {
	v, _ := cr.r.ReadFloat32()
	return v
}
func (cr *ClassReader) readFloat64() float64 {
	v, _ := cr.r.ReadFloat64()
	return v
}

func (cr *ClassReader) readUTF8(length int) string {
	str, _ := cr.r.ReadString(length)
	return str
}

func (cr *ClassReader) readBytes(length int) []byte {
	bytes, _ := cr.r.ReadBytes(length)
	return bytes
}

func (cr *ClassReader) readInterfaces(count uint16) []Interface {
	interfaces := make([]Interface, count)
	for i := uint16(0); i < count; i++ {
		index := cr.readU2()
		interfaces[i] = Interface{Index: index}
	}
	return interfaces
}

func (cr *ClassReader) readFields(count uint16) []Field {
	fields := make([]Field, count)
	for i := uint16(0); i < count; i++ {
		f := Field{}
		f.AccessFlags = cr.readU2()
		f.NameIndex = cr.readU2()
		f.DescriptorIndex = cr.readU2()
		f.AttributesCount = cr.readU2()
		f.Attributes = cr.readAttributes(f.AttributesCount)
		fields[i] = f
	}
	return fields
}

func (cr *ClassReader) readMethods(count uint16) []Method {
	methods := make([]Method, count)
	for i := uint16(0); i < count; i++ {
		m := Method{}
		m.AccessFlags = cr.readU2()
		m.NameIndex = cr.readU2()
		m.DescriptorIndex = cr.readU2()
		m.AttributesCount = cr.readU2()
		m.Attributes = cr.readAttributes(m.AttributesCount)
		methods[i] = m
	}
	return methods
}

func (cr *ClassReader) readAttributes(count uint16) []Attribute {
	attributes := make([]Attribute, count)
	for i := uint16(0); i < count; i++ {
		a := Attribute{}
		a.NameIndex = cr.readU2()
		a.Length = cr.readU1()
		a.Info = cr.readBytes(int(a.Length))
		attributes[i] = a
	}
	return attributes
}
