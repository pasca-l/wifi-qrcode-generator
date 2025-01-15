package utils

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Byte byte
type Bytes []Byte

type Bit bool
type Bits []Bit

func NewBytes(src interface{}) (Bytes, error) {
	var bs Bytes

	switch v := src.(type) {
	case byte:
		bs = append(bs, Byte(v))

	case []byte:
		for _, b := range v {
			bs = append(bs, Byte(b))
		}

	case int:
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.BigEndian, int64(v))
		if err != nil {
			return nil, err
		}
		for _, b := range buf.Bytes() {
			bs = append(bs, Byte(b))
		}

	default:
		return Bytes{}, fmt.Errorf("unexpected type: %T", v)
	}

	return bs, nil
}

func (bs Bytes) ToBits(digits int) Bits {
	bits := make(Bits, 0)
	for _, b := range bs {
		bits = append(bits, b.ToBits(digits)...)
	}
	return bits
}

func (b Byte) ToBits(digits int) Bits {
	bits := make(Bits, digits)
	for i := range digits {
		// 1 << i, shifts 00000001 to the left for `i` amount
		// bitwise AND operator, gives non-zero value if the `i`th bit exists
		bits[digits-1-i] = (b & (1 << i)) != 0
	}
	return bits
}

// pads 0 until the last byte is filled
func (b Bits) AppendBitPadding() Bits {
	paddingBitLength := (8 - (len(b) % 8)) % 8
	padding := make(Bits, paddingBitLength)
	return append(b, padding...)
}

// pads with alternating 0xEC and 0x11 until capacity
func (b Bits) AppendBytePadding(capacity int) Bits {
	// constant padding bytes
	padBytes := []byte{0xEC, 0x11}

	paddingByteLength := (capacity - len(b)) / 8
	for i := range paddingByteLength {
		padding, _ := NewBytes(padBytes[i%2])
		b = append(b, padding.ToBits(8)...)
	}

	return b
}

func (b Bits) ToBitString() string {
	paddingLength := (8 - (len(b) % 8)) % 8
	padding := make(Bits, paddingLength)

	// add padding from the left
	bits := append(padding, b...)

	var bitString string
	for i, b := range bits {
		if b {
			bitString += "1"
		} else {
			bitString += "0"
		}
		if (i+1)%8 == 0 && i+1 != len(bits) {
			bitString += " "
		}
	}

	return bitString
}
