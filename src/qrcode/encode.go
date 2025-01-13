package qrcode

import (
	"fmt"
	"regexp"

	"github.com/pasca-l/wifi-qrcode-generator/utils"
)

type EncodeMode string

const (
	BinaryMode  EncodeMode = "binary"  // 8 bits per character
	NumericMode EncodeMode = "numeric" // 10 bits per 3 digits
)

type EncodeModeIndicator byte

const (
	Terminator EncodeModeIndicator = 0 // '0000'
	NumericInd EncodeModeIndicator = 1 // '0001'
	BinaryInd  EncodeModeIndicator = 4 // '0100'
)

// source string will only be converted to one encoding mode,
// even though it is possible to use multiple encodings within a single QR code
func getEncodeMode(src string) EncodeMode {
	isNumeric := regexp.MustCompile(`^[0-9]+$`).MatchString(src)
	if isNumeric {
		return NumericMode
	}

	return BinaryMode
}

func getIndicatorBits(mode EncodeMode) (utils.Bits, error) {
	switch mode {
	case NumericMode:
		bytes, err := utils.NewBytes(byte(NumericInd))
		if err != nil {
			return utils.Bits{}, err
		}
		return bytes.ToBits(4), nil

	case BinaryMode:
		bytes, err := utils.NewBytes(byte(BinaryInd))
		if err != nil {
			return utils.Bits{}, err
		}
		return bytes.ToBits(4), nil

	default:
		return utils.Bits{}, fmt.Errorf("unexpected mode: %s", mode)
	}
}

func getTerminatorBits() (utils.Bits, error) {
	bytes, err := utils.NewBytes(byte(Terminator))
	if err != nil {
		return utils.Bits{}, err
	}
	return bytes.ToBits(4), nil
}

func convertSrcToBits(mode EncodeMode, src string) (utils.Bits, error) {
	switch mode {
	case NumericMode:
		// TODO: implement bit conversion
		return utils.Bits{}, nil

	case BinaryMode:
		srcBytes := []byte(src)
		bytes, err := utils.NewBytes(srcBytes)
		if err != nil {
			return utils.Bits{}, err
		}
		return bytes.ToBits(8), nil

	default:
		return utils.Bits{}, fmt.Errorf("unexpected mode: %s", mode)
	}
}

func getSrcCountBits(src string, ind int) (utils.Bits, error) {
	srcLength := len([]byte(src))
	bytes, err := utils.NewBytes(srcLength)
	if err != nil {
		return utils.Bits{}, err
	}
	return bytes.ToBits(ind), nil
}
