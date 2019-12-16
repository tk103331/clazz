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
		visitor.VisitSource(class.SourceFile, class.SourceDebugExtension)

		module := class.Module
		moduleVisitor := visitor.VisitModule(module.Name, module.AccessFlags, module.Version)
		r.acceptModule(moduleVisitor, module)

		visitor.VisitNestHost(class.NestHost)
		visitor.VisitOuterClass(class.OuterClass.ClassName, class.OuterClass.MethodName, class.OuterClass.Descriptor)

		for _, annotation := range class.RuntimeVisibleAnnotations {
			annotationVisitor := visitor.VisitAnnotation(annotation.Descriptor, annotation.Visible)
			r.acceptAnnotation(annotationVisitor, annotation)
		}
		for _, annotation := range class.RuntimeInvisibleAnnotations {
			annotationVisitor := visitor.VisitAnnotation(annotation.Descriptor, annotation.Visible)
			r.acceptAnnotation(annotationVisitor, annotation)
		}
		// TODO RuntimeVisibleTypeAnnotations
		// TODO RuntimeInvisibleTypeAnnotations

		for _, attr := range class.Attributes {
			visitor.VisitAttribute(attr)
		}
		for _, member := range class.NestMembers {
			visitor.VisitNestMember(member)
		}
		for _, cls := range class.InnerClasses {
			visitor.VisitInnerClass(cls.Name, cls.OuterName, cls.InnerName, cls.AccessFlags)
		}
		for _, field := range class.Fields {
			fieldVisitor := visitor.VisitField(field.AccessFlags, field.Name, field.Descriptor, field.Signature, field.ConstantValue)
			r.acceptField(fieldVisitor, field)
		}
		for _, method := range class.Methods {
			methodVisitor := visitor.VisitMethod(method.AccessFlags, method.Name, method.Descriptor, method.Signature, method.Exceptions)
			r.acceptMethod(methodVisitor, method)
		}

		visitor.VisitEnd()
	}
}

func (r *ResolveDataVisitor) acceptModule(visitor ModuleVisitor, module Module) {
	if visitor != nil {
		visitor.VisitMainClass(module.MainClass)
		for _, pkg := range module.Packages {
			visitor.VisitPackage(pkg)
		}
		for _, r := range module.Requires {
			visitor.VisitRequire(r.Name, r.AccessFlags, r.Version)
		}
		for _, e := range module.Exports {
			visitor.VisitExport(e.Name, e.AccessFlags, e.Modules)
		}
		for _, o := range module.Opens {
			visitor.VisitOpen(o.Name, o.AccessFlags, o.Modules)
		}
		for _, u := range module.Uses {
			visitor.VisitUse(u)
		}
		for _, p := range module.Provides {
			visitor.VisitProvide(p.Service, p.Provides)
		}
		visitor.VisitEnd()
	}
}
func (r *ResolveDataVisitor) acceptAnnotation(visitor AnnotationVisitor, annotation Annotation) {
	for _, pair := range annotation.ElementPairs {
		r.acceptAnnotationValue(visitor, pair.Name, pair.Value)
	}
	visitor.VisitEnd()
}

