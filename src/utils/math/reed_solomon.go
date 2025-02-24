package math

import "github.com/pasca-l/wifi-qrcode-generator/utils"

type ReedSolomon struct{}

func (rs ReedSolomon) Encode(msg utils.Bytes, nsym int) (utils.Bytes, error) {
	gf := NewGaloisField(2, 8, 0x11d) // set to GF(2^8)
	gen := GeneratorPoly(gf, nsym)

	// add padding before dividing with irreducible generator polynomial
	padMsg := append(msg, make(utils.Bytes, len(gen)-1)...)
	padMsgPoly, err := NewPolynomial(padMsg.ToNativeBytes())
	if err != nil {
		return utils.Bytes{}, err
	}

	_, r := padMsgPoly.Divide(gf, gen)
	rBytes, err := utils.NewBytes(r.ToBytes())
	if err != nil {
		return utils.Bytes{}, err
	}

	// append remainder to the message as codeword
	rsMsg := append(msg, rBytes...)
	return rsMsg, nil
}
