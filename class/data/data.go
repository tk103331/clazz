package data

import (
	"bytes"
	"encoding/binary"
	"github.com/tk103331/clazz/common"
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

type ConstantReferenceData interface {
	ConstantData
	OwnerIndex() uint16
	DescriptorIndex() uint16
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
func (c ConstantFieldRefData) OwnerIndex() uint16 {
	return c.ClassIndex
}

func (c ConstantFieldRefData) DescriptorIndex() uint16 {
	return c.NameAndTypeIndex
}

type ConstantMethodRefData struct {
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

func (c ConstantMethodRefData) Tag() uint8 {
	return TAG_CONSTANT_METHODREF
}
func (c ConstantMethodRefData) OwnerIndex() uint16 {
	return c.ClassIndex
}

func (c ConstantMethodRefData) DescriptorIndex() uint16 {
	return c.NameAndTypeIndex
}

type ConstantInterfaceMethodRefData struct {
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

func (c ConstantInterfaceMethodRefData) Tag() uint8 {
	return TAG_CONSTANT_INTERFACE_METHODREF
}

func (c ConstantInterfaceMethodRefData) OwnerIndex() uint16 {
	return c.ClassIndex
}

func (c ConstantInterfaceMethodRefData) DescriptorIndex() uint16 {
	return c.NameAndTypeIndex
}

type ConstantNameAndTypeData struct {
	NameIndex       uint16
	DescriptorIndex uint16
}

func (c ConstantNameAndTypeData) Tag() uint8 {
	return TAG_CONSTANT_NAME_AND_TYPE
}

type ConstantMethodHandleData struct {
	ReferenceKind  uint8
	ReferenceIndex uint16
}

func (c ConstantMethodHandleData) Tag() uint8 {
	return TAG_CONSTANT_METHOD_HANDLE
}

type ConstantMethodTypeData struct {
	DescriptorIndex uint16
}

func (c ConstantMethodTypeData) Tag() uint8 {
	return TAG_CONSTANT_METHOD_TYPE
}

type ConstantDynamicData struct {
	BootstrapMethodIndex uint16
	NameAndTypeIndex     uint16
}

func (c ConstantDynamicData) Tag() uint8 {
	return TAG_CONSTANT_DYNAMIC
}

type ConstantInvokeDynamicData struct {
	BootstrapMethodIndex uint16
	NameAndTypeIndex     uint16
}

func (c ConstantInvokeDynamicData) Tag() uint8 {
	return TAG_CONSTANT_INVOKE_DYNAMIC
}

type ConstantPackageData struct {
	NameIndex uint16
}

func (c ConstantPackageData) Tag() uint8 {
	return TAG_CONSTANT_PACKAGE
}

type ConstantModuleData struct {
	NameIndex uint16
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
func (v AttributeValue) Reader() *AttributeValueReader {
	return &AttributeValueReader{reader: common.NewReader(bytes.NewBuffer(v))}
}

type AttributeValueReader struct {
	reader *common.DataReader
}

func (r AttributeValueReader) ReadUint8() uint8 {
	v, _ := r.reader.ReadUint8()
	return v
}

func (r AttributeValueReader) ReadUint16() uint16 {
	v, _ := r.reader.ReadUint16()
	return v
}

func (r AttributeValueReader) ReadUint32() uint32 {
	v, _ := r.reader.ReadUint32()
	return v
}