func (r *ResolveDataVisitor) acceptAnnotationValue(visitor AnnotationVisitor, name string, value ElementValue) {
	switch value.Tag() {
	case data.ELEMENT_TAG_BOOLEAN:
		visitor.Visit(name, value.(ElementBooleanValue).Value)
	case data.ELEMENT_TAG_BYTE:
		visitor.Visit(name, value.(ElementByteValue).Value)
	case data.ELEMENT_TAG_CHAR:
		visitor.Visit(name, value.(ElementCharValue).Value)
	case data.ELEMENT_TAG_SHORT:
		visitor.Visit(name, value.(ElementShortValue).Value)
	case data.ELEMENT_TAG_INTEGER:
		visitor.Visit(name, value.(ElementIntegerValue).Value)
	case data.ELEMENT_TAG_LONG:
		visitor.Visit(name, value.(ElementLongValue).Value)
	case data.ELEMENT_TAG_FLOAT:
		visitor.Visit(name, value.(ElementFloatValue).Value)
	case data.ELEMENT_TAG_DOUBLE:
		visitor.Visit(name, value.(ElementDoubleValue).Value)
	case data.ELEMENT_TAG_STRING:
		visitor.Visit(name, value.(ElementStringValue).Value)
	case data.ELEMENT_TAG_CLASS:
		visitor.Visit(name, value.(ElementClassValue).Value)
	case data.ELEMENT_TAG_ANNOTATION:
		annotation := value.(ElementAnnotationValue).Value
		annotationVisitor := visitor.VisitAnnotation(name, annotation.Descriptor)
		r.acceptAnnotation(annotationVisitor, annotation)
	case data.ELEMENT_TAG_ENUM:
		enumValue := value.(ElementEnumValue)
		visitor.VisitEnum(name, enumValue.TypeName, enumValue.ConstName)
	case data.ELEMENT_TAG_ARRAY:
		annotationVisitor := visitor.VisitArray(name)
		for _, elemValue := range value.(ElementArrayValue).Values {
			r.acceptAnnotationValue(annotationVisitor, "", elemValue)
		}
	}
}

func (r *ResolveDataVisitor) acceptField(visitor FieldVisitor, field Field) {
	for _, annotation := range field.RuntimeVisibleAnnotations {
		annotationVisitor := visitor.VisitAnnotation(annotation.Descriptor, annotation.Visible)
		r.acceptAnnotation(annotationVisitor, annotation)
	}
	for _, annotation := range field.RuntimeInvisibleAnnotations {
		annotationVisitor := visitor.VisitAnnotation(annotation.Descriptor, annotation.Visible)
		r.acceptAnnotation(annotationVisitor, annotation)
	}
	// TODO RuntimeVisibleTypeAnnotations
	// TODO RuntimeInvisibleTypeAnnotations
	for _, attribute := range field.Attributes {
		visitor.VisitAttribute(attribute)
	}
	visitor.VisitEnd()
}
func (r *ResolveDataVisitor) acceptMethod(visitor MethodVisitor, method Method) {
	for _, parameter := range method.Parameters {
		visitor.VisitParameter(parameter.ParameterName, parameter.AccessFlags)
	}
	annotationDefaultVisitor := visitor.VisitAnnotationDefault()
	r.acceptAnnotationValue(annotationDefaultVisitor, "", method.AnnotationDefault)

	for _, annotation := range method.RuntimeVisibleAnnotations {
		annotationVisitor := visitor.VisitAnnotation(annotation.Descriptor, annotation.Visible)
		r.acceptAnnotation(annotationVisitor, annotation)
	}
	for _, annotation := range method.RuntimeInvisibleAnnotations {
		annotationVisitor := visitor.VisitAnnotation(annotation.Descriptor, annotation.Visible)
		r.acceptAnnotation(annotationVisitor, annotation)
	}
	// TODO RuntimeVisibleTypeAnnotations
	// TODO RuntimeInvisibleTypeAnnotations
	for index, parameter := range method.RuntimeVisibleParameterAnnotations {
		count := len(method.RuntimeVisibleParameterAnnotations)
		visitor.VisitAnnotableParameterCount(count, true)
		for _, annotation := range parameter.Annotations {
			annotationVisitor := visitor.VisitParameterAnnotation(index, annotation.Descriptor, annotation.Visible)
			r.acceptAnnotation(annotationVisitor, annotation)
		}
	}
	for index, parameter := range method.RuntimeInvisibleParameterAnnotations {
		count := len(method.RuntimeVisibleParameterAnnotations)
		visitor.VisitAnnotableParameterCount(count, true)
		for _, annotation := range parameter.Annotations {
			annotationVisitor := visitor.VisitParameterAnnotation(index, annotation.Descriptor, annotation.Visible)
			r.acceptAnnotation(annotationVisitor, annotation)
		}
	}
	for _, attribute := range method.Attributes {
		visitor.VisitAttribute(attribute)
	}
	// TODO VisitCode
	visitor.VisitEnd()
}

