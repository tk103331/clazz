package data

import (
	"bufio"
	"encoding/binary"
	"io"
)

// DataReader is wrapper of bufio.Reader.
type DataReader struct {
	r *bufio.Reader
}

func NewReader(reader io.Reader) *DataReader {
	return &DataReader{bufio.NewReader(reader)}
}

// Read reads a byte.
func (dr *DataReader) ReadByte() (byte, error) {
	return dr.r.ReadByte()
}

func (dr *DataReader) ReadBytes(length int) ([]byte, error) {
	bytes := make([]byte, length)
	_, err := dr.r.Read(bytes)
	return bytes, err
}

func (dr *DataReader) ReadRune() (rune, error) {
	r, _, err := dr.r.ReadRune()
	return r, err
}

// Read reads a byte.
func (dr *DataReader) ReadBool() (bool, error) {
	value, err := dr.r.ReadByte()
	if err != nil {
		return false, err
	}
	if value < 0 {
		return false, io.EOF
	}
	return value != 0, err
}

// Read reads a int8.
func (dr *DataReader) ReadInt8() (int8, error) {
	var value int8
	err := binary.Read(dr.r, binary.BigEndian, &value)
	return value, err
}

// Read reads a uint8.
func (dr *DataReader) ReadUint8() (uint8, error) {
	var value uint8
	err := binary.Read(dr.r, binary.BigEndian, &value)
	return value, err
}

// Read reads a int16.
func (dr *DataReader) ReadInt16() (int16, error) {
	var value int16
	err := binary.Read(dr.r, binary.BigEndian, &value)
	return value, err
}

// Read reads a uint16.
func (dr *DataReader) ReadUint16() (uint16, error) {
	var value uint16
	err := binary.Read(dr.r, binary.BigEndian, &value)
	return value, err
}

// Read reads a int32.
func (dr *DataReader) ReadInt32() (int32, error) {
	var value int32
	err := binary.Read(dr.r, binary.BigEndian, &value)
	return value, err
}

// Read reads a uint32.
func (dr *DataReader) ReadUint32() (uint32, error) {
	var value uint32
	err := binary.Read(dr.r, binary.BigEndian, &value)
	return value, err
}

// Read reads a float32.
func (dr *DataReader) ReadFloat32() (float32, error) {
	var value float32
	err := binary.Read(dr.r, binary.BigEndian, &value)
	return value, err
}

// ReadChar reads a java char , it is a uint16.
func (dr *DataReader) ReadChar() (uint16, error) {
	return dr.ReadUint16()
}

// ReadInt64 reads a int64.
func (dr *DataReader) ReadInt64() (int64, error) {
	var value int64
	err := binary.Read(dr.r, binary.BigEndian, &value)
	if err != nil {
		return 0, err
	}
	return value, nil
}

// ReadInt64 reads a uint64.
func (dr *DataReader) ReadUint64() (uint64, error) {
	var value uint64
	err := binary.Read(dr.r, binary.BigEndian, &value)
	if err != nil {
		return 0, err
	}
	return value, nil
}

// ReadInt64 reads a float64.
func (dr *DataReader) ReadFloat64() (float64, error) {
	var value float64
	err := binary.Read(dr.r, binary.BigEndian, &value)
	if err != nil {
		return 0, err
	}
	return value, nil
}

// ReadInt64 reads a java long, it is a int64.
func (dr *DataReader) ReadLong() (int64, error) {
	return dr.ReadInt64()
}

// ReadString reads a string with bytes[length].
func (dr *DataReader) ReadString(length int) (string, error) {
	runes := make([]rune, 0)
	total := 0
	for total < int(length) {
		r, size, err := dr.r.ReadRune()
		if err != nil {
			return string(runes), err
		}
		total = total + size
		runes = append(runes, r)
	}

	if total > int(length) {
		err := dr.r.UnreadRune()
		if err != nil {
			return string(runes), err
		}
		bs := make([]byte, total-int(length))
		_, err = dr.r.Read(bs)
		if err != nil {
			return string(runes), err
		}
		bytes := append([]byte(string(runes)), bs...)
		return string(bytes), nil
	}

	return string(runes), nil
}
