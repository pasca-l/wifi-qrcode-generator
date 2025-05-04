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
			want:    utils.Bytes{70, 4, 134, 86, 198, 198, 242, 5, 118, 247, 38, 198, 66, 16, 236, 17, 236, 17, 236},
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
			msg:     utils.Bytes{70, 4, 134, 86, 198, 198, 242, 5, 118, 247, 38, 198, 66, 16, 236, 17, 236, 17, 236},
			mode:    BinaryMode,
			version: 1,
			ecl:     L,
			want:    utils.Bytes{70, 4, 134, 86, 198, 198, 242, 5, 118, 247, 38, 198, 66, 16, 236, 17, 236, 17, 236, 238, 82, 154, 31, 241, 184, 8},
			wantErr: nil,
		},
		{
			msg:     utils.Bytes{70, 4, 134, 86, 198, 198, 242, 5, 118, 247, 38, 198, 66, 16, 236, 17, 236, 17, 236},
			mode:    BinaryMode,
			version: 3,
			ecl:     H,
			want:    utils.Bytes{70, 4, 134, 86, 198, 198, 242, 5, 118, 247, 38, 198, 66, 16, 236, 17, 236, 17, 236, 75, 80, 221, 190, 141, 1, 121, 191, 250, 57, 189, 152, 212, 203, 171, 113, 135, 37, 18, 140, 235, 207, 70, 4, 134, 86, 198, 198, 242, 5, 118, 247, 38, 198, 66, 16, 236, 17, 236, 17, 236, 75, 80, 221, 190, 141, 1, 121, 191, 250, 57, 189, 152, 212, 203, 171, 113, 135, 37, 18, 140, 235, 207},
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