func (r *ResolveDataVisitor) VisitEnd() {
	r.class = &Class{}
	r.constantDynamicValues = make(map[uint16]ConstantDynamic)
	r.resolveAll()
}

func (r *ResolveDataVisitor) resolveInteger(index uint16) int32 {
	integerData := r.Data().ConstantPool[index].(data.ConstantIntegerData)
	return integerData.IntegerValue
}
func (r *ResolveDataVisitor) resolveLong(index uint16) int64 {
	longata := r.Data().ConstantPool[index].(data.ConstantLongData)
	return longata.LongValue
}
func (r *ResolveDataVisitor) resolveFloat(index uint16) float32 {
	floatData := r.Data().ConstantPool[index].(data.ConstantFloatData)
	return floatData.FloatValue
}
func (r *ResolveDataVisitor) resolveDouble(index uint16) float64 {
	doubleData := r.Data().ConstantPool[index].(data.ConstantDoubleData)
	return doubleData.DoubleValue
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
	attributes := make([]Attribute, 0)
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
			// TODO
		case data.DEPRECATED:
			class.Deprecated = true
		case data.SYNTHETIC:
			class.AccessFlags |= data.ACC_SYNTHETIC
		case data.SOURCE_DEBUG_EXTENSION:
			class.SourceDebugExtension = r.resolveUTF8(attr.Value.Uint16())
		case data.RUNTIME_INVISIBLE_ANNOTATIONS:
			class.RuntimeVisibleAnnotations = r.resolveRuntimeAnnotations(attr.Value, false)
		case data.RUNTIME_INVISIBLE_TYPE_ANNOTATIONS:
			// TODO
		case data.RECORD:
		case data.MODULE:
			module = r.resolveModuleAttributes(attr.Value)
		case data.MODULE_MAIN_CLASS:
			moduleMainClass = r.resolveClassName(attr.Value.Uint16())
		case data.MODULE_PACKAGES:
			modulePackages = r.resolveModulePackages(attr.Value)
		case data.BOOTSTRAP_METHODS:
			class.BootstrapMethods = r.resolveBootstrapMethods(attr.Value)
		default:
			attributes = append(attributes, Attribute{Name: name, Content: attr.Value})
		}

	}

	if len(module.Name) > 0 {
		module.MainClass = moduleMainClass
		module.Packages = modulePackages
		class.Module = module
	}
	class.Attributes = attributes
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
			// TODO
		case data.RUNTIME_INVISIBLE_ANNOTATIONS:
			field.RuntimeInvisibleAnnotations = r.resolveRuntimeAnnotations(attr.Value, false)
		case data.RUNTIME_INVISIBLE_TYPE_ANNOTATIONS:
			// TODO
		default:
			attributes = append(attributes, Attribute{Name: name, Content: attr.Value})
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
			method.Code = r.resolveMethodCode(attr.Value)
		case data.EXCEPTIONS:
			method.Exceptions = r.resolveMethodExceptions(attr.Value)
		case data.DEPRECATED:
			method.Deprecated = true
		case data.SYNTHETIC:
			method.AccessFlags |= data.ACC_SYNTHETIC
		case data.ANNOTATION_DEFAULT:
			method.AnnotationDefault = r.readElementValue(attr.Value.Reader())
		case data.RUNTIME_VISIBLE_ANNOTATIONS:
			method.RuntimeVisibleAnnotations = r.resolveRuntimeAnnotations(attr.Value, true)
		case data.RUNTIME_VISIBLE_TYPE_ANNOTATIONS:
			// TODO
		case data.RUNTIME_INVISIBLE_ANNOTATIONS:
			method.RuntimeInvisibleAnnotations = r.resolveRuntimeAnnotations(attr.Value, false)
		case data.RUNTIME_INVISIBLE_TYPE_ANNOTATIONS:
			// TODO
		case data.RUNTIME_VISIBLE_PARAMETER_ANNOTATIONS:
			method.RuntimeVisibleParameterAnnotations = r.resolveRuntimeParameterAnnotations(attr.Value, true)
		case data.RUNTIME_INVISIBLE_PARAMETER_ANNOTATIONS:
			method.RuntimeInvisibleParameterAnnotations = r.resolveRuntimeParameterAnnotations(attr.Value, true)
		case data.METHOD_PARAMETERS:
			method.Parameters = r.resolveMethodParameter(attr.Value)
		default:
			attributes = append(attributes, Attribute{Name: name, Content: attr.Value})
		}
	}
	method.Attributes = attributes
	return method
}

