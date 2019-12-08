package data

import (
	"encoding/binary"
)

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
const TAG_CONSTANT_INTERFACE_METHODREF uint8 = 11
const TAG_CONSTANT_NAME_AND_TYPE uint8 = 12
const TAG_CONSTANT_METHOD_HANDLE uint8 = 15
const TAG_CONSTANT_METHOD_TYPE uint8 = 16
const TAG_CONSTANT_DYNAMIC uint8 = 17
const TAG_CONSTANT_INVOKE_DYNAMIC uint8 = 18
const TAG_CONSTANT_MODULE uint8 = 19
const TAG_CONSTANT_PACKAGE uint8 = 20

type ClassData struct {
	MagicNumber     uint32
	MinorVersion    uint16
	MajorVersion    uint16
	ConstantCount   uint16
	ConstantPool    []ConstantData
	AccessFlags     uint16
	ThisClass       uint16
	SuperClass      uint16
	InterfacesCount uint16
	Interfaces      []InterfaceData
	FieldsCount     uint16
	Fields          []FieldData
	MethodsCount    uint16
	Methods         []MethodData
	AttributesCount uint16
	Attributes      []AttributeData
}

func (d *ClassData) resolveClassName(index uint16) string {
	classData := d.ConstantPool[index].(ConstantClassData)
	return d.resolveUTF8(classData.NameIndex)
}
func (d *ClassData) resolveString(index uint16) string {
	strData := d.ConstantPool[index].(ConstantStringData)
	return d.resolveUTF8(strData.ValueIndex)
}
func (d *ClassData) resolveUTF8(index uint16) string {
	utf8Data := d.ConstantPool[index].(ConstantUTF8Data)
	return utf8Data.UTF8Value
}

type ConstantData interface {
	Tag() uint8
}

type ConstantUTF8Data struct {
	Length    uint16
	UTF8Value string
}

func (c ConstantUTF8Data) Tag() uint8 {
	return TAG_CONSTANT_UTF8
}

type ConstantIntegerData struct {
	IntegerValue int32
}

func (c ConstantIntegerData) Tag() uint8 {
	return TAG_CONSTANT_INTEGER
}

type ConstantFloatData struct {
	FloatValue float32
}

func (c ConstantFloatData) Tag() uint8 {
	return TAG_CONSTANT_FLOAT
}

type ConstantLongData struct {
	LongValue int64
}

func (c ConstantLongData) Tag() uint8 {
	return TAG_CONSTANT_LONG
}

type ConstantDoubleData struct {
	DoubleValue float64
}

func (c ConstantDoubleData) Tag() uint8 {
	return TAG_CONSTANT_DOUBLE
}

type ConstantClassData struct {
	NameIndex uint16
}

func (c ConstantClassData) Tag() uint8 {
	return TAG_CONSTANT_CLASS
}

type ConstantStringData struct {
	ValueIndex uint16
}

func (c ConstantStringData) Tag() uint8 {
	return TAG_CONSTANT_STRING
}

type ConstantFieldRefData struct {
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

func (c ConstantFieldRefData) Tag() uint8 {
	return TAG_CONSTANT_FIELDREF
}

type ConstantMethodRefData struct {
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

func (c ConstantMethodRefData) Tag() uint8 {
	return TAG_CONSTANT_METHODREF
}

type ConstantInterfaceMethodRefData struct {
	InterfaceIndex   uint16
	NameAndTypeIndex uint16
}

func (c ConstantInterfaceMethodRefData) Tag() uint8 {
	return TAG_CONSTANT_INTERFACE_METHODREF
}

type ConstantNameAndTypeData struct {
	NameIndex uint16
	TypeIndex uint16
}

func (c ConstantNameAndTypeData) Tag() uint8 {
	return TAG_CONSTANT_NAME_AND_TYPE
}

type ConstantMethodHandleData struct {
	value1 uint8
	value2 uint16
}

func (c ConstantMethodHandleData) Tag() uint8 {
	return TAG_CONSTANT_METHOD_HANDLE
}

type ConstantMethodTypeData struct {
	value uint16
}

func (c ConstantMethodTypeData) Tag() uint8 {
	return TAG_CONSTANT_METHOD_TYPE
}

type ConstantDynamicData struct {
	value uint32
}

func (c ConstantDynamicData) Tag() uint8 {
	return TAG_CONSTANT_DYNAMIC
}

type ConstantInvokeDynamicData struct {
	value uint32
}

func (c ConstantInvokeDynamicData) Tag() uint8 {
	return TAG_CONSTANT_INVOKE_DYNAMIC
}

type ConstantPackageData struct {
	value uint16
}

func (c ConstantPackageData) Tag() uint8 {
	return TAG_CONSTANT_PACKAGE
}

type ConstantModuleData struct {
	value uint16
}

func (c ConstantModuleData) Tag() uint8 {
	return TAG_CONSTANT_MODULE
}

type InterfaceData struct {
	Index uint16
}

type FieldData struct {
	AccessFlags     uint16
	NameIndex       uint16
	DescriptorIndex uint16
	AttributesCount uint16
	Attributes      []AttributeData
}
type MethodData struct {
	AccessFlags     uint16
	NameIndex       uint16
	DescriptorIndex uint16
	AttributesCount uint16
	Attributes      []AttributeData
}
type AttributeData struct {
	NameIndex uint16
	Length    uint32
	Value     AttributeValue
}

type AttributeValue []byte

func (v AttributeValue) Uint16() uint16 {
	return binary.BigEndian.Uint16(v)
}
func (v AttributeValue) Uint32() uint32 {
	return binary.BigEndian.Uint32(v)
}
func (v AttributeValue) Uint64() uint64 {
	return binary.BigEndian.Uint64(v)
}
func (v AttributeValue) Uint16Array() []uint16 {
	count := len(v) / 2
	ret := make([]uint16, count)
	for i := 0; i < count; i++ {
		ret[i] = binary.BigEndian.Uint16(v[i*2 : i*2+2])
	}
	return ret
}