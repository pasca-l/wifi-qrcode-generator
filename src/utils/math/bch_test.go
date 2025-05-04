package math

import (
	"reflect"
	"testing"

	"github.com/pasca-l/wifi-qrcode-generator/utils"
)

func TestBCHEncodeFormatInfo(t *testing.T) {
	testcases := []struct {
		ecl     utils.Bits
		mask    utils.Bits
		want    utils.Bits
		wantErr error
	}{
		{
			ecl:     utils.Bits{false, true},          // ecl of L
			mask:    utils.Bits{false, false, false},  // mask type of 0
			want:    utils.Bytes{119, 196}.ToBits(15), // 0x77C4
			wantErr: nil,
		},
	}

	for _, tt := range testcases {
		t.Run("testing BCH.EncodeFormatInfo()", func(t *testing.T) {
			bch := BCH{}
			got, err := bch.EncodeFormatInfo(tt.ecl, tt.mask)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("BCH.EncodeFormatInfo() error = '%v'; expected '%v'", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BCH.EncodeFormatInfo() = %v; want %v", got, tt.want)
			}
		})
	}
}

func TestBCHEncodeVersionInfo(t *testing.T) {
	testcases := []struct {
		version utils.Bits
		want    utils.Bits
		wantErr error
	}{
		{
			version: utils.Bits{false, false, false, true, true, true}, // version 7
			want:    utils.Bytes{0, 124, 148}.ToBits(18),               // 0x007C94
			wantErr: nil,
		},
	}

	for _, tt := range testcases {
		t.Run("testing BCH.EncodeVersionInfo()", func(t *testing.T) {
			bch := BCH{}
			got, err := bch.EncodeVersionInfo(tt.version)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("BCH.EncodeVersionInfo() error = '%v'; expected '%v'", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BCH.EncodeVersionInfo() = %v; want %v", got, tt.want)
			}
		})
	}
}