func (r *ResolveDataVisitor) resolveInnerClasses(attrValue data.AttributeValue) []InnerClass {
	reader := attrValue.Reader()
	count := reader.ReadUint16()
	innerClasses := make([]InnerClass, count)
	for i := uint16(0); i < count; i++ {
		currentClassIndex := reader.ReadUint16()
		outerClassIndex := reader.ReadUint16()
		innerClassIndex := reader.ReadUint16()
		accessFlags := reader.ReadUint16()
		innerClasses[i] = InnerClass{r.resolveClassName(currentClassIndex), r.resolveClassName(outerClassIndex), r.resolveUTF8(innerClassIndex), accessFlags}

	}
	return innerClasses
}

func (r *ResolveDataVisitor) resolveOuterClass(attrValue data.AttributeValue) OuterClass {
	reader := attrValue.Reader()
	className := r.resolveClassName(reader.ReadUint16())
	methodName, descriptor := r.resolveNameAndType(reader.ReadUint16())
	return OuterClass{ClassName: className, MethodName: methodName, Descriptor: descriptor}
}

func (r *ResolveDataVisitor) resolveModulePackages(attrValue data.AttributeValue) []string {
	reader := attrValue.Reader()
	packageCount := reader.ReadUint16()
	packages := make([]string, packageCount)
	for i := uint16(0); i < packageCount; i++ {
		packages[i] = r.resolveUTF8(reader.ReadUint16())
	}
	return packages
}

func (r *ResolveDataVisitor) resolveModuleAttributes(attrValue data.AttributeValue) Module {
	reader := attrValue.Reader()
	moduleName := r.resolveClassName(reader.ReadUint16())
	accessFlags := reader.ReadUint16()
	version := r.resolveUTF8(reader.ReadUint16())

	requireCount := reader.ReadUint16()
	requires := make([]ModuleRequire, requireCount)
	for i := uint16(0); i < requireCount; i++ {
		name := r.resolveUTF8(reader.ReadUint16())
		access := reader.ReadUint16()
		version := r.resolveUTF8(reader.ReadUint16())
		requires[i] = ModuleRequire{Name: name, AccessFlags: access, Version: version}
	}

	exportCount := reader.ReadUint16()
	exports := make([]ModuleExport, exportCount)
	for i := uint16(0); i < exportCount; i++ {
		pkgName := r.resolveUTF8(reader.ReadUint16())
		access := reader.ReadUint16()
		exportToCount := reader.ReadUint16()
		var exportTos []string
		if exportToCount != 0 {
			exportTos = make([]string, exportToCount)
			for j := uint16(0); j < exportToCount; j++ {
				exportTos[j] = r.resolveUTF8(reader.ReadUint16())
			}
		}
		exports[i] = ModuleExport{Name: pkgName, AccessFlags: access, Modules: exportTos}
	}

	openCount := reader.ReadUint16()
	opens := make([]ModuleOpen, openCount)
	for i := uint16(0); i < openCount; i++ {
		pkgName := r.resolveUTF8(reader.ReadUint16())
		access := reader.ReadUint16()
		openToCount := reader.ReadUint16()
		var openTos []string
		if openToCount != 0 {
			openTos = make([]string, openToCount)
			for j := uint16(0); j < openToCount; j++ {
				openTos[j] = r.resolveUTF8(reader.ReadUint16())
			}
		}
		opens[i] = ModuleOpen{Name: pkgName, AccessFlags: access, Modules: openTos}
	}

	useCount := reader.ReadUint16()
	uses := make([]string, useCount)
	for i := uint16(0); i < useCount; i++ {
		uses[i] = r.resolveClassName(reader.ReadUint16())
	}

	provideCount := reader.ReadUint16()
	provides := make([]ModuleProvide, provideCount)
	for i := uint16(0); i < provideCount; i++ {
		service := r.resolveUTF8(reader.ReadUint16())
		provideWithCount := reader.ReadUint16()
		provideWiths := make([]string, provideWithCount)
		for j := uint16(0); j < provideWithCount; j++ {
			provideWiths[j] = r.resolveClassName(reader.ReadUint16())
		}
		provides[i] = ModuleProvide{Service: service, Provides: provideWiths}
	}

	return Module{Name: moduleName, AccessFlags: accessFlags, Version: version, Requires: requires, Exports: exports, Opens: opens, Uses: uses, Provides: provides}
}

