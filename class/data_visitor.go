package class

import (
	"github.com/tk103331/clazz/class/data"
)

type ResolveDataVisitor struct {
	data.DataVisitor
	class                 *Class
	constantDynamicValues map[uint16]ConstantDynamic
	BootstrapMethods      []BootstrapMethod
}

func (r *ResolveDataVisitor) Class() Class {
	return *r.class
}

func (r *ResolveDataVisitor) Accept(visitor Visitor) {
	if visitor != nil {
		class := r.class
		visitor.Visit(class.Version, class.AccessFlags, class.ThisClass, class.SuperClass, class.Signature, class.Interfaces)
		visitor.VisitEnd()
	}
}

func (r *ResolveDataVisitor) VisitEnd() {
	r.class = &Class{}
	r.constantDynamicValues = make(map[uint16]ConstantDynamic)
	r.resolveAll()
}

func (r *ResolveDataVisitor) resolveClassName(index uint16) string {
	classData := r.Data().ConstantPool[index].(data.ConstantClassData)
	return r.resolveUTF8(classData.NameIndex)
}
func (r *ResolveDataVisitor) resolveString(index uint16) string {
	strData := r.Data().ConstantPool[index].(data.ConstantStringData)
	return r.resolveUTF8(strData.ValueIndex)
}
func (r *ResolveDataVisitor) resolveUTF8(index uint16) string {
	utf8Data := r.Data().ConstantPool[index].(data.ConstantUTF8Data)
	return utf8Data.UTF8Value
}
func (r *ResolveDataVisitor) resolveNameAndType(index uint16) (string, string) {
	nameAndTypeData := r.Data().ConstantPool[index].(data.ConstantNameAndTypeData)
	return r.resolveUTF8(nameAndTypeData.NameIndex), r.resolveUTF8(nameAndTypeData.DescriptorIndex)
}
func (r *ResolveDataVisitor) resolveReference(index uint16) ConstantReference {
	referenceData := r.Data().ConstantPool[index].(data.ConstantReferenceData)
	owner := r.resolveUTF8(referenceData.OwnerIndex())
	name, descriptor := r.resolveNameAndType(referenceData.DescriptorIndex())
	isInterface := referenceData.Tag() == data.TAG_CONSTANT_INTERFACE_METHODREF
	return ConstantReference{Tag: referenceData.Tag(), Owner: owner, Name: name, Descriptor: descriptor, IsInterface: isInterface}
}

func (r ResolveDataVisitor) resolveAll() {
	classData := r.Data()
	class := r.class
	class.ThisClass = r.resolveClassName(classData.ThisClass)
	class.SuperClass = r.resolveClassName(classData.SuperClass)
	interfaces := make([]string, classData.InterfacesCount)
	for i := uint16(0); i < classData.InterfacesCount; i++ {
		interfaces[i] = r.resolveClassName(classData.Interfaces[i].Index)
	}
	class.Interfaces = interfaces

	class.Version = uint32(classData.MinorVersion) << 16 & uint32(classData.MajorVersion)

	var module Module
	var moduleMainClass string
	var modulePackages []string
	for _, attr := range classData.Attributes {
		name := r.resolveUTF8(attr.NameIndex)
		switch name {
		case data.SOURCE_FILE:
			class.SourceFile = r.resolveUTF8(attr.Value.Uint16())
		case data.INNER_CLASSES:
			class.InnerClasses = r.resolveInnerClasses(attr.Value)
		case data.ENCLOSING_METHOD:
			class.OuterClass = r.resolveOuterClass(attr.Value)
		case data.NEST_HOST:
			class.NestHost = r.resolveUTF8(attr.Value.Uint16())
		case data.NEST_MEMBERS:
			class.NestMembers = r.resolveNestMembers(attr.Value)
		case data.PERMITTED_SUBTYPES:
		case data.SIGNATURE:
			class.Signature = r.resolveUTF8(attr.Value.Uint16())
		case data.RUNTIME_VISIBLE_ANNOTATIONS:
			class.RuntimeVisibleAnnotations = r.resolveRuntimeAnnotations(attr.Value, true)
		case data.RUNTIME_VISIBLE_TYPE_ANNOTATIONS:
		case data.DEPRECATED:
			class.Deprecated = true
		case data.SYNTHETIC:
			class.AccessFlags |= data.ACC_SYNTHETIC
		case data.SOURCE_DEBUG_EXTENSION:
			class.SourceDebugExtension = r.resolveUTF8(attr.Value.Uint16())
		case data.RUNTIME_INVISIBLE_ANNOTATIONS:
			class.RuntimeVisibleAnnotations = r.resolveRuntimeAnnotations(attr.Value, false)
		case data.RUNTIME_INVISIBLE_TYPE_ANNOTATIONS:
		case data.RECORD:
		case data.MODULE:
			module = r.resolveModuleAttributes(attr.Value)
		case data.MODULE_MAIN_CLASS:
			moduleMainClass = r.resolveClassName(attr.Value.Uint16())
		case data.MODULE_PACKAGES:
			modulePackages = r.resolveModulePackages(attr.Value)
		case data.BOOTSTRAP_METHODS:
			class.BootstrapMethods = r.resolveBootstrapMethods(attr.Value)
		}

	}

	if len(module.Name) > 0 {
		module.MainClass = moduleMainClass
		module.Packages = modulePackages
		class.Module = module
	}

}

