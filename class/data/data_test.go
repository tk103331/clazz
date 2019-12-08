package data

import (
	"os"
	"testing"
)

func TestReadWrite(t *testing.T) {
	open, _ := os.Open("../Hello.class")
	reader := NewReader(open)
	reader.Read()

	create, _ := os.Create("Hello.class")
	writer := NewWriter(create)
	reader.Accept(writer)
}