func (r *ResolveDataVisitor) resolveRuntimeAnnotations(attrValue data.AttributeValue, visible bool) []Annotation {
	reader := attrValue.Reader()
	annotationCount := reader.ReadUint16()
	annotations := make([]Annotation, annotationCount)
	for i := uint16(0); i < annotationCount; i++ {
		annotation := r.readAnnotation(reader)
		annotation.Visible = visible
		annotations[i] = annotation
	}
	return annotations
}
func (r *ResolveDataVisitor) resolveRuntimeParameterAnnotations(attrValue data.AttributeValue, visible bool) []ParameterAnnotation {
	reader := attrValue.Reader()
	parameterCount := reader.ReadUint8()
	parameterAnnotations := make([]ParameterAnnotation, parameterCount)

	for n := uint8(0); n < parameterCount; n++ {
		annotationCount := reader.ReadUint16()
		annotations := make([]Annotation, annotationCount)
		for i := uint16(0); i < annotationCount; i++ {
			annotation := r.readAnnotation(reader)
			annotation.Visible = visible
			annotations[i] = annotation
		}
		parameterAnnotations[n] = ParameterAnnotation{Annotations: annotations}
	}

	return parameterAnnotations
}

func (r *ResolveDataVisitor) readAnnotation(reader *data.AttributeValueReader) Annotation {

	descriptor := r.resolveUTF8(reader.ReadUint16())
	elementPairCount := reader.ReadUint16()
	elementPairs := make([]ElementPair, elementPairCount)
	for j := uint16(0); j < elementPairCount; j++ {
		name := r.resolveUTF8(reader.ReadUint16())
		value := r.readElementValue(reader)
		elementPairs[j] = ElementPair{Name: name, Value: value}
	}
	return Annotation{Descriptor: descriptor, ElementPairs: elementPairs}
}

func (r *ResolveDataVisitor) readElementValue(reader *data.AttributeValueReader) ElementValue {
	tag := reader.ReadUint8()
	switch tag {
	case data.ELEMENT_TAG_BOOLEAN:
		return ElementBooleanValue{Value: r.resolveInteger(reader.ReadUint16()) == 0}
	case data.ELEMENT_TAG_BYTE:
		return ElementByteValue{Value: int8(r.resolveInteger(reader.ReadUint16()))}
	case data.ELEMENT_TAG_CHAR:
		return ElementCharValue{Value: uint16(r.resolveInteger(reader.ReadUint16()))}
	case data.ELEMENT_TAG_SHORT:
		return ElementShortValue{Value: int16(r.resolveInteger(reader.ReadUint16()))}
	case data.ELEMENT_TAG_INTEGER:
		return ElementIntegerValue{Value: r.resolveInteger(reader.ReadUint16())}
	case data.ELEMENT_TAG_LONG:
		return ElementLongValue{Value: r.resolveLong(reader.ReadUint16())}
	case data.ELEMENT_TAG_FLOAT:
		return ElementFloatValue{Value: r.resolveFloat(reader.ReadUint16())}
	case data.ELEMENT_TAG_DOUBLE:
		return ElementDoubleValue{Value: r.resolveDouble(reader.ReadUint16())}
	case data.ELEMENT_TAG_STRING:
		return ElementStringValue{Value: r.resolveUTF8(reader.ReadUint16())}
	case data.ELEMENT_TAG_CLASS:
		return ElementClassValue{Value: r.resolveClassName(reader.ReadUint16())}
	case data.ELEMENT_TAG_ANNOTATION:
		return ElementAnnotationValue{Value: r.readAnnotation(reader)}
	case data.ELEMENT_TAG_ENUM:
		return ElementEnumValue{TypeName: r.resolveUTF8(reader.ReadUint16()), ConstName: r.resolveUTF8(reader.ReadUint16())}
	case data.ELEMENT_TAG_ARRAY:
		itemCount := reader.ReadUint16()
		values := make([]ElementValue, itemCount)
		for i := uint16(0); i < itemCount; i++ {
			value := r.readElementValue(reader)
			values[i] = value
		}
		return ElementArrayValue{Values: values}
	default:
		return nil
	}
}

