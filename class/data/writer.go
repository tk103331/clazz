package data

import (
	"github.com/tk103331/clazz/data"
	"io"
)

type Writer struct {
	writer *data.DataWriter
	data   *ClassData
}

func NewWriter(writer io.Writer) *Writer {
	return &Writer{writer: data.NewWriter(writer)}
}

func (w Writer) VisitMagicNumber(magic uint32) {
	w.writer.WriteUint32(magic)
}

func (w Writer) VisitVersion(minorVersion, majorVersion uint16) {
	w.writer.WriteUint16(minorVersion)
	w.writer.WriteUint16(majorVersion)
}

func (w Writer) VisitConstants(constants []ConstantData) {
	writer := w.writer
	count := len(constants)
	writer.WriteUint16(uint16(count))
	for i, data := range constants {
		if i == 0 {
			continue
		}
		tag := data.Tag()
		writer.WriteByte(tag)
		switch tag {
		case TAG_CONSTANT_UTF8:
			utf8Data := data.(ConstantUTF8Data)
			writer.WriteUint16(utf8Data.Length)
			writer.WriteBytes([]byte(utf8Data.UTF8Value))
		case TAG_CONSTANT_INTEGER:
			integerData := data.(ConstantIntegerData)
			writer.WriteInt32(integerData.IntegerValue)
		case TAG_CONSTANT_FLOAT:
			floatData := data.(ConstantFloatData)
			writer.WriteFloat32(floatData.FloatValue)
		case TAG_CONSTANT_LONG:
			longData := data.(ConstantLongData)
			writer.WriteInt64(longData.LongValue)
		case TAG_CONSTANT_DOUBLE:
			doubleData := data.(ConstantDoubleData)
			writer.WriteFloat64(doubleData.DoubleValue)
		case TAG_CONSTANT_CLASS:
			classData := data.(ConstantClassData)
			writer.WriteUint16(classData.NameIndex)
		case TAG_CONSTANT_STRING:
			stringData := data.(ConstantStringData)
			writer.WriteUint16(stringData.ValueIndex)
		case TAG_CONSTANT_FIELDREF:
			fieldRefData := data.(ConstantFieldRefData)
			writer.WriteUint16(fieldRefData.ClassIndex)
			writer.WriteUint16(fieldRefData.NameAndTypeIndex)
		case TAG_CONSTANT_METHODREF:
			methodRefData := data.(ConstantMethodRefData)
			writer.WriteUint16(methodRefData.ClassIndex)
			writer.WriteUint16(methodRefData.NameAndTypeIndex)
		case TAG_CONSTANT_INTERFACE_METHODREF:
			interfaceMethodRefData := data.(ConstantInterfaceMethodRefData)
			writer.WriteUint16(interfaceMethodRefData.ClassIndex)
			writer.WriteUint16(interfaceMethodRefData.NameAndTypeIndex)
		case TAG_CONSTANT_NAME_AND_TYPE:
			nameAndTypeData := data.(ConstantNameAndTypeData)
			writer.WriteUint16(nameAndTypeData.NameIndex)
			writer.WriteUint16(nameAndTypeData.DescriptorIndex)
		case TAG_CONSTANT_METHOD_HANDLE:
			methodHandleData := data.(ConstantMethodHandleData)
			writer.WriteUint8(methodHandleData.ReferenceKind)
			writer.WriteUint16(methodHandleData.ReferenceIndex)
		case TAG_CONSTANT_METHOD_TYPE:
			methodTypeData := data.(ConstantMethodTypeData)
			writer.WriteUint16(methodTypeData.DescriptorIndex)
		case TAG_CONSTANT_DYNAMIC:
			dynamicData := data.(ConstantDynamicData)
			writer.WriteUint32(dynamicData.value)
		case TAG_CONSTANT_INVOKE_DYNAMIC:
			invokeDynamicData := data.(ConstantInvokeDynamicData)
			writer.WriteUint16(invokeDynamicData.BootstrapMethodIndex)
			writer.WriteUint16(invokeDynamicData.NameAndTypeIndex)
		case TAG_CONSTANT_MODULE:
			moduleData := data.(ConstantModuleData)
			writer.WriteUint16(moduleData.NameIndex)
		case TAG_CONSTANT_PACKAGE:
			packageData := data.(ConstantPackageData)
			writer.WriteUint16(packageData.NameIndex)
		}
	}
}

func (w Writer) Visit(thisClass, superClass, access uint16) {
	w.writer.WriteUint16(access)
	w.writer.WriteUint16(thisClass)
	w.writer.WriteUint16(superClass)
}

func (w Writer) VisitInterfaces(interfaces []InterfaceData) {
	writer := w.writer
	count := uint16(len(interfaces))
	writer.WriteUint16(count)
	for _, inter := range interfaces {
		writer.WriteUint16(inter.Index)
	}
}

func (w Writer) VisitFields(fields []FieldData) {
	writer := w.writer
	count := uint16(len(fields))
	writer.WriteUint16(count)
	for _, f := range fields {
		writer.WriteUint16(f.AccessFlags)
		writer.WriteUint16(f.NameIndex)
		writer.WriteUint16(f.DescriptorIndex)
		w.writeAttributes(f.Attributes)
	}
}

func (w Writer) VisitMethods(methods []MethodData) {
	writer := w.writer
	count := uint16(len(methods))
	writer.WriteUint16(count)
	for _, m := range methods {
		writer.WriteUint16(m.AccessFlags)
		writer.WriteUint16(m.NameIndex)
		writer.WriteUint16(m.DescriptorIndex)
		w.writeAttributes(m.Attributes)
	}
}

func (w Writer) VisitAttributes(attributes []AttributeData) {
	w.writeAttributes(attributes)
}

func (w Writer) writeAttributes(attributes []AttributeData) {
	writer := w.writer
	count := uint16(len(attributes))
	writer.WriteUint16(count)
	for _, attr := range attributes {
		writer.WriteUint16(attr.NameIndex)
		writer.WriteUint32(attr.Length)
		writer.WriteBytes(attr.Value)
	}
}

func (w Writer) VisitEnd() {
	w.writer.Flush()
}
