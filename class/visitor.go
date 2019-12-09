package class

import "fmt"

// A visitor to visit a Java class.
// The methods of this class must be called in the following order:
// visit [ visitSource ] [ visitModule ][ visitNestHost ][ visitPermittedSubtype ][ visitOuterClass ] ( visitAnnotation | visitTypeAnnotation | visitAttribute )* ( visitNestMember | visitInnerClass | visitField | visitMethod )* visitEnd.
type Visitor interface {
	Visit(version uint32, access uint16, name string, signature string, superName string, interfaces []string)
	VisitSource(source string, debug string)
	VisitModule(name string, access uint16, version string) ModuleVisitor
	VisitNestHost(nestHost string)
	VisitOuterClass(owner string, name string, descriptor string)
	VisitAnnotation(descriptor string, visible bool) AnnotationVisitor
	VisitTypeAnnotation(typeRef int, typePath TypePath, descriptor string, visible bool) AnnotationVisitor
	VisitAttribute(attribute Attribute)
	VisitNestMember(nestMember string)
	VisitInnerClass(name string, outerName string, innerName string, access uint16)
	VisitField(access uint16, name string, descriptor string, signature string, value interface{}) FieldVisitor
	VisitMethod(access uint16, name string, descriptor string, signature string, exceptions []string)
	//VisitRecordComponentExperimental(access uint16, name string, descriptor string, signature string)
	VisitEnd()
}

// A visitor to visit a Java module.
// The methods of this class must be called in the following order:
// ( visitMainClass | ( visitPackage | visitRequire | visitExport | visitOpen | visitUse | visitProvide )* ) visitEnd.
type ModuleVisitor interface {
	VisitMainClass(mainClass string)
	VisitPackage(packageName string)
	VisitRequire(moduleName string, access uint16, version string)
	VisitExport(packageName string, access uint16, modules []string)
	VisitOpen(packageName string, access uint16, modules []string)
	VisitUse(service string)
	VisitProvide(service string, providers []string)
	VisitEnd()
}

// A visitor to visit a Java method.
// The methods of this class must be called in the following order:
// ( visitParameter )* [ visitAnnotationDefault ] ( visitAnnotation | visitAnnotableParameterCount | visitParameterAnnotation visitTypeAnnotation | visitAttribute )* [ visitCode ( visitFrame | visit<i>X</i>Insn | visitLabel | visitInsnAnnotation | visitTryCatchBlock | visitTryCatchAnnotation | visitLocalVariable | visitLocalVariableAnnotation | visitLineNumber )* visitMaxs ] visitEnd. In addition, the visit<i>X</i>Insn and visitLabel methods must be called in the sequential order of the bytecode instructions of the visited code, visitInsnAnnotation must be called after the annotated instruction, visitTryCatchBlock must be called before the labels passed as arguments have been visited, visitTryCatchBlockAnnotation must be called after the corresponding try catch block has been visited, and the visitLocalVariable, visitLocalVariableAnnotation and visitLineNumber methods must be called after the labels passed as arguments have been visited.
type MethodVisitor interface {
	VisitParameter(name string, access uint16)
	VisitAnnotationDefault() AnnotationVisitor
	VisitAnnotation(descriptor string, visible bool) AnnotationVisitor
	VisitTypeAnnotation(typeRef int, typePath TypePath, descriptor string, visible bool) AnnotationVisitor
	VisitAnnotableParameterCount(parameterCount int, visible bool)
	VisitParameterAnnotation(parameterIndex int, descriptor string, visible bool) AnnotationVisitor
	VisitAttribute(attribute Attribute)
	VisitCode()
	VisitFrame(frameType int, numLocal int, locals []interface{}, numStack int, stacks []interface{})
	VisitInstruction(opCode uint16)
	VisitIntInstruction(opCode uint16, operand int32)
	VisitVarInstruction(opCode uint16, variable int)
	VisitTypeInstruction(opCode uint16, typeName string)
	VisitFieldInstruction(opCode uint16, owner string, name string, descriptor string)
	VisitMethodInstruction(opCode uint16, owner string, name string, descriptor string, isInterface bool)
	VisitInvokeDynamicInstruction(opCode uint16, name string, descriptor string, bootstrapMethodHandle Handle, bootstrapMethodArguments []interface{})
}

// A visitor to visit a Java field.
// The methods of this class must be called in the following order:
// ( visitAnnotation | visitTypeAnnotation | visitAttribute )* visitEnd.
type FieldVisitor interface {
	VisitAnnotation(descriptor string, visible bool) AnnotationVisitor
	VisitTypeAnnotation(typeRef int, typePath TypePath, descriptor string, visible bool) AnnotationVisitor
	VisitAttribute(attribute Attribute)
	VisitEnd()
}

// A visitor to visit a Java annotation.
// The methods of this class must be called in the following order:
// ( visit | visitEnum | visitAnnotation | visitArray )* visitEnd.
type AnnotationVisitor interface {
	Visit(name string, value interface{})
	VisitEnum(name string, descriptor string, value string)
	VisitAnnotation(descriptor string, visible bool) AnnotationVisitor
	VisitArray(name string) AnnotationVisitor
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
