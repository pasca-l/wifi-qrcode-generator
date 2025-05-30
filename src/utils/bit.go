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

func NewBytes(src any) (Bytes, error) {
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

func (bs Bytes) ToNativeBytes() []byte {
	bytes := make([]byte, 0, len(bs))
	for _, b := range bs {
		bytes = append(bytes, byte(b))
	}
	return bytes
}

func (bs Bytes) ToBits(digits int) Bits {
	bits := make(Bits, 0, digits)
	for _, b := range bs {
		bits = append(bits, b.ToBits(8)...)
	}
	return bits[len(bits)-digits:]
}

func (b Byte) ToNativeByte() byte {
	return byte(b)
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

func (bs Bits) ToNativeBools() []bool {
	bools := make([]bool, 0, len(bs))
	for _, bit := range bs {
		bools = append(bools, bool(bit))
	}
	return bools
}

// pads 0 until the last byte is filled
func (bs Bits) AppendBitPadding() Bits {
	paddingBitLength := (8 - (len(bs) % 8)) % 8
	padding := make(Bits, paddingBitLength)
	return append(bs, padding...)
}

// pads with alternating 0xEC and 0x11 until capacity
func (bs Bits) AppendBytePadding(capacity int) Bits {
	// constant padding bytes
	padBytes := []byte{0xEC, 0x11}

	paddingByteLength := (capacity - len(bs)) / 8
	for i := range paddingByteLength {
		padding, _ := NewBytes(padBytes[i%2])
		bs = append(bs, padding.ToBits(8)...)
	}

	return bs
}

func (bs Bits) ToBytes() (Bytes, error) {
	if len(bs)%8 != 0 {
		return Bytes{}, fmt.Errorf("bits must have length with multiple of 8: given length %d", len(bs))
	}

	Bs := make(Bytes, 0)
	for i := 0; i < len(bs); i += 8 {
		var B Byte
		for j := range 8 {
			if bs[i+j] {
				// set the bit to 1 if true
				B |= (1 << uint(7-j))
			}
		}
		Bs = append(Bs, B)
	}

	return Bs, nil
}

func (bs Bits) ToBitString() string {
	paddingLength := (8 - (len(bs) % 8)) % 8
	padding := make(Bits, paddingLength)

	// add padding from the left
	bits := append(padding, bs...)

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
