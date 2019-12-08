package class

type Class struct {
	Version     uint32
	AccessFlags uint16
	ThisClass   string
	SuperClass  string
	Interfaces  []string
	Fields      []Field
	Methods     []Method
	Attributes  []Attribute
}

type Field struct {
	Name        string
	AccessFlags uint16
	Descriptor  string
}

type Method struct {
	Name        string
	AccessFlags uint16
	Descriptor  string
}

type InnerClass struct {
	Name        string
	OuterName   string
	InnerName   string
	AccessFlags uint16
}

type Module struct {
	Name        string
	AccessFlags uint16
	Version     string
	MainClass   string
	Packages    []string
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
	Name string
}
