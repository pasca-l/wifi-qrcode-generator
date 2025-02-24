package qrcode

import (
	"reflect"
	"testing"

	"github.com/pasca-l/wifi-qrcode-generator/utils"
)

func TestQRCodeSpecEncodeSrc(t *testing.T) {
	testcases := []struct {
		src     string
		ecl     ErrorCorrectionLevel
		mode    EncodeMode
		version Version
		want    utils.Bits
		wantErr error
	}{
		{
			src:     "Hello World!",
			ecl:     L,
			mode:    BinaryMode,
			version: 1,
			want:    utils.Bytes{70, 4, 134, 86, 198, 198, 242, 5, 118, 247, 38, 198, 66, 16, 236, 17, 236, 17, 236}.ToBits(19 * 8),
			wantErr: nil,
		},
	}

	for _, tt := range testcases {
		t.Run("testing QRCodeSpec.EncodeSrc()", func(t *testing.T) {
			spec := QRCodeSpec{
				src:     tt.src,
				ecl:     tt.ecl,
				mode:    tt.mode,
				version: tt.version,
			}
			got, err := spec.EncodeSrc()
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("QRCodeSpec.EncodeSrc()() error = '%v'; expected '%v'", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QRCodeSpec.EncodeSrc() = %v; expected %v", got, tt.want)
			}
		})
	}
}
