package data

type Visitor interface {
	VisitStart()
	VisitMagicNumber(magic uint32)
	VisitVersion(minorVersion, majorVersion uint16)
	VisitConstants(constants []ConstantData)
	Visit(thisClass, superClass, access uint16)
	VisitInterfaces(interfaces []InterfaceData)
	VisitFields(fields []FieldData)
	VisitMethods(methods []MethodData)
	VisitAttributes(attributes []AttributeData)
	VisitEnd()
}

type DataVisitor struct {
	data *ClassData
}

func NewDataVisitor(data *ClassData) *DataVisitor {
	return &DataVisitor{data: data}
}

func (d *DataVisitor) Data() ClassData {
	return *d.data
}

func (d *DataVisitor) VisitStart() {
	if d.data == nil {
		d.data = &ClassData{}
	}
}

func (d *DataVisitor) VisitMagicNumber(magic uint32) {
	d.data.MagicNumber = magic
}

func (d *DataVisitor) VisitVersion(minorVersion, majorVersion uint16) {
	d.data.MinorVersion = minorVersion
	d.data.MajorVersion = majorVersion
}

func (d *DataVisitor) VisitConstants(constants []ConstantData) {
	d.data.ConstantCount = uint16(len(constants))
	d.data.ConstantPool = constants
}

func (d *DataVisitor) Visit(thisClass, superClass, access uint16) {
	d.data.ThisClass = thisClass
	d.data.SuperClass = superClass
	d.data.AccessFlags = access
}

func (d *DataVisitor) VisitInterfaces(interfaces []InterfaceData) {
	d.data.InterfacesCount = uint16(len(interfaces))
	d.data.Interfaces = interfaces
}

func (d *DataVisitor) VisitFields(fields []FieldData) {
	d.data.FieldsCount = uint16(len(fields))
	d.data.Fields = fields
}

func (d *DataVisitor) VisitMethods(methods []MethodData) {
	d.data.MethodsCount = uint16(len(methods))
	d.data.Methods = methods
}

func (d *DataVisitor) VisitAttributes(attributes []AttributeData) {
	d.data.AttributesCount = uint16(len(attributes))
	d.data.Attributes = attributes
}

func (d *DataVisitor) VisitEnd() {
	panic("implement me")
}
