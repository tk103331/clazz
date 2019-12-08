package class

import (
	"github.com/tk103331/clazz/class/data"
)

type ResolveDataVisitor struct {
	data.DataVisitor
	class   *Class
	visitor Visitor
}

func (d *ResolveDataVisitor) Class() Class {
	return *d.class
}

func (d ResolveDataVisitor) VisitEnd() {
	d.resolveAll()
}

func (d *ResolveDataVisitor) resolveClassName(index uint16) string {
	classData := d.Data().ConstantPool[index].(data.ConstantClassData)
	return d.resolveUTF8(classData.NameIndex)
}
func (d *ResolveDataVisitor) resolveString(index uint16) string {
	strData := d.Data().ConstantPool[index].(data.ConstantStringData)
	return d.resolveUTF8(strData.ValueIndex)
}
func (d *ResolveDataVisitor) resolveUTF8(index uint16) string {
	utf8Data := d.Data().ConstantPool[index].(data.ConstantUTF8Data)
	return utf8Data.UTF8Value
}

func (r ResolveDataVisitor) resolveAll() {
	d := r.Data()
	visitor := r.visitor
	thisClass := r.resolveClassName(d.ThisClass)
	superClass := r.resolveClassName(d.SuperClass)
	interfaces := make([]string, d.InterfacesCount)

	version := uint32(d.MinorVersion) << 16 & uint32(d.MajorVersion)
	var sourceFile string
	var sourceDebugExtension string
	var innerClasses []InnerClass
	var nestHost string
	var module Module
	for _, attr := range d.Attributes {
		name := r.resolveUTF8(attr.NameIndex)
		switch name {
		case SOURCE_FILE:
			sourceFile = r.resolveUTF8(attr.Value.Uint16())
		case INNER_CLASSES:
			innerClasses = r.resolveInnerClasses(attr.Value)
		case ENCLOSING_METHOD:
		case NEST_HOST:
			nestHost = r.resolveUTF8(attr.Value.Uint16())
		case NEST_MEMBERS:
		case PERMITTED_SUBTYPES:
		case SIGNATURE:
		case RUNTIME_VISIBLE_ANNOTATIONS:
		case RUNTIME_VISIBLE_TYPE_ANNOTATIONS:
		case DEPRECATED:
		case SYNTHETIC:
		case SOURCE_DEBUG_EXTENSION:
			sourceDebugExtension = r.resolveUTF8(attr.Value.Uint16())
		case RUNTIME_INVISIBLE_ANNOTATIONS:
		case RUNTIME_INVISIBLE_TYPE_ANNOTATIONS:
		case RECORD:
		case MODULE:
			module = r.resolveModuleAttributes(attr.Value)
		case MODULE_MAIN_CLASS:
		case MODULE_PACKAGES:
		case BOOTSTRAP_METHODS:

		}

	}
	var signature string
	visitor.Visit(version, d.AccessFlags, thisClass, signature, superClass, interfaces)

	if len(sourceFile) > 0 || len(sourceDebugExtension) > 0 {
		visitor.VisitSource(sourceFile, sourceDebugExtension)
	}

	if len(module.Name) > 0 {
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
		field := r.resolveField(f)
		visitor.VisitField(field)
	}

	for _, m := range d.Methods {
		method := r.resolveMethod(m)
		visitor.VisitMethod(method)
	}

	for _, attr := range d.Attributes {
		name := r.resolveUTF8(attr.NameIndex)
		visitor.VisitAttribute(Attribute{Name: name})
	}
}

func (r *ResolveDataVisitor) resolveField(fieldData data.FieldData) Field {
	field := Field{}

	field.AccessFlags = fieldData.AccessFlags
	field.Name = r.resolveUTF8(fieldData.NameIndex)
	field.Descriptor = r.resolveUTF8(fieldData.DescriptorIndex)
	for _, attr := range fieldData.Attributes {
		name := r.resolveUTF8(attr.NameIndex)
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
func (r *ResolveDataVisitor) resolveMethod(methodData data.MethodData) Method {
	method := Method{}
	method.AccessFlags = methodData.AccessFlags
	method.Name = r.resolveUTF8(methodData.NameIndex)
	method.Descriptor = r.resolveUTF8(methodData.DescriptorIndex)
	for _, attr := range methodData.Attributes {
		name := r.resolveUTF8(attr.NameIndex)
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

func (r *ResolveDataVisitor) resolveInnerClasses(attrValue data.AttributeValue) []InnerClass {
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
		innerClasses[i] = InnerClass{r.resolveClassName(currentClassIndex), r.resolveClassName(outerClassIndex), r.resolveUTF8(innerClassIndex), accessFlags}
		offset += 4

	}
	return innerClasses
}
func (r *ResolveDataVisitor) resolveModuleAttributes(attrValue data.AttributeValue) Module {
	array := attrValue.Uint16Array()
	moduleName := r.resolveClassName(array[0])
	accessFlags := array[1]
	version := r.resolveUTF8(array[2])
	return Module{Name: moduleName, AccessFlags: accessFlags, Version: version}
}
