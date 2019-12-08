package class

import (
	"github.com/tk103331/clazz/data"
	"io"
)

type ClassReader struct {
	r  *data.DataReader
	cv ClassVisitor
	d  *ClassData
}

func NewReader(reader io.Reader) *ClassReader {
	return &ClassReader{r: data.NewReader(reader), d: &ClassData{}}
}

func (cr *ClassReader) Accept(visitor ClassVisitor) {
	if visitor == nil {
		return
	}
	d := cr.d
	thisClass := cr.resolveClassName(d.ThisClass)
	superClass := cr.resolveClassName(d.SuperClass)
	interfaces := make([]string, d.InterfacesCount)

	var sourceFile string
	var sourceDebugExtension string
	var innerClasses []InnerClass
	var nestHost string
	var module Module
	for _, attr := range d.Attributes {
		name := cr.resolveUTF8(attr.NameIndex)
		switch name {
		case SOURCE_FILE:
			sourceFile = cr.resolveUTF8(attr.Value.Uint16())
		case INNER_CLASSES:
			innerClasses = cr.resolveInnerClasses(attr.Value)
		case ENCLOSING_METHOD:
		case NEST_HOST:
			nestHost = cr.resolveUTF8(attr.Value.Uint16())
		case NEST_MEMBERS:
		case PERMITTED_SUBTYPES:
		case SIGNATURE:
		case RUNTIME_VISIBLE_ANNOTATIONS:
		case RUNTIME_VISIBLE_TYPE_ANNOTATIONS:
		case DEPRECATED:
		case SYNTHETIC:
		case SOURCE_DEBUG_EXTENSION:
			sourceDebugExtension = cr.resolveUTF8(attr.Value.Uint16())
		case RUNTIME_INVISIBLE_ANNOTATIONS:
		case RUNTIME_INVISIBLE_TYPE_ANNOTATIONS:
		case RECORD:
		case MODULE:
			module = cr.resolveModuleAttributes(attr.Value)
		case MODULE_MAIN_CLASS:
		case MODULE_PACKAGES:
		case BOOTSTRAP_METHODS:

		}

	}
	var signature string
	visitor.Visit(d.Version, d.AccessFlags, thisClass, signature, superClass, interfaces)

	sourceAttr, sourceOk := d.attrs[SOURCE_FILE]
	debugAttr, debugOk := d.attrs[SOURCE_DEBUG_EXTENSION]
	if sourceOk {
		sourceFile = cr.resolveUTF8(sourceAttr.Value.Uint16())
	}
	if debugOk {
		sourceDebugExtension = cr.resolveUTF8(debugAttr.Value.Uint16())
	}
	if sourceOk || debugOk {
		visitor.VisitSource(sourceFile, sourceDebugExtension)
	}

	if moduleAttr, ok := d.attrs[MODULE]; ok {
		module = cr.resolveModuleAttributes(moduleAttr.Value)
		visitor.VisitModule(module.Name, module.AccessFlags, module.Version)
	}

	if len(nestHost) > 0 {
		visitor.VisitNestHost(nestHost)
	}

	if len(innerClasses) > 0 {
		for _, innerClass := range innerClasses {
			visitor.VisitInnerClass(innerClass.Name, innerClass.OuterName, innerClass.InnerName, innerClass.AccessFlags)
		}
	}

	for _, f := range d.Fields {
		field := cr.resolveField(f)
		visitor.VisitField(field)
	}

	for _, m := range d.Methods {
		method := cr.resolveMethod(m)
		visitor.VisitMethod(method)
	}

	for _, attr := range d.Attributes {
		visitor.VisitAttribute(attr)
	}
}
func (cr *ClassReader) resolveField(fieldInfo FieldInfo) Field {
	field := Field{}

	field.AccessFlags = fieldInfo.AccessFlags
	field.Name = cr.resolveUTF8(fieldInfo.NameIndex)
	field.Descriptor = cr.resolveUTF8(fieldInfo.DescriptorIndex)
	for _, attr := range fieldInfo.Attributes {
		name := cr.resolveUTF8(attr.NameIndex)
		switch name {
		case CONSTANT_VALUE:
		case SIGNATURE:
		case DEPRECATED:
		case SYNTHETIC:
		case RUNTIME_VISIBLE_ANNOTATIONS:
		case RUNTIME_VISIBLE_TYPE_ANNOTATIONS:
		case RUNTIME_INVISIBLE_ANNOTATIONS:
		case RUNTIME_INVISIBLE_TYPE_ANNOTATIONS:
		}
	}

	return field
}
func (cr *ClassReader) resolveMethod(methodInfo MethodInfo) Method {
	method := Method{}
	method.AccessFlags = methodInfo.AccessFlags
	method.Name = cr.resolveUTF8(methodInfo.NameIndex)
	method.Descriptor = cr.resolveUTF8(methodInfo.DescriptorIndex)
	for _, attr := range methodInfo.Attributes {
		name := cr.resolveUTF8(attr.NameIndex)
		switch name {
		case CODE:
		case EXCEPTIONS:
		case DEPRECATED:
		case SYNTHETIC:
		case RUNTIME_VISIBLE_ANNOTATIONS:
		case RUNTIME_VISIBLE_TYPE_ANNOTATIONS:
		case RUNTIME_INVISIBLE_ANNOTATIONS:
		case RUNTIME_INVISIBLE_TYPE_ANNOTATIONS:
		case RUNTIME_VISIBLE_PARAMETER_ANNOTATIONS:
		case RUNTIME_INVISIBLE_PARAMETER_ANNOTATIONS:
		case METHOD_PARAMETERS:
		}
	}
	return method
}
func (cr *ClassReader) resolveClassName(index uint16) string {
	classInfo := cr.d.ConstantPool[index].(ConstantClassInfo)
	return cr.resolveUTF8(classInfo.NameIndex)
}
func (cr *ClassReader) resolveString(index uint16) string {
	strInfo := cr.d.ConstantPool[index].(ConstantStringInfo)
	return cr.resolveUTF8(strInfo.ValueIndex)
}
func (cr *ClassReader) resolveUTF8(index uint16) string {
	utf8Info := cr.d.ConstantPool[index].(ConstantUTF8Info)
	return utf8Info.UTF8Value
}
func (cr *ClassReader) resolveInnerClasses(attrValue AttributeValue) []InnerClass {
	array := attrValue.Uint16Array()
	offset := 0
	count := array[offset]
	innerClasses := make([]InnerClass, count)
	offset++
	for i := uint16(0); i < count; i++ {
		currentClassIndex := array[offset]
		outerClassIndex := array[offset+1]
		innerClassIndex := array[offset+2]
		accessFlags := array[offset+3]
		innerClasses[i] = InnerClass{cr.resolveClassName(currentClassIndex), cr.resolveClassName(outerClassIndex), cr.resolveUTF8(innerClassIndex), accessFlags}
		offset += 4

	}
	return innerClasses
}
func (cr *ClassReader) resolveModuleAttributes(attrValue AttributeValue) Module {
	array := attrValue.Uint16Array()
	moduleName := cr.resolveClassName(array[0])
	accessFlags := array[1]
	version := cr.resolveUTF8(array[2])
	return Module{Name: moduleName, AccessFlags: accessFlags, Version: version}
}

