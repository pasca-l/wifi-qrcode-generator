package math

import "testing"

// referenced: http://www.ee.unb.ca/cgi-bin/tervo/calc2.pl

func TestGaloisFieldAdd(t *testing.T) {
	testcases := []struct {
		a, b int
		want int
	}{
		{
			a:    0b10001001,
			b:    0b00101010,
			want: 0b10100011,
		},
	}

	for _, tt := range testcases {
		t.Run("testing GaloisField.Add()", func(t *testing.T) {
			gf := NewGaloisField(2, 8, 0x11d)
			if got := gf.Add(tt.a, tt.b); got != tt.want {
				t.Errorf("GaloisField.Add(%v, %v) = %v; expected %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestGaloisFieldMultiply(t *testing.T) {
	testcases := []struct {
		a, b int
		want int
	}{
		{
			a:    0b10001001,
			b:    0b00101010,
			want: 0b11000011,
		},
	}

	for _, tt := range testcases {
		t.Run("testing GaloisField.Multiply()", func(t *testing.T) {
			gf := NewGaloisField(2, 8, 0x11d)
			if got := gf.Multiply(tt.a, tt.b); got != tt.want {
				t.Errorf("GaloisField.Multiply(%v, %v) = %v; expected %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestGaloisFieldFastMultiply(t *testing.T) {
	testcases := []struct {
		a, b byte
		want byte
	}{
		{
			a:    0b10001001,
			b:    0b00101010,
			want: 0b11000011,
		},
		{
			a:    0b00000001,
			b:    0b00101010,
			want: 0b00101010,
		},
	}

	for _, tt := range testcases {
		t.Run("testing GaloisField.FastMultiply()", func(t *testing.T) {
			gf := NewGaloisField(2, 8, 0x11d)
			if got := gf.FastMultiply(tt.a, tt.b); got != tt.want {
				t.Errorf("GaloisField.FastMultiply(%v, %v) = %v; expected %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestGaloisFieldFastDivision(t *testing.T) {
	testcases := []struct {
		divident, divisor byte
		want              byte
		wantErr           error
	}{
		{
			divident: 0b10001001,
			divisor:  0b00101010,
			want:     0b11011100,
			wantErr:  nil,
		},
	}

	for _, tt := range testcases {
		t.Run("testing GaloisField.FastDivision()", func(t *testing.T) {
			gf := NewGaloisField(2, 8, 0x11d)
			got, err := gf.FastDivision(tt.divident, tt.divisor)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("GaloisField.FastDivision() with error '%v'; expected '%v'", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("GaloisField.FastDivision(%v, %v) = %v; expected %v", tt.divident, tt.divisor, got, tt.want)
			}
		})
	}
}

func TestGaloisFieldFastPower(t *testing.T) {
	testcases := []struct {
		x    byte
		pow  int
		want byte
	}{
		{
			x:    0b10001001,
			pow:  2,
			want: 0b01010010,
		},
	}

	for _, tt := range testcases {
		t.Run("testing GaloisField.FastPower()", func(t *testing.T) {
			gf := NewGaloisField(2, 8, 0x11d)
			if got := gf.FastPower(tt.x, tt.pow); got != tt.want {
				t.Errorf("GaloisField.FastPower(%v, %v) = %v; expected %v", tt.x, tt.pow, got, tt.want)
			}
		})
	}
}

func TestGaloisFieldFastInverse(t *testing.T) {
	testcases := []struct {
		x    byte
		want byte
	}{
		{
			x:    0b00101010,
			want: 0b00011111,
		},
	}

	for _, tt := range testcases {
		t.Run("testing GaloisField.FastInverse()", func(t *testing.T) {
			gf := NewGaloisField(2, 8, 0x11d)
			if got := gf.FastInverse(tt.x); got != tt.want {
				t.Errorf("GaloisField.FastInverse(%v) = %v; expected %v", tt.x, got, tt.want)
			}
		})
	}
}
