package class

import (
	"github.com/tk103331/clazz/class/data"
	"strings"
)

type Class struct {
	Version                     uint32
	AccessFlags                 uint16
	Signature                   string
	ThisClass                   string
	SuperClass                  string
	Deprecated                  bool
	Interfaces                  []string
	Fields                      []Field
	Methods                     []Method
	Attributes                  []Attribute
	SourceFile                  string
	SourceDebugExtension        string
	Module                      Module
	InnerClasses                []InnerClass
	OuterClass                  OuterClass
	NestHost                    string
	RuntimeVisibleAnnotations   []Annotation
	RuntimeInvisibleAnnotations []Annotation
	NestMembers                 []string
	BootstrapMethods            []BootstrapMethod
}

type Field struct {
	Name                        string
	AccessFlags                 uint16
	Descriptor                  string
	Signature                   string
	Deprecated                  bool
	RuntimeVisibleAnnotations   []Annotation
	RuntimeInvisibleAnnotations []Annotation
	Attributes                  []Attribute
	Exceptions                  []string
	ConstantValue               interface{}
}

type Method struct {
	Name                                 string
	AccessFlags                          uint16
	Descriptor                           string
	Signature                            string
	Deprecated                           bool
	RuntimeVisibleAnnotations            []Annotation
	RuntimeInvisibleAnnotations          []Annotation
	RuntimeVisibleParameterAnnotations   []ParameterAnnotation
	RuntimeInvisibleParameterAnnotations []ParameterAnnotation
	Attributes                           []Attribute
	Exceptions                           []string
	Parameters                           []MethodParameter
}

type InnerClass struct {
	Name        string
	OuterName   string
	InnerName   string
	AccessFlags uint16
}
type OuterClass struct {
	ClassName  string
	MethodName string
	Descriptor string
}

type Module struct {
	Name        string
	AccessFlags uint16
	Version     string
	MainClass   string
	Packages    []string
	Requires    []ModuleRequire
	Exports     []ModuleExport
	Opens       []ModuleOpen
	Uses        []string
	Provides    []ModuleProvide
}

type ModuleRequire struct {
	Name        string
	AccessFlags uint16
	Version     string
}
type ModuleExport struct {
	Name        string
	AccessFlags uint16
	Modules     []string
}
type ModuleOpen struct {
	Name        string
	AccessFlags uint16
	Modules     []string
}
type ModuleProvide struct {
	Service  string
	Provides []string
}

type Attribute struct {
	Name    string
	Content []byte
}

type TypePath struct {
}

type Handle struct {
	Tag         uint8
	Owner       string
	Name        string
	Descriptor  string
	IsInterface bool
}
type Annotation struct {
	Descriptor   string
	Visible      bool
	ElementPairs []ElementPair
}

type ParameterAnnotation struct {
	Annotations []Annotation
}

type MethodParameter struct {
	ParameterName string
	AccessFlags   uint16
}

type ConstantReference struct {
	Tag         uint8
	Owner       string
	Name        string
	Descriptor  string
	IsInterface bool
}

type ConstantDynamic struct {
	Name                     string
	Descriptor               string
	BootstrapMethod          Handle
	BootstrapMethodArguments []interface{}
}

type BootstrapMethod struct {
	Handle    Handle
	Arguments []interface{}
}

func NewObjectType(internalName string) Type {
	sort := data.TYPE_SORT_INTERNAL
	if strings.HasPrefix(internalName, "[") {
		sort = data.TYPE_SORT_ARRAY
	}
	return Type{sort: sort, value: internalName, begin: 0, end: len(internalName)}
}

func NewMethodType(methodDescriptor string) Type {
	return Type{sort: data.TYPE_SORT_METHOD, value: methodDescriptor, begin: 0, end: len(methodDescriptor)}
}

type Type struct {
	sort  int
	value string
	begin int
	end   int
}