func (cr *ClassReader) Read() error {
	cr.d.MagicNumber = cr.readU4()
	//cr.d.MinorVersion = cr.readU2()
	//cr.d.MajorVersion = cr.readU2()
	cr.d.Version = cr.readU4()

	cr.d.ConstantCount = cr.readU2()
	cr.d.ConstantPool = cr.readConstantPool(cr.d.ConstantCount)

	cr.d.AccessFlags = cr.readU2()
	cr.d.ThisClass = cr.readU2()
	cr.d.SuperClass = cr.readU2()

	cr.d.InterfacesCount = cr.readU2()
	cr.d.Interfaces = cr.readInterfaces(cr.d.InterfacesCount)

	cr.d.FieldsCount = cr.readU2()
	cr.d.Fields = cr.readFields(cr.d.FieldsCount)

	cr.d.MethodsCount = cr.readU2()
	cr.d.Methods = cr.readMethods(cr.d.MethodsCount)

	cr.d.AttributesCount = cr.readU2()
	cr.d.Attributes = cr.readAttributes(cr.d.AttributesCount)
	attrs := make(map[string]AttributeInfo, cr.d.AttributesCount)
	for _, attr := range cr.d.Attributes {
		name := cr.resolveUTF8(attr.NameIndex)
		attrs[name] = attr
	}
	cr.d.attrs = attrs
	return nil
}

