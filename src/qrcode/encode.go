package qrcode

import "regexp"

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

func getEncodeMode(src string) EncodeMode {
	isNumeric := regexp.MustCompile(`^[0-9]+$`).MatchString(src)
	if isNumeric {
		return NumericMode
	}

	return BinaryMode
}