func (r *ResolveDataVisitor) resolveConstantValue(constIndex uint16) interface{} {
	pool := r.Data().ConstantPool
	if constIndex < 0 || int(constIndex) > len(pool) {
		return nil
	}
	constantData := pool[constIndex]
	tag := constantData.Tag()
	switch tag {
	case data.TAG_CONSTANT_UTF8:
		utf8Data := constantData.(data.ConstantUTF8Data)
		return utf8Data.UTF8Value
	case data.TAG_CONSTANT_INTEGER:
		integerData := constantData.(data.ConstantIntegerData)
		return integerData.IntegerValue
	case data.TAG_CONSTANT_FLOAT:
		floatData := constantData.(data.ConstantFloatData)
		return floatData.FloatValue
	case data.TAG_CONSTANT_LONG:
		longData := constantData.(data.ConstantLongData)
		return longData.LongValue
	case data.TAG_CONSTANT_DOUBLE:
		doubleData := constantData.(data.ConstantDoubleData)
		return doubleData.DoubleValue
	case data.TAG_CONSTANT_CLASS:
		classData := constantData.(data.ConstantClassData)
		className := r.resolveUTF8(classData.NameIndex)
		return NewObjectType(className)
	case data.TAG_CONSTANT_STRING:
		stringData := constantData.(data.ConstantStringData)
		return r.resolveUTF8(stringData.ValueIndex)
	case data.TAG_CONSTANT_METHOD_HANDLE:
		methodHandleData := constantData.(data.ConstantMethodHandleData)
		reference := r.resolveReference(methodHandleData.ReferenceIndex)
		return Handle{Tag: methodHandleData.Tag(), Owner: reference.Owner, Name: reference.Name, Descriptor: reference.Descriptor}
	case data.TAG_CONSTANT_METHOD_TYPE:
		methodTypeData := constantData.(data.ConstantMethodTypeData)
		descriptor := r.resolveUTF8(methodTypeData.DescriptorIndex)
		return NewMethodType(descriptor)
	case data.TAG_CONSTANT_DYNAMIC:
		dynamicData := constantData.(data.ConstantDynamicData)
		return dynamicData
	default:
		return nil
	}
}

func (r ResolveDataVisitor) resolveConstantDynamicHandle() {

}

