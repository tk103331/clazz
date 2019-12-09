package data

import (
	"github.com/tk103331/clazz/data"
	"io"
)

type Reader struct {
	reader *data.DataReader
	data   *ClassData
}

func NewReader(reader io.Reader) *Reader {
	return &Reader{reader: data.NewReader(reader), data: &ClassData{}}
}

func (r *Reader) Accept(visitor Visitor) {
	if visitor != nil {
		visitor.VisitStart()
		visitor.VisitMagicNumber(r.data.MagicNumber)
		visitor.VisitVersion(r.data.MinorVersion, r.data.MajorVersion)
		visitor.VisitConstants(r.data.ConstantPool)
		visitor.Visit(r.data.ThisClass, r.data.SuperClass, r.data.AccessFlags)
		visitor.VisitInterfaces(r.data.Interfaces)
		visitor.VisitFields(r.data.Fields)
		visitor.VisitMethods(r.data.Methods)
		visitor.VisitAttributes(r.data.Attributes)
		visitor.VisitEnd()
	}
}

func (r *Reader) Read() error {
	r.data.MagicNumber = r.readU4()
	r.data.MinorVersion = r.readU2()
	r.data.MajorVersion = r.readU2()

	r.data.ConstantCount = r.readU2()
	r.data.ConstantPool = r.readConstantPool(r.data.ConstantCount)

	r.data.AccessFlags = r.readU2()
	r.data.ThisClass = r.readU2()
	r.data.SuperClass = r.readU2()

	r.data.InterfacesCount = r.readU2()
	r.data.Interfaces = r.readInterfaces(r.data.InterfacesCount)

	r.data.FieldsCount = r.readU2()
	r.data.Fields = r.readFields(r.data.FieldsCount)

	r.data.MethodsCount = r.readU2()
	r.data.Methods = r.readMethods(r.data.MethodsCount)

	r.data.AttributesCount = r.readU2()
	r.data.Attributes = r.readAttributes(r.data.AttributesCount)

	return nil
}

func (r *Reader) readConstantPool(count uint16) []ConstantData {
	pool := make([]ConstantData, count)
	pool[0] = nil
	for i := uint16(1); i < count; i++ {
		tag := r.readU1()
		switch tag {
		case TAG_CONSTANT_UTF8:
			length := r.readU2()
			str := r.readUTF8(int(length))
			pool[i] = ConstantUTF8Data{Length: length, UTF8Value: str}
		case TAG_CONSTANT_INTEGER:
			integer := r.readInt32()
			pool[i] = ConstantIntegerData{IntegerValue: integer}
		case TAG_CONSTANT_FLOAT:
			float := r.readFloat32()
			pool[i] = ConstantFloatData{FloatValue: float}
		case TAG_CONSTANT_LONG:
			long := r.readInt64()
			pool[i] = ConstantLongData{LongValue: long}
		case TAG_CONSTANT_DOUBLE:
			double := r.readFloat64()
			pool[i] = ConstantDoubleData{DoubleValue: double}
		case TAG_CONSTANT_CLASS:
			index := r.readU2()
			pool[i] = ConstantClassData{NameIndex: index}
		case TAG_CONSTANT_STRING:
			index := r.readU2()
			pool[i] = ConstantStringData{ValueIndex: index}
		case TAG_CONSTANT_FIELDREF:
			classIndex := r.readU2()
			fieldIndex := r.readU2()
			pool[i] = ConstantFieldRefData{ClassIndex: classIndex, NameAndTypeIndex: fieldIndex}
		case TAG_CONSTANT_METHODREF:
			classIndex := r.readU2()
			methodIndex := r.readU2()
			pool[i] = ConstantMethodRefData{ClassIndex: classIndex, NameAndTypeIndex: methodIndex}
		case TAG_CONSTANT_INTERFACE_METHODREF:
			interfaceIndex := r.readU2()
			methodIndex := r.readU2()
			pool[i] = ConstantInterfaceMethodRefData{ClassIndex: interfaceIndex, NameAndTypeIndex: methodIndex}
		case TAG_CONSTANT_NAME_AND_TYPE:
			nameIndex := r.readU2()
			typeIndex := r.readU2()
			pool[i] = ConstantNameAndTypeData{NameIndex: nameIndex, DescriptorIndex: typeIndex}
		case TAG_CONSTANT_METHOD_HANDLE:
			pool[i] = ConstantMethodHandleData{ReferenceKind: r.readU1(), ReferenceIndex: r.readU2()}
		case TAG_CONSTANT_METHOD_TYPE:
			pool[i] = ConstantMethodTypeData{DescriptorIndex: r.readU2()}
		case TAG_CONSTANT_DYNAMIC:
			pool[i] = ConstantDynamicData{value: r.readU4()}
		case TAG_CONSTANT_INVOKE_DYNAMIC:
			pool[i] = ConstantInvokeDynamicData{BootstrapMethodIndex: r.readU2(), NameAndTypeIndex: r.readU2()}
		case TAG_CONSTANT_MODULE:
			pool[i] = ConstantModuleData{NameIndex: r.readU2()}
		case TAG_CONSTANT_PACKAGE:
			pool[i] = ConstantPackageData{NameIndex: r.readU2()}
		}
	}
	return pool
}

