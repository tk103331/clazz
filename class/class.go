package class

const MAGIC_NUMBER = 0xcafebabe
const TAG_CONSTANT_UTF8 uint8 = 1
const TAG_CONSTANT_INTEGER uint8 = 3
const TAG_CONSTANT_FLOAT uint8 = 4
const TAG_CONSTANT_LONG uint8 = 5
const TAG_CONSTANT_DOUBLE uint8 = 6
const TAG_CONSTANT_CLASS uint8 = 7
const TAG_CONSTANT_STRING uint8 = 8
const TAG_CONSTANT_FIELDREF uint8 = 9
const TAG_CONSTANT_METHODREF uint8 = 10
const TAG_CONSTANT_IFMETHODREF uint8 = 11
const TAG_CONSTANT_NAMEANDTYPE uint8 = 12

type Data struct {
	MagicNumber     uint32
	MinorVersion    uint16
	MajorVersion    uint16
	ConstantCount   uint16
	ConstantPool    []Constant
	AccessFlags     uint16
	ThisClass       uint16
	SuperClass      uint16
	InterfacesCount uint16
	Interfaces      []Interface
	FieldsCount     uint16
	Fields          []Field
	MethodsCount    uint16
	Methods         []Method
	AttributesCount uint16
	Attributes      []Attribute
}

type Constant interface {
	Tag() uint8
}

type ConstantUTF8 struct {
	Length uint16
	Value  string
}

func (c ConstantUTF8) Tag() uint8 {
	return TAG_CONSTANT_UTF8
}

type ConstantInteger struct {
	Value int32
}

func (c ConstantInteger) Tag() uint8 {
	return TAG_CONSTANT_INTEGER
}

type ConstantFloat struct {
	Value float32
}

func (c ConstantFloat) Tag() uint8 {
	return TAG_CONSTANT_FLOAT
}

type ConstantLong struct {
	Value int64
}

func (c ConstantLong) Tag() uint8 {
	return TAG_CONSTANT_LONG
}

type ConstantDouble struct {
	Value float64
}

func (c ConstantDouble) Tag() uint8 {
	return TAG_CONSTANT_DOUBLE
}

type ConstantClass struct {
	Index uint16
}

func (c ConstantClass) Tag() uint8 {
	return TAG_CONSTANT_CLASS
}

type ConstantString struct {
	Index uint16
}

func (c ConstantString) Tag() uint8 {
	return TAG_CONSTANT_STRING
}

type ConstantFieldRef struct {
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

func (c ConstantFieldRef) Tag() uint8 {
	return TAG_CONSTANT_FIELDREF
}

type ConstantMethodRef struct {
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

func (c ConstantMethodRef) Tag() uint8 {
	return TAG_CONSTANT_METHODREF
}

type ConstantIfMethodRef struct {
	InterfaceIndex   uint16
	NameAndTypeIndex uint16
}

func (c ConstantIfMethodRef) Tag() uint8 {
	return TAG_CONSTANT_IFMETHODREF
}

type ConstantNameAndType struct {
	NameIndex uint16
	TypeIndex uint16
}

func (c ConstantNameAndType) Tag() uint8 {
	return TAG_CONSTANT_NAMEANDTYPE
}

type Interface struct {
	Index uint16
}

type Field struct {
	AccessFlags     uint16
	NameIndex       uint16
	DescriptorIndex uint16
	AttributesCount uint16
	Attributes      []Attribute
}
type Method struct {
	AccessFlags     uint16
	NameIndex       uint16
	DescriptorIndex uint16
	AttributesCount uint16
	Attributes      []Attribute
}
type Attribute struct {
	NameIndex uint16
	Length    uint8
	Info      []byte
}
