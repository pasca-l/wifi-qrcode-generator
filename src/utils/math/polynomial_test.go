package math

import (
	"errors"
	"reflect"
	"testing"
)

func TestNewPolynomial(t *testing.T) {
	testcases := []struct {
		args    interface{}
		want    Polynomial
		wantErr error
	}{
		{
			args:    []byte{1},
			want:    Polynomial{1},
			wantErr: nil,
		},
		{
			args:    "Hello World!",
			want:    Polynomial{},
			wantErr: errors.New("unexpected type: string"),
		},
	}

	for _, tt := range testcases {
		t.Run("testing NewPolynomial()", func(t *testing.T) {
			got, err := NewPolynomial(tt.args)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("NewPolynomial() error = '%v'; expected '%v'", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPolynomial(%v) = %v; expected %v", tt.args, got, tt.want)
			}
		})
	}
}

func TestGeneratorPoly(t *testing.T) {
	testcases := []struct {
		nsym int
		want Polynomial
	}{
		{
			nsym: 10,
			want: Polynomial{233, 150, 83, 27, 112, 147, 154, 25, 219, 180, 119},
		},
	}

	for _, tt := range testcases {
		t.Run("testing GeneratorPoly()", func(t *testing.T) {
			gf := NewGaloisField(2, 8, 0x11d)
			if got := GeneratorPoly(gf, tt.nsym); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GeneratorPoly() = %v; expected %v", got, tt.want)
			}
		})
	}
}
