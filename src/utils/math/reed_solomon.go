package math

import (
	"slices"

	"github.com/pasca-l/wifi-qrcode-generator/utils"
)

type ReedSolomon struct{}

func (rs ReedSolomon) Encode(msg utils.Bytes, nsym int) (utils.Bytes, error) {
	gen := GeneratorPoly(GF256, nsym)

	// add padding before dividing with irreducible generator polynomial
	padMsg := slices.Concat(msg, make(utils.Bytes, len(gen)-1))
	padMsgPoly, err := NewPolynomial(padMsg.ToNativeBytes())
	if err != nil {
		return utils.Bytes{}, err
	}

	_, r := padMsgPoly.Divide(GF256, gen)
	rBytes, err := utils.NewBytes(r.ToBytes())
	if err != nil {
		return utils.Bytes{}, err
	}

	// append remainder to the message as codeword
	rsMsg := slices.Concat(msg, rBytes)
	return rsMsg, nil
}
