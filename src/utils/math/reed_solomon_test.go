package math

import (
	"reflect"
	"testing"

	"github.com/pasca-l/wifi-qrcode-generator/utils"
)

func TestReedSolomonEncode(t *testing.T) {
	testcases := []struct {
		msg     utils.Bytes
		nsym    int
		want    utils.Bytes
		wantErr error
	}{
		{
			msg:     utils.Bytes{1},
			nsym:    1,
			want:    utils.Bytes{1, 1},
			wantErr: nil,
		},
		{
			// encoded "Hello World!" with QRCodeSpec{mode: binary, version: 1}
			msg: utils.Bytes{64, 196, 134, 86, 198, 198, 242, 5, 118, 247, 38, 198, 66, 16, 236, 17, 236, 17, 236},
			// nsym determined by QRCodeSpec{version: 1, ecl: L}, which uses
			// Block{blockLength:26, codewordLength:19} -> 26-19 = 7 as nsym
			nsym:    7,
			want:    utils.Bytes{64, 196, 134, 86, 198, 198, 242, 5, 118, 247, 38, 198, 66, 16, 236, 17, 236, 17, 236, 30, 201, 34, 105, 71, 33, 134},
			wantErr: nil,
		},
	}

	for _, tt := range testcases {
		t.Run("testing ReedSolomon.Encode()", func(t *testing.T) {
			rs := ReedSolomon{}
			got, err := rs.Encode(tt.msg, tt.nsym)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("ReedSolomon.Encode() error = '%v'; expected '%v'", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReedSolomon.Encode(%v, %v) = %v; expected %v", tt.msg, tt.nsym, got, tt.want)
			}
		})
	}
}