func (r *Reader) readU1() uint8 {
	v, _ := r.reader.ReadUint8()
	return v
}
func (r *Reader) readU2() uint16 {
	v, _ := r.reader.ReadUint16()
	return v
}
func (r *Reader) readU4() uint32 {
	v, _ := r.reader.ReadUint32()
	return v
}
func (r *Reader) readInt32() int32 {
	v, _ := r.reader.ReadInt32()
	return v
}
func (r *Reader) readInt64() int64 {
	v, _ := r.reader.ReadInt64()
	return v
}
func (r *Reader) readFloat32() float32 {
	v, _ := r.reader.ReadFloat32()
	return v
}
func (r *Reader) readFloat64() float64 {
	v, _ := r.reader.ReadFloat64()
	return v
}

func (r *Reader) readUTF8(length int) string {
	str, _ := r.reader.ReadString(length)
	return str
}

func (r *Reader) readBytes(length int) []byte {
	bytes, _ := r.reader.ReadBytes(length)
	return bytes
}

func (r *Reader) readInterfaces(count uint16) []InterfaceData {
	interfaces := make([]InterfaceData, count)
	for i := uint16(0); i < count; i++ {
		index := r.readU2()
		interfaces[i] = InterfaceData{Index: index}
	}
	return interfaces
}

func (r *Reader) readFields(count uint16) []FieldData {
	fields := make([]FieldData, count)
	for i := uint16(0); i < count; i++ {
		f := FieldData{}
		f.AccessFlags = r.readU2()
		f.NameIndex = r.readU2()
		f.DescriptorIndex = r.readU2()
		f.AttributesCount = r.readU2()
		f.Attributes = r.readAttributes(f.AttributesCount)
		fields[i] = f
	}
	return fields
}

func (r *Reader) readMethods(count uint16) []MethodData {
	methods := make([]MethodData, count)
	for i := uint16(0); i < count; i++ {
		m := MethodData{}
		m.AccessFlags = r.readU2()
		m.NameIndex = r.readU2()
		m.DescriptorIndex = r.readU2()
		m.AttributesCount = r.readU2()
		m.Attributes = r.readAttributes(m.AttributesCount)
		methods[i] = m
	}
	return methods
}

func (r *Reader) readAttributes(count uint16) []AttributeData {
	attributes := make([]AttributeData, count)
	for i := uint16(0); i < count; i++ {
		a := AttributeData{}
		a.NameIndex = r.readU2()
		a.Length = r.readU4()
		a.Value = r.readBytes(int(a.Length))
		attributes[i] = a
	}
	return attributes
}
