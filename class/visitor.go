package class

import "fmt"

type Visitor interface {
	Visit(version uint32, access uint16, name string, signature string, superName string, interfaces []string)
	VisitSource(source string, debug string)
	VisitModule(name string, access uint16, version string)
	VisitNestHost(nestHost string)
	VisitOuterClass(owner string, name string, descriptor string)
	VisitInnerClass(name string, outerName string, innerName string, access uint16)
	VisitAnnotation(descriptor string, visible bool)
	VisitAttribute(attribute Attribute)
	VisitField(field Field)
	VisitMethod(method Method)
	VisitEnd()
}

type PrintVisitor struct {
}

func (p PrintVisitor) Visit(version uint32, access uint16, name string, signature string, superName string, interfaces []string) {
	fmt.Printf("class %s \n", name)
	fmt.Printf("\tminor verison: %d \n", version>>16)
	fmt.Printf("\tmajor verison: %d \n", version&0x00ff)
	fmt.Printf("\tflags: %b \n", access)
	fmt.Printf("\tsuper: %s \n", superName)
	fmt.Printf("\tinterfaces: %s \n", interfaces)

}

func (p PrintVisitor) VisitSource(source string, debug string) {
	fmt.Printf("Source: \"%s\"\n", source)
}

func (p PrintVisitor) VisitModule(name string, access uint16, version string) {
	fmt.Printf("Module Name:%s", name)
}
func (p PrintVisitor) VisitNestHost(name string) {
	fmt.Printf("NestHost Name:%s", name)
}

func (p PrintVisitor) VisitOuterClass(owner string, name string, descriptor string) {
	fmt.Printf("Name: %s  Descriptor: %s  Owner: %s ", name, descriptor, owner)
}

func (p PrintVisitor) VisitInnerClass(name string, outerName string, innerName string, access uint16) {
	fmt.Printf("Name: %s OuterName: %s InnerName: %s Access:%d\n", name, outerName, innerName, access)
}

func (p PrintVisitor) VisitAnnotation(descriptor string, visible bool) {
	fmt.Printf("Descriptor: %s Visible: %v", descriptor, visible)
}

func (p PrintVisitor) VisitAttribute(attribute Attribute) {
	fmt.Printf("Attribute: %v\n", attribute)
}

func (p PrintVisitor) VisitField(field Field) {
	fmt.Printf("Field: %v\n", field)
}

func (p PrintVisitor) VisitMethod(method Method) {
	fmt.Printf("Method: %v\n", method)
}

func (p PrintVisitor) VisitEnd() {
	fmt.Println("end")
}