func (r *ResolveDataVisitor) resolveMethodParameter(attrValue data.AttributeValue) []MethodParameter {
	reader := attrValue.Reader()
	parameterCount := reader.ReadUint8()
	parameters := make([]MethodParameter, parameterCount)
	for i := uint8(0); i < parameterCount; i++ {
		nameIndex := reader.ReadUint16()
		access := reader.ReadUint16()
		parameters[i] = MethodParameter{ParameterName: r.resolveUTF8(nameIndex), AccessFlags: access}

	}
	return parameters
}

func (r *ResolveDataVisitor) resolveMethodExceptions(attrValue data.AttributeValue) []string {
	reader := attrValue.Reader()
	exceptionCount := reader.ReadUint16()
	execeptions := make([]string, exceptionCount)
	for i := uint16(0); i < exceptionCount; i++ {
		execeptions[i] = r.resolveClassName(reader.ReadUint16())
	}
	return execeptions
}

func (r *ResolveDataVisitor) resolveMethodCode(attrValue data.AttributeValue) MethodCode {
	reader := attrValue.Reader()

	maxStack := reader.ReadUint16()
	maxLocal := reader.ReadUint16()
	instructionCount := reader.ReadUint32()
	instructions := make([]Instruction, instructionCount)
	for i := uint32(0); i < instructionCount; i++ {
		opCode := reader.ReadUint8()
		instructions[i] = CodeInstruction{opCode: opCode}
	}
	exceptionCount := reader.ReadUint16()
	exceptions := make([]Exception, exceptionCount)
	for i := uint16(0); i < exceptionCount; i++ {
		start := reader.ReadUint32()
		end := reader.ReadUint32()
		handler := reader.ReadUint32()
		className := r.resolveClassName(reader.ReadUint16())
		exceptions[i] = Exception{StartPC: start, EndPC: end, HandlerPC: handler, CatchType: className}
	}
	attributeCount := reader.ReadUint16()
	attributes := make([]Attribute, attributeCount)
	for i := uint16(0); i < attributeCount; i++ {
		name := r.resolveUTF8(reader.ReadUint16())
		length := reader.ReadUint32()
		bytes := reader.ReadBytes(length)
		attributes[i] = Attribute{Name: name, Content: bytes}
	}

	return MethodCode{MaxStack: maxStack, MaxLocal: maxLocal, InstructionCount: instructionCount, Instructions: instructions}
}

func (r *ResolveDataVisitor) resolveNestMembers(attrValue data.AttributeValue) []string {
	reader := attrValue.Reader()
	nestMemberCount := reader.ReadUint16()
	nestMembers := make([]string, nestMemberCount)
	for i := uint16(0); i < nestMemberCount; i++ {
		member := r.resolveUTF8(reader.ReadUint16())
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
	reader := attrValue.Reader()
	methodCount := reader.ReadUint16()
	methods := make([]BootstrapMethod, methodCount)
	for i := uint16(0); i < methodCount; i++ {
		handle := r.resolveConstantValue(reader.ReadUint16()).(Handle)
		argCount := reader.ReadUint16()
		args := make([]interface{}, argCount)
		for j := uint16(0); j < argCount; j++ {
			args[j] = r.resolveConstantValue(reader.ReadUint16())
		}
		methods[i] = BootstrapMethod{Handle: handle, Arguments: args}
	}
	return methods
}
