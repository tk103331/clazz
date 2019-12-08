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
}

type Attribute struct {
	Name string
}