func (r *ResolveDataVisitor) resolveField(fieldData data.FieldData) Field {
	field := Field{}

	field.AccessFlags = fieldData.AccessFlags
	field.Name = r.resolveUTF8(fieldData.NameIndex)
	field.Descriptor = r.resolveUTF8(fieldData.DescriptorIndex)
	attributes := make([]Attribute, 0)
	for _, attr := range fieldData.Attributes {
		name := r.resolveUTF8(attr.NameIndex)
		switch name {
		case data.CONSTANT_VALUE:
			constValueIndex := attr.Value.Uint16()
			if constValueIndex > 0 {
				field.ConstantValue = r.resolveConstantValue(constValueIndex)
			}
		case data.SIGNATURE:
			field.Signature = r.resolveUTF8(attr.Value.Uint16())
		case data.DEPRECATED:
			field.Deprecated = true
		case data.SYNTHETIC:
			field.AccessFlags |= data.ACC_SYNTHETIC
		case data.RUNTIME_VISIBLE_ANNOTATIONS:
			field.RuntimeVisibleAnnotations = r.resolveRuntimeAnnotations(attr.Value, true)
		case data.RUNTIME_VISIBLE_TYPE_ANNOTATIONS:
		case data.RUNTIME_INVISIBLE_ANNOTATIONS:
			field.RuntimeInvisibleAnnotations = r.resolveRuntimeAnnotations(attr.Value, false)
		case data.RUNTIME_INVISIBLE_TYPE_ANNOTATIONS:
		default:
			attributes = append(attributes, Attribute{Name: name})
		}
	}
	field.Attributes = attributes
	return field
}
func (r *ResolveDataVisitor) resolveMethod(methodData data.MethodData) Method {
	method := Method{}
	method.AccessFlags = methodData.AccessFlags
	method.Name = r.resolveUTF8(methodData.NameIndex)
	method.Descriptor = r.resolveUTF8(methodData.DescriptorIndex)
	attributes := make([]Attribute, 0)
	for _, attr := range methodData.Attributes {
		name := r.resolveUTF8(attr.NameIndex)
		switch name {
		case data.CODE:
		case data.EXCEPTIONS:
		case data.DEPRECATED:
			method.Deprecated = true
		case data.SYNTHETIC:
			method.AccessFlags |= data.ACC_SYNTHETIC
		case data.RUNTIME_VISIBLE_ANNOTATIONS:
			method.RuntimeVisibleAnnotations = r.resolveRuntimeAnnotations(attr.Value, true)
		case data.RUNTIME_VISIBLE_TYPE_ANNOTATIONS:
		case data.RUNTIME_INVISIBLE_ANNOTATIONS:
			method.RuntimeInvisibleAnnotations = r.resolveRuntimeAnnotations(attr.Value, false)
		case data.RUNTIME_INVISIBLE_TYPE_ANNOTATIONS:
		case data.RUNTIME_VISIBLE_PARAMETER_ANNOTATIONS:
		case data.RUNTIME_INVISIBLE_PARAMETER_ANNOTATIONS:
		case data.METHOD_PARAMETERS:
		default:
			attributes = append(attributes, Attribute{Name: name})
		}
	}
	method.Attributes = attributes
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

func (r *ResolveDataVisitor) resolveOuterClass(attrValue data.AttributeValue) OuterClass {
	array := attrValue.Uint16Array()
	className := r.resolveClassName(array[0])
	methodName, descriptor := r.resolveNameAndType(array[1])
	return OuterClass{ClassName: className, MethodName: methodName, Descriptor: descriptor}
}

func (r *ResolveDataVisitor) resolveModulePackages(attrValue data.AttributeValue) []string {
	offset := 0
	array := attrValue.Uint16Array()
	packageCount := array[offset]
	offset += 1
	packages := make([]string, packageCount)
	offset += 1
	for i := uint16(0); i < packageCount; i++ {
		packages[i] = r.resolveUTF8(array[offset])
		offset += 1
	}
	return packages
}

func (r *ResolveDataVisitor) resolveModuleAttributes(attrValue data.AttributeValue) Module {
	offset := 0
	array := attrValue.Uint16Array()
	moduleName := r.resolveClassName(array[offset])
	accessFlags := array[offset+1]
	version := r.resolveUTF8(array[offset+2])

	requireCount := array[offset]
	offset += 1
	requires := make([]ModuleRequire, requireCount)
	for i := uint16(0); i < requireCount; i++ {
		name := r.resolveUTF8(array[offset])
		access := array[offset+1]
		version := r.resolveUTF8(array[offset+2])
		requires[i] = ModuleRequire{Name: name, AccessFlags: access, Version: version}
		offset += 3
	}

	exportCount := array[offset]
	offset += 1
	exports := make([]ModuleExport, exportCount)
	for i := uint16(0); i < exportCount; i++ {
		pkgName := r.resolveUTF8(array[offset])
		access := array[offset+1]
		exportToCount := array[offset+2]
		offset += 3
		var exportTos []string
		if exportToCount != 0 {
			exportTos = make([]string, exportToCount)
			for j := uint16(0); j < exportToCount; j++ {
				exportTos[j] = r.resolveUTF8(array[offset])
				offset += 1
			}
		}
		exports[i] = ModuleExport{Name: pkgName, AccessFlags: access, Modules: exportTos}
	}

	openCount := array[offset]
	offset += 1
	opens := make([]ModuleOpen, openCount)
	for i := uint16(0); i < openCount; i++ {
		pkgName := r.resolveUTF8(array[offset])
		access := array[offset+1]
		openToCount := array[offset+2]
		offset += 3
		var openTos []string
		if openToCount != 0 {
			openTos = make([]string, openToCount)
			for j := uint16(0); j < openToCount; j++ {
				openTos[j] = r.resolveUTF8(array[offset])
				offset += 1
			}
		}
		opens[i] = ModuleOpen{Name: pkgName, AccessFlags: access, Modules: openTos}
	}

	useCount := array[offset]
	offset += 1
	uses := make([]string, useCount)
	for i := uint16(0); i < useCount; i++ {
		uses[i] = r.resolveClassName(array[offset])
		offset += 1
	}

	provideCount := array[offset]
	offset += 1
	provides := make([]ModuleProvide, provideCount)
	for i := uint16(0); i < provideCount; i++ {
		service := r.resolveUTF8(array[offset])
		provideWithCount := array[offset+1]
		offset += 2
		provideWiths := make([]string, provideWithCount)
		for j := uint16(0); j < provideWithCount; j++ {
			provideWiths[j] = r.resolveClassName(array[offset])
			offset += 1
		}
		provides[i] = ModuleProvide{Service: service, Provides: provideWiths}
	}

	return Module{Name: moduleName, AccessFlags: accessFlags, Version: version, Requires: requires, Exports: exports, Opens: opens, Uses: uses, Provides: provides}
}

func (r *ResolveDataVisitor) resolveRuntimeAnnotations(attrValue data.AttributeValue, visible bool) []Annotation {
	array := attrValue.Uint16Array()
	offset := 0
	annotationCount := array[offset]
	offset += 1
	annotations := make([]Annotation, annotationCount)
	for i := uint16(0); i < annotationCount; i++ {
		descriptor := r.resolveUTF8(array[offset])
		offset += 1
		annotations[i] = Annotation{Descriptor: descriptor, Visible: visible}
	}
	return annotations
}
func (r *ResolveDataVisitor) resolveNestMembers(attrValue data.AttributeValue) []string {
	array := attrValue.Uint16Array()
	offset := 0
	nestMemberCount := array[offset]
	offset += 1
	nestMembers := make([]string, nestMemberCount)
	for i := uint16(0); i < nestMemberCount; i++ {
		member := r.resolveUTF8(array[offset])
		offset += 1
		nestMembers[i] = member
	}
	return nestMembers
}
func (r *ResolveDataVisitor) resolveConstantDynamic(index uint16) ConstantDynamic {
	if constantDynamic, ok := r.constantDynamicValues[index]; ok {
		return constantDynamic
	}
	dynamicData := r.Data().ConstantPool[index].(data.ConstantDynamicData)
	name, descriptor := r.resolveNameAndType(dynamicData.NameAndTypeIndex)
	bootstrapMethod := r.Class().BootstrapMethods[dynamicData.BootstrapMethodIndex]
	return ConstantDynamic{Name: name, Descriptor: descriptor, BootstrapMethod: bootstrapMethod.Handle, BootstrapMethodArguments: bootstrapMethod.Arguments}
}

func (r *ResolveDataVisitor) resolveBootstrapMethods(attrValue data.AttributeValue) []BootstrapMethod {
	array := attrValue.Uint16Array()
	offset := 0
	methodCount := array[offset]
	offset += 1
	methods := make([]BootstrapMethod, methodCount)
	for i := uint16(0); i < methodCount; i++ {
		handle := r.resolveConstantValue(array[offset]).(Handle)
		argCount := array[offset+1]
		offset += 2
		args := make([]interface{}, argCount)
		for j := uint16(0); j < argCount; j++ {
			args[j] = r.resolveConstantValue(array[offset])
			offset += 1
		}
		methods[i] = BootstrapMethod{Handle: handle, Arguments: args}
	}
	return methods
}
