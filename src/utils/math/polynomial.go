package math

import "fmt"

type Polynomial []byte

func NewPolynomial(src interface{}) (Polynomial, error) {
	var poly Polynomial

	switch v := src.(type) {
	case []byte:
		for _, b := range v {
			poly = append(poly, b)
		}

	default:
		return Polynomial{}, fmt.Errorf("unexpected type: %T", v)
	}

	return poly, nil
}

func (p Polynomial) ToBytes() []byte {
	return p
}

func GeneratorPoly(gf GaloisField, nsym int) Polynomial {
	g := Polynomial{1}

	// iterates over the number of error correction symbols
	for i := range nsym {
		poly := Polynomial{1, gf.FastPower(2, i)}
		g = g.Multiply(gf, poly)
	}

	return g
}

func (p Polynomial) Scale(gf GaloisField, x byte) Polynomial {
	r := make([]byte, 0, len(p))
	for i := range len(p) {
		r = append(r, gf.FastMultiply(p[i], x))
	}
	return r
}

func (p Polynomial) Add(gf GaloisField, q Polynomial) Polynomial {
	r := make([]byte, max(len(p), len(q)))
	for i := range len(p) {
		r[i+len(r)-len(p)] = p[i]
	}
	for i := range len(q) {
		r[i+len(r)-len(q)] = byte(gf.Add(int(r[i+len(r)-len(q)]), int(q[i])))
	}
	return r
}

func (p Polynomial) Multiply(gf GaloisField, q Polynomial) Polynomial {
	r := make([]byte, len(p)+len(q)-1)
	for i := range len(p) {
		for j := range len(q) {
			r[i+j] = byte(gf.Add(int(r[i+j]), int(gf.FastMultiply(p[i], q[j]))))
		}
	}
	return r
}

// polynomial division by using extended synthetic division,
// and optimized for GF(2^p) computation
func (p Polynomial) Divide(gf GaloisField, divisor Polynomial) (Polynomial, Polynomial) {
	r := make([]byte, len(p))
	copy(r, p)

	for i := range len(p) - (len(divisor) - 1) {

		// avoid case of log(0) as it is undefined
		if r[i] != 0 {

			// in synthetic division, the first coefficient is skipped,
			// as the first value is only used for normalizing
			for j := 1; j < len(divisor); j++ {

				// avoid case of log(0) as it is undefined
				if divisor[j] != 0 {
					r[i+j] = byte(gf.Add(int(r[i+j]), int(gf.FastMultiply(divisor[j], r[i]))))
				}
			}
		}
	}

	// division result contains both quotient and remainder,
	// where the remainder has the same degree as the divisor
	separator := len(p) - (len(divisor) - 1)

	// returning quotient, and remainder
	return r[:separator], r[separator:]
}
