package math

import "fmt"

// Galois field with prime power p^n, GF(p^n)
type GaloisField struct {
	p    int
	n    int
	prim int // irreducible primitive polynomial
}

func NewGaloisField(p int, n int, prim int) GaloisField {
	return GaloisField{
		p:    p,
		n:    n,
		prim: prim,
	}
}

// as byte will be used, only p=2 and n=8 is considered, GF(2^8) = GF(256)
// 0x11d = 100011101 is a common primitive polynomial used for this case
// primitive polynomial: x^8 + x^4 + x^3 + x^2 + 1
var GF256 = NewGaloisField(2, 8, 0x11d)

// in binary galois field, addition (and subtraction) is equivalent to XOR
// in GF(2^8), addition between 2 byte input is expected
func (gf GaloisField) Add(a, b int) int {
	return a ^ b
}

// in GF(2^8), multiplication between 2 byte input is expected
func (gf GaloisField) Multiply(a, b int) int {
	// find position of most significan bit (1)
	bitLength := func(n int) int {
		length := 0
		for n > 0 {
			n >>= 1
			length++
		}
		return length
	}

	// bitwise carry-less multiplication
	// in GF(2^8), 2 byte input is expected resulting in maximum as uint16
	clMult := func(x, y int) int {
		z, i := 0, 0
		for (y >> i) > 0 {
			if (y & (1 << i)) != 0 {
				z ^= x << i
			}
			i++
		}
		return z
	}

	// bitwise carry-less long division, returning remainder
	// in GF(2^8), 2 uint16 input is expected resulting in byte
	clDivRemainder := func(dividend, divisor int) int {
		dividendLen := bitLength(dividend)
		divisorLen := bitLength(divisor)

		if dividendLen < divisorLen {
			return dividend
		}

		for i := dividendLen - divisorLen; i >= 0; i-- {
			if dividend&(1<<(i+divisorLen-1)) != 0 {
				dividend ^= divisor << i
			}
		}

		return dividend
	}

	// product with modular reduction
	return clDivRemainder(clMult(a, b), gf.prim)
}

var expLUT, logLUT = GF256.initLookUpTable()

// lookup table for exponentials and logs
func (gf GaloisField) initLookUpTable() ([]byte, []byte) {
	expLUT := make([]byte, 256)
	logLUT := make([]byte, 256)

	x := byte(1)
	for i := range 255 {
		expLUT[i] = x
		logLUT[x] = byte(i)
		x = byte(gf.Multiply(int(x), 2))
	}
	return expLUT, logLUT
}

func (gf GaloisField) FastMultiply(a, b byte) byte {
	if a == 0 || b == 0 {
		return 0
	}

	return expLUT[(int(logLUT[a])+int(logLUT[b]))%255]
}

func (gf GaloisField) FastDivision(divident, divisor byte) (byte, error) {
	if divisor == 0 {
		return 0, fmt.Errorf("divisor cannot be 0")
	}
	if divident == 0 {
		return 0, nil
	}

	return expLUT[(logLUT[divident]+255-logLUT[divisor])%255], nil
}

func (gf GaloisField) FastPower(x byte, pow int) byte {
	return expLUT[(logLUT[x]*byte(pow))%255]
}

func (gf GaloisField) FastInverse(x byte) byte {
	return expLUT[255-logLUT[x]]
}
