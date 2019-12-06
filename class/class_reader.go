package class

import (
	"github.com/tk103331/clazz/data"
	"io"
)

type ClassReader struct {
	r  *data.DataReader
	cv ClassVisitor
}

func NewReader(reader io.Reader) *ClassReader {
	return &ClassReader{r: data.NewReader(reader)}
}

func (cr *ClassReader) SetVisitor(visitor ClassVisitor) {
	cr.cv = visitor
}

func (cr *ClassReader) Read() error {

	return nil
}

func (cr *ClassReader) readInfo() error {
	u4Magic, _ := cr.r.ReadInt32()
	u2Minor, _ := cr.r.ReadInt16()
	u2Major, _ := cr.r.ReadInt16()

	magic := int(u4Magic) & 0xffffffff
	minor := int(u2Minor) & 0xffff
	major := int(u2Major) & 0xffff

	return nil
}
