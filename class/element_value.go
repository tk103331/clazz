package class

import "github.com/tk103331/clazz/class/data"

type ElementPair struct {
	Name  string
	Value ElementValue
}

type ElementValue interface {
	Tag() uint8
}
type ElementBooleanValue struct {
	Value bool
}

func (e ElementBooleanValue) Tag() uint8 {
	return data.ELEMENT_TAG_BOOLEAN
}

type ElementByteValue struct {
	Value int8
}

func (e ElementByteValue) Tag() uint8 {
	return data.ELEMENT_TAG_BYTE
}

type ElementCharValue struct {
	Value uint16
}

func (e ElementCharValue) Tag() uint8 {
	return data.ELEMENT_TAG_CHAR
}

type ElementShortValue struct {
	Value int16
}

func (e ElementShortValue) Tag() uint8 {
	return data.ELEMENT_TAG_SHORT
}

type ElementIntegerValue struct {
	Value int32
}

func (e ElementIntegerValue) Tag() uint8 {
	return data.ELEMENT_TAG_INTEGER
}

type ElementLongValue struct {
	Value int64
}

func (e ElementLongValue) Tag() uint8 {
	return data.ELEMENT_TAG_LONG
}

type ElementFloatValue struct {
	Value float32
}

func (e ElementFloatValue) Tag() uint8 {
	return data.ELEMENT_TAG_FLOAT
}

type ElementDoubleValue struct {
	Value float64
}

func (e ElementDoubleValue) Tag() uint8 {
	return data.ELEMENT_TAG_DOUBLE
}

type ElementStringValue struct {
	Value string
}

func (e ElementStringValue) Tag() uint8 {
	return data.ELEMENT_TAG_STRING
}

type ElementEnumValue struct {
	TypeName  string
	ConstName string
}

func (e ElementEnumValue) Tag() uint8 {
	return data.ELEMENT_TAG_ENUM
}

type ElementClassValue struct {
	Value interface{}
}

func (e ElementClassValue) Tag() uint8 {
	return data.ELEMENT_TAG_CLASS
}

type ElementAnnotationValue struct {
	Value Annotation
}

func (e ElementAnnotationValue) Tag() uint8 {
	return data.ELEMENT_TAG_ANNOTATION
}

type ElementArrayValue struct {
	Length uint16
	Values []ElementValue
}

func (e ElementArrayValue) Tag() uint8 {
	return data.ELEMENT_TAG_ARRAY
}
