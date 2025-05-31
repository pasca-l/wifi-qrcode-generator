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
			want: Polynomial{1, 216, 194, 158, 111, 193, 194, 111, 213, 157, 193},
		},
	}

	for _, tt := range testcases {
		t.Run("testing GeneratorPoly()", func(t *testing.T) {
			if got := GeneratorPoly(GF256, tt.nsym); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GeneratorPoly() = %v; expected %v", got, tt.want)
			}
		})
	}
}

// referenced: http://www.ee.unb.ca/cgi-bin/tervo/calc2.pl

func TestPolynomialAdd(t *testing.T) {
	testcases := []struct {
		p, q Polynomial
		want Polynomial
	}{
		{
			p:    Polynomial{1, 0, 7, 6},
			q:    Polynomial{1, 6, 3},
			want: Polynomial{1, 1, 1, 5},
		},
	}

	for _, tt := range testcases {
		t.Run("testing Polynomial.Add()", func(t *testing.T) {
			if got := tt.p.Add(GF256, tt.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Polynomial.Add() = %v; expected %v", got, tt.want)
			}
		})
	}
}
func TestPolynomialMultiply(t *testing.T) {
	testcases := []struct {
		p, q Polynomial
		want Polynomial
	}{
		{
			p:    Polynomial{1, 0, 7, 6},
			q:    Polynomial{1, 6, 3},
			want: Polynomial{1, 6, 4, 20, 29, 10},
		},
	}

	for _, tt := range testcases {
		t.Run("testing Polynomial.Multiply()", func(t *testing.T) {
			if got := tt.p.Multiply(GF256, tt.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Polynomial.Multiply() = %v; expected %v", got, tt.want)
			}
		})
	}
}

func TestPolynomialDivide(t *testing.T) {
	testcases := []struct {
		p, q      Polynomial
		quotient  Polynomial
		remainder Polynomial
	}{
		{
			p:         Polynomial{1, 0, 7, 6},
			q:         Polynomial{1, 6, 3},
			quotient:  Polynomial{1, 6},
			remainder: Polynomial{16, 12},
		},
	}

	for _, tt := range testcases {
		t.Run("testing Polynomial.Divide()", func(t *testing.T) {
			q, r := tt.p.Divide(GF256, tt.q)
			if !reflect.DeepEqual(q, tt.quotient) || !reflect.DeepEqual(r, tt.remainder) {
				t.Errorf("Polynomial.Divide() = (%v, %v); expected (%v, %v)", q, r, tt.quotient, tt.remainder)
			}
		})
	}
}
