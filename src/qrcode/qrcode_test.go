package qrcode

import (
	"reflect"
	"testing"

	"github.com/pasca-l/wifi-qrcode-generator/utils"
)

func TestQRCodeSpecEncodeSrc(t *testing.T) {
	testcases := []struct {
		src     string
		mode    EncodeMode
		version Version
		ecl     ErrorCorrectionLevel
		want    utils.Bytes
		wantErr error
	}{
		{
			src:     "Hello World!",
			mode:    BinaryMode,
			version: 1,
			ecl:     L,
			want:    utils.Bytes{64, 196, 134, 86, 198, 198, 242, 5, 118, 247, 38, 198, 66, 16, 236, 17, 236, 17, 236},
			wantErr: nil,
		},
	}

	for _, tt := range testcases {
		t.Run("testing QRCodeSpec.EncodeSrc()", func(t *testing.T) {
			spec := QRCodeSpec{
				mode:    tt.mode,
				version: tt.version,
				ecl:     tt.ecl,
			}
			got, err := spec.EncodeSrc(tt.src)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("QRCodeSpec.EncodeSrc() error = '%v'; expected '%v'", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QRCodeSpec.EncodeSrc() = %v; expected %v", got, tt.want)
			}
		})
	}
}

func TestQRCodeSpecApplyErrorCorrection(t *testing.T) {
	testcases := []struct {
		msg     utils.Bytes
		mode    EncodeMode
		version Version
		ecl     ErrorCorrectionLevel
		want    utils.Bytes
		wantErr error
	}{
		{
			// encoded "Hello World!"
			msg:     utils.Bytes{64, 196, 134, 86, 198, 198, 242, 5, 118, 247, 38, 198, 66, 16, 236, 17, 236, 17, 236},
			mode:    BinaryMode,
			version: 1,
			ecl:     L,
			want:    utils.Bytes{64, 196, 134, 86, 198, 198, 242, 5, 118, 247, 38, 198, 66, 16, 236, 17, 236, 17, 236, 30, 201, 34, 105, 71, 33, 134},
			wantErr: nil,
		},
		{
			msg:     utils.Bytes{64, 196, 134, 86, 198, 198, 242, 5, 118, 247, 38, 198, 66, 16, 236, 17, 236, 17, 236},
			mode:    BinaryMode,
			version: 3,
			ecl:     H,
			want:    utils.Bytes{64, 196, 134, 86, 198, 198, 242, 5, 118, 247, 38, 198, 66, 16, 236, 17, 236, 17, 236, 160, 200, 191, 131, 142, 240, 194, 73, 34, 225, 201, 8, 149, 179, 18, 91, 147, 3, 170, 132, 123, 117, 64, 196, 134, 86, 198, 198, 242, 5, 118, 247, 38, 198, 66, 16, 236, 17, 236, 17, 236, 160, 200, 191, 131, 142, 240, 194, 73, 34, 225, 201, 8, 149, 179, 18, 91, 147, 3, 170, 132, 123, 117},
			wantErr: nil,
		},
	}

	for _, tt := range testcases {
		t.Run("testing QRCodeSpec.ApplyErrorCorrection()", func(t *testing.T) {
			spec := QRCodeSpec{
				mode:    tt.mode,
				version: tt.version,
				ecl:     tt.ecl,
			}
			got, err := spec.ApplyErrorCorrection(tt.msg)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("QRCodeSpec.ApplyErrorCorrection() error = '%v'; expected '%v'", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QRCodeSpec.ApplyErrorCorrection() = %v; expected %v", got, tt.want)
			}
		})
	}
}
