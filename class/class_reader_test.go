package class

import (
	"github.com/tk103331/clazz/data"
	"github.com/tk103331/clazz/tools"
	"os"
	"testing"
)

func TestReadMagic(t *testing.T) {
	f, _ := os.Open("Hello.class")

	r := data.NewReader(f)
	magic, _ := r.ReadInt32()
	tools.AssertEqual(t, 0xcafebabe, int(magic)&0xffffffff)
}
