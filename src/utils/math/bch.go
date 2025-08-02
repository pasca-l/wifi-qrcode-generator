package math

import (
	"fmt"
	"slices"

	"github.com/pasca-l/wifi-qrcode-generator/utils"
)

type BCH struct{}

// appends 10 bit error correction for 5 bit format information
func (bch BCH) EncodeFormatInfo(ecl utils.Bits, mask utils.Bits) (utils.Bits, error) {
	if len(ecl) != 2 {
		return nil, fmt.Errorf("invalid ecl length: %d, expected 2", len(ecl))
	}
	if len(mask) != 3 {
		return nil, fmt.Errorf("invalid mask length: %d, expected 3", len(mask))
	}

	// convert to native byte
	formatInfo := slices.Concat(ecl, mask)
	formatInfoBytes, err := slices.Concat(utils.Bits{false, false, false}, formatInfo).ToBytes()
	if err != nil {
		return nil, err
	}

	b := formatInfoBytes[0].ToNativeByte()
	rem := int(b)
	for range 10 {
		// applies polynomial division, using 0x537 as the generator polynomial
		// ((rem >> 9) * 0x537) checks if the 10th bit is set, if so XORs
		rem = (rem << 1) ^ ((rem >> 9) * 0x537)
	}
	// rem&0x3FF assures for 10 digit bits
	// 0x5412 (0b101010000010010) is the mask
	encoded := (int(b)<<10 | rem&0x3FF) ^ 0x5412

	bytes, err := utils.NewBytes(encoded)
	if err != nil {
		return nil, err
	}
	bits := bytes.ToBits(15)

	return bits, nil
}

// appends 12 bit error correction for 6 bit version information
func (bch BCH) EncodeVersionInfo(version utils.Bits) (utils.Bits, error) {
	if len(version) != 6 {
		return nil, fmt.Errorf("invalid version length: %d, expected 6", len(version))
	}

	// convert to native byte
	versionBytes, err := append(utils.Bits{false, false}, version...).ToBytes()
	if err != nil {
		return nil, err
	}

	b := versionBytes[0].ToNativeByte()
	rem := int(b)
	for range 12 {
		// applies polynomial division, using 0x1F25 as the generator polynomial
		// ((rem >> 11) * 0x1F25) checks if the 12th bit is set, if so XORs
		rem = (rem << 1) ^ ((rem >> 11) * 0x1F25)
	}
	// rem&0xFFF assures for 12 digit bits
	encoded := int(b)<<12 | rem&0xFFF

	bytes, err := utils.NewBytes(encoded)
	if err != nil {
		return nil, err
	}
	bits := bytes.ToBits(18)

	return bits, nil
}
