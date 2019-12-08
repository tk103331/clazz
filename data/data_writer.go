package data

import (
	"bufio"
	"encoding/binary"
	"io"
)

// DataWriter is wrapper of bufio.Writer.
type DataWriter struct {
	w *bufio.Writer
}

func NewWriter(writer io.Writer) *DataWriter {
	return &DataWriter{w: bufio.NewWriter(writer)}
}

func (dw *DataWriter) WriteByte(value byte) error {
	return dw.w.WriteByte(value)
}
func (dw *DataWriter) WriteBytes(value []byte) error {
	_, err := dw.w.Write(value)
	return err
}

func (dw *DataWriter) WriteUint8(value uint8) error {
	return binary.Write(dw.w, binary.BigEndian, value)
}
func (dw *DataWriter) WriteInt16(value int16) error {
	return binary.Write(dw.w, binary.BigEndian, value)
}
func (dw *DataWriter) WriteUint16(value uint16) error {
	return binary.Write(dw.w, binary.BigEndian, value)
}
func (dw *DataWriter) WriteUint32(value uint32) error {
	return binary.Write(dw.w, binary.BigEndian, value)
}
func (dw *DataWriter) WriteInt32(value int32) error {
	return binary.Write(dw.w, binary.BigEndian, value)
}
func (dw *DataWriter) WriteFloat32(value float32) error {
	return binary.Write(dw.w, binary.BigEndian, value)
}
func (dw *DataWriter) WriteFloat64(value float64) error {
	return binary.Write(dw.w, binary.BigEndian, value)
}
func (dw *DataWriter) WriteChar(value uint16) error {
	return dw.WriteUint16(value)
}
func (dw *DataWriter) WriteInt64(value int64) error {
	return binary.Write(dw.w, binary.BigEndian, value)
}
func (dw *DataWriter) WriteLong(value int64) error {
	return dw.WriteInt64(value)
}
func (dw *DataWriter) WriteBool(value bool) error {
	if value {
		return dw.WriteByte(1)
	} else {
		return dw.WriteByte(0)
	}
}
func (dw *DataWriter) Flush() error {
	return dw.w.Flush()
}