func (cr *ClassReader) readConstantPool(count uint16) []ConstantInfo {
	pool := make([]ConstantInfo, count)
	pool[0] = nil
	for i := uint16(1); i < count; i++ {
		tag := cr.readU1()
		switch tag {
		case TAG_CONSTANT_UTF8:
			length := cr.readU2()
			str := cr.readUTF8(int(length))
			pool[i] = ConstantUTF8Info{Length: length, UTF8Value: str}
		case TAG_CONSTANT_INTEGER:
			integer := cr.readInt32()
			pool[i] = ConstantIntegerInfo{IntegerValue: integer}
		case TAG_CONSTANT_FLOAT:
			float := cr.readFloat32()
			pool[i] = ConstantFloatInfo{FloatValue: float}
		case TAG_CONSTANT_LONG:
			long := cr.readInt64()
			pool[i] = ConstantLongInfo{LongValue: long}
		case TAG_CONSTANT_DOUBLE:
			double := cr.readFloat64()
			pool[i] = ConstantDoubleInfo{DoubleValue: double}
		case TAG_CONSTANT_CLASS:
			index := cr.readU2()
			pool[i] = ConstantClassInfo{NameIndex: index}
		case TAG_CONSTANT_STRING:
			index := cr.readU2()
			pool[i] = ConstantStringInfo{ValueIndex: index}
		case TAG_CONSTANT_FIELDREF:
			classIndex := cr.readU2()
			fieldIndex := cr.readU2()
			pool[i] = ConstantFieldRefInfo{ClassIndex: classIndex, NameAndTypeIndex: fieldIndex}
		case TAG_CONSTANT_METHODREF:
			classIndex := cr.readU2()
			methodIndex := cr.readU2()
			pool[i] = ConstantMethodRefInfo{ClassIndex: classIndex, NameAndTypeIndex: methodIndex}
		case TAG_CONSTANT_INTERFACE_METHODREF:
			interfaceIndex := cr.readU2()
			methodIndex := cr.readU2()
			pool[i] = ConstantInterfaceMethodRefInfo{InterfaceIndex: interfaceIndex, NameAndTypeIndex: methodIndex}
		case TAG_CONSTANT_NAME_AND_TYPE:
			nameIndex := cr.readU2()
			typeIndex := cr.readU2()
			pool[i] = ConstantNameAndTypeInfo{NameIndex: nameIndex, TypeIndex: typeIndex}
		case TAG_CONSTANT_METHOD_HANDLE:
			cr.readU1()
			cr.readU2()
			pool[i] = ConstantMethodHandleInfo{}
		case TAG_CONSTANT_METHOD_TYPE:
			cr.readU2()
			pool[i] = ConstantMethodTypeInfo{}
		case TAG_CONSTANT_DYNAMIC:
			cr.readU4()
			pool[i] = ConstantDynamicInfo{}
		case TAG_CONSTANT_INVOKE_DYNAMIC:
			cr.readU4()
			pool[i] = ConstantInvokeDynamicInfo{}
		case TAG_CONSTANT_MODULE:
			cr.readU2()
			pool[i] = ConstantModuleInfo{}
		case TAG_CONSTANT_PACKAGE:
			cr.readU2()
			pool[i] = ConstantModuleInfo{}
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

func (cr *ClassReader) readInterfaces(count uint16) []InterfaceInfo {
	interfaces := make([]InterfaceInfo, count)
	for i := uint16(0); i < count; i++ {
		index := cr.readU2()
		interfaces[i] = InterfaceInfo{Index: index}
	}
	return interfaces
}

func (cr *ClassReader) readFields(count uint16) []FieldInfo {
	fields := make([]FieldInfo, count)
	for i := uint16(0); i < count; i++ {
		f := FieldInfo{}
		f.AccessFlags = cr.readU2()
		f.NameIndex = cr.readU2()
		f.DescriptorIndex = cr.readU2()
		f.AttributesCount = cr.readU2()
		f.Attributes = cr.readAttributes(f.AttributesCount)
		fields[i] = f
	}
	return fields
}

func (cr *ClassReader) readMethods(count uint16) []MethodInfo {
	methods := make([]MethodInfo, count)
	for i := uint16(0); i < count; i++ {
		m := MethodInfo{}
		m.AccessFlags = cr.readU2()
		m.NameIndex = cr.readU2()
		m.DescriptorIndex = cr.readU2()
		m.AttributesCount = cr.readU2()
		m.Attributes = cr.readAttributes(m.AttributesCount)
		methods[i] = m
	}
	return methods
}

func (cr *ClassReader) readAttributes(count uint16) []AttributeInfo {
	attributes := make([]AttributeInfo, count)
	for i := uint16(0); i < count; i++ {
		a := AttributeInfo{}
		a.NameIndex = cr.readU2()
		a.Length = cr.readU4()
		a.Value = cr.readBytes(int(a.Length))
		attributes[i] = a
	}
	return attributes
}
