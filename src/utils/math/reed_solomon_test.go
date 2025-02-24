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
			want:    utils.Bytes{1, 71},
			wantErr: nil,
		},
		{
			// encoded "Hello World!" with mode: binary, version: 1
			msg: utils.Bytes{70, 4, 134, 86, 198, 198, 242, 5, 118, 247, 38, 198, 66, 16, 236, 17, 236, 17, 236},
			//  nsym determined by ecl: L for the spec above
			nsym:    7,
			want:    utils.Bytes{70, 4, 134, 86, 198, 198, 242, 5, 118, 247, 38, 198, 66, 16, 236, 17, 236, 17, 236, 238, 82, 154, 31, 241, 184, 8},
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
