package utils

import (
	"errors"
	"reflect"
	"testing"
)

func TestNewBytes(t *testing.T) {
	testcases := []struct {
		args    interface{}
		want    Bytes
		wantErr error
	}{
		{
			args:    byte(1),
			want:    Bytes{1},
			wantErr: nil,
		},
		{
			args:    1234,
			want:    Bytes{0, 0, 0, 0, 0, 0, 4, 210},
			wantErr: nil,
		},
		{
			args:    []byte("Hello World!"),
			want:    Bytes{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100, 33},
			wantErr: nil,
		},
		{
			args:    "Hello World!",
			want:    Bytes{},
			wantErr: errors.New("unexpected type: string"),
		},
	}

	for _, tt := range testcases {
		t.Run("testing NewBytes()", func(t *testing.T) {
			got, err := NewBytes(tt.args)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("NewBytes() error = '%v'; expected '%v'", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBytes(%v) = %v; expected %v", tt.args, got, tt.want)
			}
		})
	}
}

func TestBytesToBits(t *testing.T) {
	testcases := []struct {
		bs   Bytes
		args int
		want Bits
	}{
		{
			bs:   Bytes{1},
			args: 1,
			want: Bits{true},
		},
		{
			// []byte{1, 1} -> {0001, 0001}
			bs:   Bytes{1, 1},
			args: 4,
			want: Bits{false, false, false, true, false, false, false, true},
		},
	}

	for _, tt := range testcases {
		t.Run("testing Bytes.ToBits()", func(t *testing.T) {
			if got := tt.bs.ToBits(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bytes.ToBits(%v) = %v; expected %v", tt.args, got, tt.want)
			}
		})
	}
}

func TestByteToBits(t *testing.T) {
	testcases := []struct {
		b    Byte
		args int
		want Bits
	}{
		{
			b:    Byte(0),
			args: 1,
			want: Bits{false},
		},
		{
			b:    Byte(1),
			args: 1,
			want: Bits{true},
		},
		{
			// byte(1) -> 001
			b:    Byte(1),
			args: 3,
			want: Bits{false, false, true},
		},
		{
			// byte(2) -> 10 -> (true,) false
			b:    Byte(2),
			args: 1,
			want: Bits{false},
		},
	}

	for _, tt := range testcases {
		t.Run("testing Byte.ToBits()", func(t *testing.T) {
			if got := tt.b.ToBits(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Byte.ToBits(%v) = %v; expected %v", tt.args, got, tt.want)
			}
		})
	}
}

func TestBitsAppendBytePadding(t *testing.T) {
	testcases := []struct {
		b    Bits
		args int
		want Bits
	}{
		{
			// byte() -> 0xEC
			b:    Bits{},
			args: 8,
			want: Bits{true, true, true, false, true, true, false, false},
		},
		{
			// byte() -> {0xEC, 0x11}
			b:    Bits{},
			args: 16,
			want: Bits{true, true, true, false, true, true, false, false, false, false, false, true, false, false, false, true},
		},
	}

	for _, tt := range testcases {
		t.Run("testing Bits.AppendBytePadding()", func(t *testing.T) {
			if got := tt.b.AppendBytePadding(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bits.AppendBytePadding(%v) = %v; expected %v", tt.args, got, tt.want)
			}
		})
	}
}

func TestBitsAppendBitPadding(t *testing.T) {
	testcases := []struct {
		b    Bits
		want Bits
	}{
		{
			b:    Bits{},
			want: Bits{},
		},
		{
			b:    Bits{true},
			want: Bits{true, false, false, false, false, false, false, false},
		},
		{
			b:    Bits{false, true, true, false, true, false, false, true},
			want: Bits{false, true, true, false, true, false, false, true},
		},
	}

	for _, tt := range testcases {
		t.Run("testing Bits.AppendBitPadding()", func(t *testing.T) {
			if got := tt.b.AppendBitPadding(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bits.AppendBitPadding() = %v; expected %v", got, tt.want)
			}
		})
	}
}

func TestBitsToBitString(t *testing.T) {
	testcases := []struct {
		b    Bits
		want string
	}{
		{
			b:    Bits{},
			want: "",
		},
		{
			b:    Bits{true},
			want: "00000001",
		},
		{
			b:    Bits{true, true, true, false, true, false, false, true, true},
			want: "00000001 11010011",
		},
	}

	for _, tt := range testcases {
		t.Run("testing Bits.ToBitString()", func(t *testing.T) {
			if got := tt.b.ToBitString(); got != tt.want {
				t.Errorf("Bits.ToBitString() = %v; expected %v", got, tt.want)
			}
		})
	}
}
