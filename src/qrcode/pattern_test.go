package qrcode

import (
	"reflect"
	"testing"

	"github.com/pasca-l/wifi-qrcode-generator/utils"
)

func TestCreateFinderPattern(t *testing.T) {
	testcases := []struct {
		want Pattern
	}{
		{
			want: Pattern{
				{true, true, true, true, true, true, true},
				{true, false, false, false, false, false, true},
				{true, false, true, true, true, false, true},
				{true, false, true, true, true, false, true},
				{true, false, true, true, true, false, true},
				{true, false, false, false, false, false, true},
				{true, true, true, true, true, true, true},
			},
		},
	}

	for _, tt := range testcases {
		t.Run("testing createFinderPattern()", func(t *testing.T) {
			got := createFinderPattern()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createFinderPattern() = %v; expected %v", got, tt.want)
			}
		})
	}
}

func TestCreateAlignmentPattern(t *testing.T) {
	testcases := []struct {
		want Pattern
	}{
		{
			want: Pattern{
				{true, true, true, true, true},
				{true, false, false, false, true},
				{true, false, true, false, true},
				{true, false, false, false, true},
				{true, true, true, true, true},
			},
		},
	}

	for _, tt := range testcases {
		t.Run("testing createAlignmentPattern()", func(t *testing.T) {
			got := createAlignmentPattern()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createAlignmentPattern() = %v; expected %v", got, tt.want)
			}
		})
	}
}

func TestAddFunctionalPattern(t *testing.T) {
	testcases := []struct {
		ver     Version
		want    Pattern
		wantErr error
	}{
		{
			ver: 1,
			want: Pattern{
				utils.Bytes{31, 192, 127}.ToBits(21).ToNativeBools(),
				utils.Bytes{16, 64, 65}.ToBits(21).ToNativeBools(),
				utils.Bytes{23, 64, 93}.ToBits(21).ToNativeBools(),
				utils.Bytes{23, 64, 93}.ToBits(21).ToNativeBools(),
				utils.Bytes{23, 64, 93}.ToBits(21).ToNativeBools(),
				utils.Bytes{16, 64, 65}.ToBits(21).ToNativeBools(),
				utils.Bytes{31, 213, 127}.ToBits(21).ToNativeBools(),
				utils.Bytes{0, 0, 0}.ToBits(21).ToNativeBools(),
				utils.Bytes{0, 64, 0}.ToBits(21).ToNativeBools(),
				utils.Bytes{0, 0, 0}.ToBits(21).ToNativeBools(),
				utils.Bytes{0, 64, 0}.ToBits(21).ToNativeBools(),
				utils.Bytes{0, 0, 0}.ToBits(21).ToNativeBools(),
				utils.Bytes{0, 64, 0}.ToBits(21).ToNativeBools(),
				utils.Bytes{0, 0, 0}.ToBits(21).ToNativeBools(),
				utils.Bytes{31, 192, 0}.ToBits(21).ToNativeBools(),
				utils.Bytes{16, 64, 0}.ToBits(21).ToNativeBools(),
				utils.Bytes{23, 64, 0}.ToBits(21).ToNativeBools(),
				utils.Bytes{23, 64, 0}.ToBits(21).ToNativeBools(),
				utils.Bytes{23, 64, 0}.ToBits(21).ToNativeBools(),
				utils.Bytes{16, 64, 0}.ToBits(21).ToNativeBools(),
				utils.Bytes{31, 192, 0}.ToBits(21).ToNativeBools(),
			},
			wantErr: nil,
		},
		{
			ver: 2,
			want: Pattern{
				utils.Bytes{1, 252, 0, 127}.ToBits(25).ToNativeBools(),
				utils.Bytes{1, 4, 0, 65}.ToBits(25).ToNativeBools(),
				utils.Bytes{1, 116, 0, 93}.ToBits(25).ToNativeBools(),
				utils.Bytes{1, 116, 0, 93}.ToBits(25).ToNativeBools(),
				utils.Bytes{1, 116, 0, 93}.ToBits(25).ToNativeBools(),
				utils.Bytes{1, 4, 0, 65}.ToBits(25).ToNativeBools(),
				utils.Bytes{1, 253, 85, 127}.ToBits(25).ToNativeBools(),
				utils.Bytes{0, 0, 0, 0}.ToBits(25).ToNativeBools(),
				utils.Bytes{0, 4, 0, 0}.ToBits(25).ToNativeBools(),
				utils.Bytes{0, 0, 0, 0}.ToBits(25).ToNativeBools(),
				utils.Bytes{0, 4, 0, 0}.ToBits(25).ToNativeBools(),
				utils.Bytes{0, 0, 0, 0}.ToBits(25).ToNativeBools(),
				utils.Bytes{0, 4, 0, 0}.ToBits(25).ToNativeBools(),
				utils.Bytes{0, 0, 0, 0}.ToBits(25).ToNativeBools(),
				utils.Bytes{0, 4, 0, 0}.ToBits(25).ToNativeBools(),
				utils.Bytes{0, 0, 0, 0}.ToBits(25).ToNativeBools(),
				utils.Bytes{0, 4, 1, 240}.ToBits(25).ToNativeBools(),
				utils.Bytes{0, 0, 1, 16}.ToBits(25).ToNativeBools(),
				utils.Bytes{1, 252, 1, 80}.ToBits(25).ToNativeBools(),
				utils.Bytes{1, 4, 1, 16}.ToBits(25).ToNativeBools(),
				utils.Bytes{1, 116, 1, 240}.ToBits(25).ToNativeBools(),
				utils.Bytes{1, 116, 0, 0}.ToBits(25).ToNativeBools(),
				utils.Bytes{1, 116, 0, 0}.ToBits(25).ToNativeBools(),
				utils.Bytes{1, 4, 0, 0}.ToBits(25).ToNativeBools(),
				utils.Bytes{1, 252, 0, 0}.ToBits(25).ToNativeBools(),
			},
			wantErr: nil,
		},
	}

	for _, tt := range testcases {
		t.Run("testing addFunctionalPattern()", func(t *testing.T) {
			size := calcSizeFromVersion(tt.ver)
			pat := NewPattern(size)
			got, err := pat.addFunctionPattern(tt.ver)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("addFunctionalPattern() error = '%v'; expected '%v'", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addFunctionalPattern() = %v; expected %v", got, tt.want)
			}
		})
	}
}
