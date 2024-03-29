package class

import (
	"github.com/tk103331/clazz/common"
	"github.com/tk103331/clazz/tools"
	"os"
	"testing"
)

func TestReadMagic(t *testing.T) {
	f, _ := os.Open("Hello.class")

	r := common.NewReader(f)
	magic, _ := r.ReadUint32()
	tools.AssertEqual(t, 0xcafebabe, int(magic))
}

func TestReadClass(t *testing.T) {
	f, _ := os.Open("Hello.class")

	reader := NewReader(f)
	reader.Read()

	reader.Accept(PrintVisitor{})
}
