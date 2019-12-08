package class

import (
	"github.com/tk103331/clazz/class/data"
	"io"
)

type Reader struct {
	reader *data.Reader
	class  *Class
}

func NewReader(reader io.Reader) *Reader {
	return &Reader{reader: data.NewReader(reader)}
}

func (r *Reader) Read() {
	r.reader.Read()
}

func (r *Reader) Accept(visitor Visitor) {
	if visitor == nil {
		return
	}
	r.reader.Accept(&ResolveDataVisitor{visitor: visitor})
}
